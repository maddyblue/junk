// Modified from: https://github.com/fhs/acme-lsp/blob/master/cmd/L/acme.go

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"9fans.net/go/acme"
	"9fans.net/go/plan9"
	"9fans.net/go/plumb"
	"github.com/mjibson/acmepls/lsp-client/protocol"
	"github.com/pkg/errors"
)

func open(b []byte) error {
	port, err := plumb.Open("send", plan9.OWRITE)
	if err != nil {
		return err
	}
	defer port.Close()
	msg := &plumb.Message{
		Dst:  "edit",
		Type: "text",
		Data: b,
	}
	return msg.Send(port)
}

type win struct {
	*acme.Win
}

func newWin() (*win, error) {
	w, err := acme.New()
	if err != nil {
		return nil, err
	}
	return &win{w}, err
}

func openWin(id int) (*win, error) {
	w, err := acme.Open(id, nil)
	if err != nil {
		return nil, err
	}
	return &win{w}, err
}

func openCurrentWin() (*win, error) {
	id, err := strconv.Atoi(os.Getenv("winid"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse $winid")
	}
	return openWin(id)
}

func (w *win) Filename() (string, error) {
	tag, err := w.ReadAll("tag")
	if err != nil {
		return "", err
	}
	i := bytes.IndexRune(tag, ' ')
	if i < 0 {
		i = len(tag)
	}
	return string(tag[:i]), nil
}

func (w *win) DocumentURI() (uri, fname string, err error) {
	fname, err = w.Filename()
	if err != nil {
		return "", "", err
	}
	return filenameToURI(fname), fname, nil
}

// ReadDotAddr returns the address of current selection.
func (w *win) ReadDotAddr() (q0, q1 int, err error) {
	_, _, err = w.ReadAddr() // open addr file
	if err != nil {
		return 0, 0, err
	}
	err = w.Ctl("addr=dot")
	if err != nil {
		return 0, 0, err
	}
	return w.ReadAddr()
}

func (w *win) Position() (*protocol.TextDocumentPositionParams, string, error) {
	fname, err := w.Filename()
	if err != nil {
		return nil, "", err
	}
	q0, _, err := w.ReadDotAddr()
	if err != nil {
		return nil, "", err
	}
	off, err := getNewlineOffsets(w.FileReadWriter("body"))
	if err != nil {
		return nil, "", err
	}
	line, col := off.OffsetToLine(q0)
	return &protocol.TextDocumentPositionParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: filenameToURI(fname),
		},
		Position: protocol.Position{
			Line:      float64(line),
			Character: float64(col),
		},
	}, fname, nil
}

func (w *win) DataURI() (string, error) {
	body, err := w.ReadAll("body")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("data:text/plain;base64,%s",
		base64.StdEncoding.EncodeToString(body),
	), nil
}

func (w *win) FileReadWriter(filename string) io.ReadWriter {
	return &winFile{
		w:    w.Win,
		name: filename,
	}
}

type winFile struct {
	w    *acme.Win
	name string
}

func (f *winFile) Read(b []byte) (int, error) {
	return f.w.Read(f.name, b)
}

func (f *winFile) Write(b []byte) (int, error) {
	return f.w.Write(f.name, b)
}

func (w *win) Edit(edits []protocol.TextEdit) error {
	if len(edits) == 0 {
		return nil
	}
	body, err := w.ReadAll("body")
	if err != nil {
		return err
	}
	if len(edits) == 1 {
		if string(body) == edits[0].NewText {
			return nil
		}
	}
	sort.Slice(edits, func(i, j int) bool {
		pi := edits[i].Range.Start
		pj := edits[j].Range.Start
		if pi.Line == pj.Line {
			return pi.Character < pj.Character
		}
		return pi.Line < pj.Line
	})
	_, err = w.Seek("body", 0, 0)
	if err != nil {
		return errors.Wrapf(err, "seed failed for window %v", w.ID())
	}
	off, err := getNewlineOffsets(bytes.NewReader(body))
	if err != nil {
		return errors.Wrapf(err, "failed to obtain newline offsets for window %v", w.ID())
	}

	w.Ctl("nomark")
	w.Ctl("mark")

	delta := 0
	for _, e := range edits {
		soff := off.LineToOffset(int(e.Range.Start.Line), int(e.Range.Start.Character))
		eoff := off.LineToOffset(int(e.Range.End.Line), int(e.Range.End.Character))
		err := w.Addr("#%d,#%d", soff+delta, eoff+delta)
		if err != nil {
			return errors.Wrapf(err, "failed to write to addr for winid=%v", w.ID())
		}
		_, err = w.Write("data", []byte(e.NewText))
		if err != nil {
			return errors.Wrapf(err, "failed to write new text to data file")
		}
		delta += len(e.NewText) - (eoff - soff)
	}
	return nil
}

type editor interface {
	Edit(edits []protocol.TextEdit) error
}

func uriToFilename(uri string) string {
	return strings.TrimPrefix(string(uri), "file://")
}

func filenameToURI(fname string) string {
	return "file://" + fname
}

/*
func applyAcmeEdits(we *lsp.WorkspaceEdit) error {
	wins, err := acme.Windows()
	if err != nil {
		return errors.Wrapf(err, "failed to read list of acme index")
	}
	winid := make(map[string]int, len(wins))
	for _, info := range wins {
		winid[info.Name] = info.ID
	}

	for uri := range we.Changes {
		fname := uriToFilename(uri)
		if _, ok := winid[fname]; !ok {
			return fmt.Errorf("%v: not open in acme", fname)
		}
	}
	for uri, edits := range we.Changes {
		fname := uriToFilename(lsp.DocumentURI(uri))
		id := winid[fname]
		w, err := openWin(id)
		if err != nil {
			return errors.Wrapf(err, "failed to open window %v", id)
		}
		if err := w.Edit(edits); err != nil {
			return errors.Wrapf(err, "failed to apply edits to window %v", id)
		}
		w.CloseFiles()
	}
	return nil
}
*/
