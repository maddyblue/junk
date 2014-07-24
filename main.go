/*
 * Copyright (c) 2014 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"bytes"
	"flag"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"code.google.com/p/go.tools/astutil"
)

const dirMode os.FileMode = 0755

var (
	dryrun     = flag.Bool("n", false, "don't perform any action, instead print them")
	create     = flag.Bool("c", false, "create the third party directory if needed")
	relative   = flag.Bool("r", false, "use a relative third party directory (needed on App Engine)")
	verbose    = flag.Bool("v", false, "print actions")
	thirdParty = flag.String("d", "_third_party", "name of third party directory")
	flagUpdate = flag.Bool("u", false, "update (go get -d -u) used packages")

	relpath, gopath, ThirdParty string
)

func main() {
	flag.Parse()
	if *dryrun {
		*verbose = true
	}

	ThirdParty = filepath.Base(*thirdParty)
	if ThirdParty != *thirdParty {
		log.Fatal("third party directory cannot contain path separators")
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
	var importSites []string
	for _, s := range filepath.SplitList(os.Getenv("GOPATH")) {
		s, err := filepath.Abs(s)
		if err != nil {
			log.Fatal(err)
		}
		s, err = filepath.EvalSymlinks(s)
		if err != nil {
			log.Fatal(err)
		}
		if sf, err := os.Open(filepath.Join(s, "src")); err == nil {
			names, err := sf.Readdirnames(0)
			if err == nil {
				for _, name := range names {
					importSites = append(importSites, name)
				}
			}
		}
		if !strings.HasPrefix(pwd, s) {
			continue
		}
		gopath = s
		relpath = strings.TrimPrefix(pwd, filepath.Join(s, "src")+string(os.PathSeparator))
		break
	}
	if relpath == "" {
		log.Fatal("couldn't determine current directory relative to $GOPATH")
	}

	if *verbose {
		log.Println("using GOPATH:", gopath)
		log.Println("using relative path:", relpath)
	}
	for {
		update()
		if !rewriteImports(importSites) {
			break
		}
	}
}

func update() {
	paths := make(map[string]struct{})
	reltp := ThirdParty + "/"
	if !*relative {
		reltp = filepath.ToSlash(relpath) + "/" + reltp
	}
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
		if *flagUpdate {
			if *verbose {
				log.Printf("go get -d -u %v", k)
			}
			if !*dryrun {
				out, err := exec.Command("go", "get", "-d", "-u", k).CombinedOutput()
				if len(out) > 0 {
					log.Println(string(out))
				}
				if err != nil {
					log.Println("go get", k, "err:", err)
				}
			}
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

func rewriteImports(importSites []string) (rewritten bool) {
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
			for _, site := range importSites {
				v, err := strconv.Unquote(im.Path.Value)
				if err != nil {
					log.Fatal(err)
				}
				// don't replace ourself
				if strings.HasPrefix(v, relpath+"/") {
					continue
				}
				if strings.HasPrefix(v, site+"/") {
					target := ThirdParty + "/" + v
					if !*relative {
						target = relpath + "/" + target
					}
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
