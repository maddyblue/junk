// Package lsp implements a Language Server Protocol client.
// See https://langserver.org/.
package lsp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/mjibson/acmepls/lsp-client/protocol"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	Capabilities protocol.ServerCapabilities

	grp    *errgroup.Group
	cancel func()
	stdin  io.Writer
	stdout *bufio.Reader
	id     int64
	msg    chan Message
	not    chan Notification
}

// NewCmd executes the specified langserver and configures a client for
// it at the specified Root URI. The langserver must interact over its
// stdin/stdout.
func NewCmd(rootURI string, command string, args ...string) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	l := &Client{
		cancel: cancel,
		msg:    make(chan Message),
		not:    make(chan Notification),
	}
	cmd := exec.CommandContext(ctx, command, args...)
	l.grp, ctx = errgroup.WithContext(ctx)
	getRW := func() (io.Reader, io.Writer) {
		pr, pw := io.Pipe()
		l.grp.Go(func() error {
			<-ctx.Done()
			pr.Close()
			pw.Close()
			return nil
		})
		return pr, pw
	}
	cmd.Stdin, l.stdin = getRW()
	cmd.Stderr = os.Stderr
	{
		pr, pw := getRW()
		cmd.Stdout = pw
		l.stdout = bufio.NewReader(pr)
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if err := func() error {
		var res struct {
			ID     ID
			Result protocol.InitializeResult
		}
		id, err := l.request("initialize", protocol.InitializeParams{
			InnerInitializeParams: protocol.InnerInitializeParams{
				RootURI: rootURI,
			},
		})
		if err != nil {
			return err
		}
		b, err := l.readMessage()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &res); err != nil {
			return err
		}
		if res.ID != id {
			return errors.New("unexpected id")
		}
		l.Capabilities = res.Result.Capabilities
		return nil
	}(); err != nil {
		l.Close()
		return nil, err
	}
	l.grp.Go(cmd.Wait)
	l.grp.Go(func() error {
		for {
			resp, err := l.readMessage()
			if err != nil {
				return err
			}
			var m Message
			var n Notification
			if err := json.Unmarshal(resp, &m); err != nil {
				return err
			} else if m.ID != 0 {
				go func() {
					l.msg <- m
				}()
				continue
			}
			if err := json.Unmarshal(resp, &n); err != nil {
				return err
			}
			go func() {
				l.not <- n
			}()
		}
	})
	return l, nil
}

func (l *Client) Message() <-chan Message {
	return l.msg
}

func (l *Client) Notification() <-chan Notification {
	return l.not
}

func (l *Client) Close() {
	l.cancel()
	l.grp.Wait()
}

func (l *Client) readMessage() ([]byte, error) {
	var buf bytes.Buffer
	var n int
	for {
		line, err := l.stdout.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if !strings.HasSuffix(line, "\r\n") {
			buf.WriteString(line)
			continue
		}
		buf.WriteString(line[:len(line)-2])
		line = buf.String()
		if line == "" {
			break
		}
		sp := strings.SplitN(line, ": ", 2)
		if len(sp) != 2 {
			return nil, fmt.Errorf("unknown header: %s", line)
		}
		switch sp[0] {
		case "Content-Length":
			n, err = strconv.Atoi(sp[1])
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown header: %s", line)
		}
		buf.Reset()
	}
	if n == 0 {
		return nil, fmt.Errorf("no Content-Length")
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(l.stdout, b); err != nil {
		return nil, err
	}
	print("READ", b)
	return b, nil
}

func print(prefix string, msg []byte) {
	var out bytes.Buffer
	prefix += "\t"
	json.Indent(&out, msg, prefix, "\t")
	fmt.Print(prefix, out.String(), "\n\n")
}

type Message struct {
	Error *struct {
		Code    int
		Message string
	} `json:,omitempty`
	ID     ID
	Result json.RawMessage
}

type Notification struct {
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

type requestNotification struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type requestMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      ID          `json:"id"`
}

type ID int64

func (l *Client) makeID() ID {
	return ID(atomic.AddInt64(&l.id, 1))
}

func (l *Client) send(msg interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Content-Length: %d\r\n", len(b))
	buf.Write([]byte("\r\n"))
	buf.Write(b)
	d := buf.Bytes()
	print("SEND", b)
	_, err = l.stdin.Write(d)
	return err
}

const vers = "2.0"

func (l *Client) request(method string, params interface{}) (ID, error) {
	msg := requestMessage{
		JSONRPC: vers,
		ID:      l.makeID(),
		Method:  method,
		Params:  params,
	}
	return msg.ID, l.send(msg)
}

func (l *Client) notification(method string, params interface{}) error {
	msg := requestNotification{
		JSONRPC: vers,
		Method:  method,
		Params:  params,
	}
	return l.send(msg)
}

func (l *Client) DidOpen(doc protocol.TextDocumentItem) error {
	return l.notification("textDocument/didOpen", protocol.DidOpenTextDocumentParams{
		TextDocument: doc,
	})
}

func (l *Client) DidClose(uri string) error {
	return l.notification("textDocument/didClose", protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
}

func (l *Client) DidChange(doc protocol.DidChangeTextDocumentParams) error {
	return l.notification("textDocument/didChange", doc)
}

func (l *Client) Definition(pos protocol.TextDocumentPositionParams) (ID, error) {
	return l.request("textDocument/definition", pos)
}

func (l *Client) Symbols(uri string) (ID, error) {
	return l.request("textDocument/documentSymbol", protocol.DocumentSymbolParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
}

func (l *Client) Hover(pos protocol.TextDocumentPositionParams) (ID, error) {
	return l.request("textDocument/hover", pos)
}

func (l *Client) Format(uri string) (ID, error) {
	return l.request("textDocument/formatting", protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
}
