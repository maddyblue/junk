package lsp

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/mjibson/acmepls/lsp-client/protocol"
)

func TestGopls(t *testing.T) {
	defer leaktest.AfterTest(t)()

	l, err := NewCmd("file://", "gopls")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	uri := "file://" + filepath.Join(cwd, "test.go")
	tdi := protocol.TextDocumentIdentifier{URI: uri}

	if _, err := l.Hover(protocol.TextDocumentPositionParams{
		TextDocument: tdi,
		Position: protocol.Position{
			Line:      9,
			Character: 7,
		},
	}); err != nil {
		t.Fatal(err)
	}

	if err := l.DidOpen(protocol.TextDocumentItem{
		URI:     uri,
		Version: 1,
	}); err != nil {
		t.Fatal(err)
	}

	if err := l.DidChange(protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: tdi,
			Version:                2,
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Text: v1,
			},
		},
	}); err != nil {
		t.Fatal(err)
	}

	select {}
	return

	for line := 0; line < 10; line++ {
		for chr := 0; chr <= 6; chr += 2 {
			l.Hover(protocol.TextDocumentPositionParams{
				TextDocument: tdi,
				Position: protocol.Position{
					Line:      float64(line),
					Character: float64(chr),
				},
			})
			select {
			case m := <-l.Message():
				fmt.Println("M", m)
			case n := <-l.Notification():
				fmt.Println("N", n)
			}
		}
	}

	select {}
	return
}

const v1 = `package lsp

import "fmt"

func main() {
	fmt.Println("test")
}
`

const v2 = `
package lsp

import "fmt"

func main() {


	fmt.Println("test")
}
`
