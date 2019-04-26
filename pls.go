package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"

	"9fans.net/go/acme"
	"github.com/golang-commonmark/markdown"
	"github.com/gorilla/websocket"
	lsp "github.com/mjibson/acmepls/lsp-client"
	"github.com/mjibson/acmepls/lsp-client/protocol"
)

type pls struct {
	lock     sync.Mutex
	m        map[*regexp.Regexp]*lsp.Client
	version  int64
	prevBody map[string]string
	prevMsg  chan interface{}
	ids      map[lsp.ID]func(json.RawMessage)
}

func NewPLS() *pls {
	p := pls{
		m:        map[*regexp.Regexp]*lsp.Client{},
		prevBody: map[string]string{},
		ids:      map[lsp.ID]func(json.RawMessage){},
	}
	go p.watch()
	return &p
}

func (p *pls) SendMsg(typ string, msg interface{}) {
	p.lock.Lock()
	if p.prevMsg != nil {
		p.prevMsg <- struct {
			Typ string
			Msg interface{}
		}{
			Typ: typ,
			Msg: msg,
		}
	}
	p.lock.Unlock()
}

func (p *pls) AddLSP(suffix, cmd string, args ...string) error {
	client, err := lsp.NewCmd(*flagDev, "file://", cmd, args...)
	if err != nil {
		return err
	}
	p.m[regexp.MustCompile(fmt.Sprintf(`\.%s`, suffix))] = client
	go func() {
		for {
			select {
			case m := <-client.Message():
				if m.Error != nil {
					log.Println(*m.Error)
					//p.SendMsg("error", *m.Error)
					continue
				}
				p.lock.Lock()
				fn, ok := p.ids[m.ID]
				p.lock.Unlock()
				if !ok {
					continue
				}
				fn(m.Result)
			}
		}
	}()
	return nil
}

func (p *pls) Command(r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	win, err := openWin(id)
	if err != nil {
		return nil, err
	}
	pos, reqName, err := win.Position()
	if err != nil {
		return nil, err
	}
	cl := p.getClient(reqName)
	if cl == nil {
		return nil, fmt.Errorf("no client")
	}

	if err := p.syncOpen(); err != nil {
		return nil, err
	}

	var cmdID lsp.ID
	method := r.FormValue("method")
	switch method {
	case "completion":
		p.Completion(cl, *pos)
	case "definition":
		cmdID, err = cl.Definition(*pos)
		p.RegisterCmd(cmdID, func(msg json.RawMessage) {
			var locs []protocol.Location
			if err := json.Unmarshal(msg, &locs); err != nil {
				return
			}
			if len(locs) == 0 {
				return
			}
			loc := locs[0]
			addr := fmt.Sprintf("%s:%v:%v", uriToFilename(loc.URI), loc.Range.Start.Line+1, loc.Range.Start.Character+1)
			open([]byte(addr))
			p.SendMsg(method, addr)
		})
	case "hover":
		p.Hover(cl, *pos)
	case "signature":
		cmdID, err = cl.Signature(*pos)
		p.RegisterCmd(cmdID, func(msg json.RawMessage) {
			var sig protocol.SignatureHelp
			if err := json.Unmarshal(msg, &sig); err != nil {
				return
			}
			p.SendMsg(method, struct {
				Filename  string
				Signature protocol.SignatureHelp
			}{
				Filename:  reqName,
				Signature: sig,
			})
		})
	case "symbols":
		cmdID, err = cl.Symbols(pos.TextDocument.URI)
		p.RegisterCmd(cmdID, func(msg json.RawMessage) {
			var symbs []protocol.DocumentSymbol
			if err := json.Unmarshal(msg, &symbs); err != nil {
				return
			}
			p.SendMsg(method, struct {
				Filename string
				Symbols  []protocol.DocumentSymbol
			}{
				Filename: reqName,
				Symbols:  symbs,
			})
		})
	default:
		return nil, fmt.Errorf("unknown method")
	}
	return nil, err
}

func (p *pls) Completion(cl *lsp.Client, pos protocol.TextDocumentPositionParams) {
	cmdID, err := cl.Completion(pos)
	if err != nil {
		log.Println(err)
		return
	}
	p.RegisterCmd(cmdID, func(msg json.RawMessage) {
		var comps protocol.CompletionList
		if err := json.Unmarshal(msg, &comps); err != nil {
			log.Println(err)
			return
		}
		if len(comps.Items) > 5 {
			comps.Items = comps.Items[:5]
		}
		p.SendMsg("completion", struct {
			Filename   string
			Completion protocol.CompletionList
		}{
			Filename:   uriToFilename(pos.TextDocument.URI),
			Completion: comps,
		})
	})
}

func (p *pls) Hover(cl *lsp.Client, pos protocol.TextDocumentPositionParams) {
	cmdID, err := cl.Hover(pos)
	if err != nil {
		log.Println(err)
		return
	}
	p.RegisterCmd(cmdID, func(msg json.RawMessage) {
		var hov protocol.Hover
		if err := json.Unmarshal(msg, &hov); err != nil {
			return
		}
		res := markdown.New().RenderToString([]byte(hov.Contents.Value))
		p.SendMsg("hover", struct {
			Filename string
			HTML     string
		}{
			Filename: uriToFilename(pos.TextDocument.URI),
			HTML:     res,
		})
	})
}

func (p *pls) RegisterCmd(id lsp.ID, fn func(msg json.RawMessage)) {
	p.lock.Lock()
	p.ids[id] = fn
	p.lock.Unlock()
}

