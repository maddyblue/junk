package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/go.tools/astutil"
)

const ThirdParty = "third_party"
const dirMode os.FileMode = 0755

var (
	dryrun  = flag.Bool("n", false, "don't perform any action, instead print them")
	create  = flag.Bool("c", false, "create the "+ThirdParty+" directory if needed")
	verbose = flag.Bool("v", false, "print actions")

	relpath, gopath string
)

func main() {
	flag.Parse()
	if *dryrun {
		*verbose = true
	}

	if *create {
		os.Mkdir(ThirdParty, dirMode)
	}
	if _, err := os.Open(ThirdParty); err != nil {
		log.Fatalf("could not open %s directory: run with -c or create it manually", ThirdParty)
	}

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
		gopath = s
		relpath = strings.TrimPrefix(pwd, filepath.Join(s, "src")+string(os.PathSeparator))
	}
	if relpath == "" {
		log.Fatal("couldn't determine current directory relative to $GOPATH")
	}

	for {
		update()
		if !rewriteImports() {
			break
		}
	}
}

var ImportSites = []string{
	"bazil.org",
	"code.google.com",
	"github.com",
	"labix.org",
	"launchpad.net",
}

func update() {
	paths := make(map[string]struct{})
	reltp := filepath.ToSlash(relpath) + "/" + ThirdParty + "/"
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}
		for _, im := range f.Imports {
			v := im.Path.Value[1 : len(im.Path.Value)-1]
			if !strings.HasPrefix(v, reltp) {
				continue
			}
			paths[strings.TrimPrefix(v, reltp)] = struct{}{}
		}
		return nil
	})
	for k := range paths {
		fpath := filepath.Join(gopath, "src", k)
		f, err := os.Open(fpath)
		if err != nil {
			log.Printf("%s required, but could not be found at %s", k, fpath)
			continue
		}
		files, err := f.Readdir(0)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		for _, source := range files {
			if source.IsDir() || strings.HasPrefix(source.Name(), ".") {
				continue
			}
			spath := filepath.Join(f.Name(), source.Name())
			destdir := filepath.Join(ThirdParty, k)
			dest := filepath.Join(destdir, source.Name())
			if err := os.MkdirAll(destdir, dirMode); err != nil {
				log.Fatal(err)
			}
			dst, err := os.Open(dest)
			created := false
			newer := false
			if os.IsNotExist(err) {
				if *verbose {
					log.Println("create", dest)
				}
				created = true
			} else if err != nil {
				log.Fatal(err)
			} else {
				destination, err := dst.Stat()
				if err != nil {
					log.Fatal(err)
				}
				newer = source.ModTime().After(destination.ModTime())
				dst.Close()
			}
			if created || newer {
				if *verbose {
					log.Printf("copy %s -> %s\n", spath, dest)
				}
				if !*dryrun {
					src, err := os.Open(spath)
					if err != nil {
						log.Fatal(err)
					}
					dst, err = os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, source.Mode())
					if err != nil {
						log.Fatal(err)
					}
					if _, err := io.Copy(dst, src); err != nil {
						log.Fatal(err)
					}
					src.Close()
					dst.Close()
				}
			}
		}
	}
}

func rewriteImports() (rewritten bool) {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
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
					target := fmt.Sprintf("%s/%s/%s", relpath, ThirdParty, v)
					if *verbose {
						log.Printf("rewrite %s: %s -> %s\n", path, v, target)
					}
					if *dryrun {
						break
					}
					if !astutil.RewriteImport(fset, f, v, target) {
						log.Fatalf("%s: could not rewrite %s", path, im.Path.Value)
					}
					rewritten = true
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
	return
}
