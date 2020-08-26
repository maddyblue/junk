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
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

const dirMode os.FileMode = 0755

var (
	dryrun       = flag.Bool("n", false, "don't perform any action, instead print them")
	create       = flag.Bool("c", false, "create the third party directory if needed")
	relative     = flag.Bool("r", false, "use a relative third party directory (needed on App Engine)")
	verbose      = flag.Bool("v", false, "print actions")
	thirdParty   = flag.String("d", "_third_party", "name of third party directory")
	flagUpdate   = flag.Bool("u", false, "update (go get -d -u) used packages")
	includeTests = flag.Bool("t", false, "import dependencies for _test.go files")

	relpath, ThirdParty string
	gopath              []string
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
	// Under certain circumstances Getwd and EvalSymlinks return different case drive letter on Windows.
	// To ensure a match in the loop below, do an Evalsymlink on pwd too.
	pwd, err = filepath.EvalSymlinks(pwd)
	if err != nil {
		log.Fatal(err)
	}
	var importSites []string
	gopath = []string{}
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
			gopath = append(gopath, s)
			continue
		}
		gopath = append([]string{s}, gopath...)
		relpath = filepath.ToSlash(strings.TrimPrefix(pwd, filepath.Join(s, "src")+string(os.PathSeparator)))
		break
	}
	if relpath == "" {
		log.Fatal("couldn't determine current directory relative to $GOPATH")
	}

	if *verbose {
		log.Println("using relative path:", relpath)
	}
	updated := make(map[string]bool)
	for {
		update(updated)
		if !rewriteImports(importSites) {
			break
		}
	}
}

func update(updated map[string]bool) {
	paths := make(map[string]struct{})
	reltp := ThirdParty + "/"
	if !*relative {
		reltp = filepath.ToSlash(relpath) + "/" + reltp
	}
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !*includeTests && strings.HasSuffix(path, "_test.go") {
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
		fpath, err := findPackage(k)
		if err != nil {
			log.Println(err)
			continue
		}
		if *flagUpdate && !updated[k] {
			updated[k] = true
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
		f, err := os.Open(fpath)
		if err != nil {
			log.Fatal(err)
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
					b, err := ioutil.ReadFile(spath)
					if err != nil {
						log.Fatal(err)
					}
					if filepath.Ext(spath) == ".go" {
						b, err = fixImportCheck(b, filepath.ToSlash(path.Join(relpath, destdir)))
						if err != nil {
							log.Fatal(err)
						}
					}
					if err := ioutil.WriteFile(dest, b, source.Mode()); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
}

func findPackage(name string) (string, error) {
	for _, gp := range gopath {
		fpath := filepath.Join(gp, "src", name)
		if _, err := os.Stat(fpath); err == nil {
			return fpath, nil
		}
	}
	return "", fmt.Errorf("%s required, but could not be found", name)
}

func fixImportCheck(body []byte, importPath string) ([]byte, error) {
	fset := token.NewFileSet()
	// todo: see if we can restrict the mode some more
	f, err := parser.ParseFile(fset, "", body, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var after *ast.CommentGroup
	var pos token.Pos = token.Pos(len(body))
	for _, v := range f.Comments {
		text := strings.TrimSpace(v.Text())
		if v.Pos() > f.Package && v.Pos() < pos && strings.HasPrefix(text, "import") {
			pos = v.Pos()
			after = v
		}
	}
	if after != nil && bytes.IndexByte(body[f.Package:pos], '\n') == -1 {
		comment := fmt.Sprintf(`// import "%s"`, importPath)
		buf := new(bytes.Buffer)
		buf.Write(body[:after.Pos()-1])
		buf.WriteString(comment)
		buf.Write(body[after.End()-1:])
		body = buf.Bytes()
	}
	return body, nil
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
				if v == relpath || strings.HasPrefix(v, relpath+"/") {
					if *verbose {
						log.Printf("skipping %s because it is in %s\n", v, relpath)
					}
					continue
				}
				if strings.HasPrefix(v, site+"/") {
					target := ThirdParty + "/" + v
					if !*relative {
						target = filepath.ToSlash(relpath) + "/" + target
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
