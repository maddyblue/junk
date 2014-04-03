package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/go.tools/astutil"
)

var dryrun = flag.Bool("n", false, "don't perform any action, instead print them")
var relpath string

func main() {
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range filepath.SplitList(os.Getenv("GOPATH")) {
		s, err := filepath.Abs(s)
		if err != nil {
			log.Fatal(err)
		}
		if !strings.HasPrefix(pwd, s) {
			continue
		}
		relpath = strings.TrimPrefix(pwd, filepath.Join(s, "src")+string(os.PathSeparator))
	}
	if relpath == "" {
		log.Fatal("couldn't determine current directory relative to $GOPATH")
	}

	rewriteImports()
}

var ImportSites = []string{
	"bazil.org",
	"code.google.com",
	"github.com",
	"labix.org",
	"launchpad.net",
}

func rewriteImports() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}
		for _, im := range f.Imports {
			for _, site := range ImportSites {
				v := im.Path.Value[1 : len(im.Path.Value)-1]
				// don't replace ourself
				if strings.HasPrefix(v, relpath+"/") {
					continue
				}
				if strings.HasPrefix(v, site+"/") {
					target := fmt.Sprintf("%s/third_party/%s", relpath, v)
					if *dryrun {
						fmt.Printf("%s: %s -> %s\n", path, v, target)
						break
					}
					if !astutil.RewriteImport(fset, f, v, target) {
						log.Fatalf("%s: could not rewrite %s", path, im.Path.Value)
					}
					break
				}
			}
		}
		var buf bytes.Buffer
		if err := printer.Fprint(&buf, fset, f); err != nil {
			log.Fatal(err)
		}
		fb, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(path, fb, info.Mode()); err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