func (p *pls) syncOpen() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	wins, err := acme.Windows()
	if err != nil {
		panic(err)
		return err
	}
	// Keep track of previously open docs so we know what to close.
	prevOpen := map[string]bool{}
	for name := range p.prevBody {
		prevOpen[name] = true
	}
	for _, w := range wins {
		w, err := openWin(w.ID)
		if err != nil {
			panic(err)
			return err
		}
		filename, err := w.Filename()
		if err != nil {
			panic(err)
			return err
		}
		delete(prevOpen, filename)
		cl := p.getClient(filename)
		if cl == nil {
			continue
		}
		body, err := w.ReadAll("body")
		if err != nil {
			panic(err)
			return err
		}
		text := string(body)
		version := float64(atomic.AddInt64(&p.version, 1))
		uri := filenameToURI(filename)
		if previous, ok := p.prevBody[filename]; !ok {
			err = cl.DidOpen(protocol.TextDocumentItem{
				URI:     uri,
				Version: version,
				Text:    text,
			})
		} else if text != previous {
			err = cl.DidChange(protocol.DidChangeTextDocumentParams{
				TextDocument: protocol.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: uri},
					Version:                version,
				},
				ContentChanges: []protocol.TextDocumentContentChangeEvent{
					protocol.TextDocumentContentChangeEvent{
						Text: text,
					},
				},
			})
		}
		p.prevBody[filename] = text
		if err != nil {
			panic(err)
			return err
		}
	}
	for name := range prevOpen {
		cl := p.getClient(name)
		if cl == nil {
			continue
		}
		cl.DidClose(name)
	}
	return nil
}

func (p *pls) WS(ws *websocket.Conn) {
	defer ws.Close()
	msgs := make(chan interface{})
	p.lock.Lock()
	if p.prevMsg != nil {
		close(p.prevMsg)
	}
	p.prevMsg = msgs
	p.lock.Unlock()
	go func() {
		p.SendMsg("state", p.getState())
	}()
	for msg := range msgs {
		if err := ws.WriteJSON(msg); err != nil {
			log.Println(err)
			return
		}
	}
	log.Println("closing ws")
}

func (p *pls) getClient(name string) *lsp.Client {
	for re, client := range p.m {
		if re.MatchString(name) {
			return client
		}
	}
	return nil
}

func (p *pls) getState() interface{} {
	wins, err := acme.Windows()
	if err != nil {
		log.Println(err)
		return nil
	}
	type Win struct {
		Info         acme.WinInfo
		Methods      []string
		Capabilities protocol.ServerCapabilities
	}
	ret := struct {
		Wins []Win
	}{
		Wins: []Win{},
	}
	for _, win := range wins {
		cl := p.getClient(win.Name)
		if cl == nil {
			continue
		}
		var methods []string
		c := cl.Capabilities
		if c.DefinitionProvider {
			methods = append(methods, "definition")
		}
		if c.HoverProvider {
			//methods = append(methods, "hover")
		}
		if c.DocumentSymbolProvider {
			methods = append(methods, "symbols")
		}
		if c.CompletionProvider != nil {
			methods = append(methods, "completion")
		}
		if c.SignatureHelpProvider != nil {
			//methods = append(methods, "signature")
		}
		if len(methods) == 0 {
			continue
		}
		ret.Wins = append(ret.Wins, Win{
			Info:         win,
			Methods:      methods,
			Capabilities: c,
		})
	}
	return ret
}

func (p *pls) watch() {
	defer func() {
		panic("no return")
	}()
	l, err := acme.Log()
	if err != nil {
		panic(err)
	}

	for {
		event, err := l.Read()
		if err != nil {
			panic(err)
		}
		fmt.Println("EVENT", event)
		win, _ := openWin(event.ID)
		uri := filenameToURI(event.Name)
		cl := p.getClient(event.Name)
		if cl == nil || win == nil {
			continue
		}
		switch event.Op {
		case "new", "del", "focus":
			if err := p.syncOpen(); err != nil {
				log.Println("SO ERR", err)
				continue
			}
			if event.Op == "del" {
				p.lock.Lock()
				delete(p.prevBody, event.Name)
				p.lock.Unlock()

				if cl != nil {
					cl.DidClose(uri)
				}
			}
			p.SendMsg("state", p.getState())
			if event.Op == "focus" {
				pos, _, _ := win.Position()
				p.Hover(cl, *pos)
			}
		case "put":
			if err := p.syncOpen(); err != nil {
				log.Println("SO ERR", err)
				continue
			}
			id, err := cl.CodeAction(uri)
			if err != nil {
				log.Println(err)
				continue
			}
			// Run organize imports first.
			p.RegisterCmd(id, func(msg json.RawMessage) {
				var actions []protocol.CodeAction
				if err := json.Unmarshal(msg, &actions); err != nil {
					return
				}
				changed := false
				for _, action := range actions {
					if action.Edit == nil || action.Edit.Changes == nil {
						continue
					}
					if edits := (*action.Edit.Changes)[uri]; len(edits) > 0 {
						win.Edit(edits)
						changed = true
					}
				}
				if changed {
					if err := p.syncOpen(); err != nil {
						log.Println("SO ERR", err)
						return
					}
				}

				// Now fmt.
				id, err := cl.Format(uri)
				if err != nil {
					log.Println(err)
					return
				}
				p.RegisterCmd(id, func(msg json.RawMessage) {
					var edits []protocol.TextEdit
					if err := json.Unmarshal(msg, &edits); err != nil {
						return
					}
					win.Edit(edits)
				})
			})
		}
	}
}
