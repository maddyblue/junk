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

/*
Party is a tool to automate the use of third party packages.

Many existing Go package managers provide solutions to the problem of relying on
other projects' code. Go handles the use of other packages elegantly, but is
prone to problems when that code has breaking changes or is removed from the
Internet. As a solution, the FAQ suggests copying any third party repositories
into your own project and changing your import paths to match. This is a
recursive process, as you must then perform the same procedure for those copied
packages and so on, until all leaf packages have no external imports.

party is a tool that performs all of this work. It:

1. Rewrites all external imports in both your project and the third party
directory from "other.com/user/package" to
"host.com/user/your_project/_third_party/other.com/user/package".

2. Updates or copies files from "$GOPATH/src/other.com/user/package" to
"$GOPATH/src/host.com/user/your_project/_third_party/other.com/user/package".

3. Returns to step 1 until no more rewrites are performed.

Usage

From your project's root directory, for the first run:

	$ party -c

Subsequent invocations can omit -c.

The flags are:
	-c
		create the third party directory if it does not exist
	-u
		run "go get -d -u <pkg>" on packages imported by party for the
		current package
	-d="_third_party"
		set the third party directory
	-v
		enable verbose output
	-n
		dry run: actions are printed instead of run
	-r
		import third party packages as relative, local imports

App Engine and Relative Imports

When working on an App Engine project, use the -r flag. This will cause all
third party imports to be local packages. This is needed because otherwise App
Engine will run init() routines from your .go files twice if your app is also in
your $GOPATH.

Bugs

There is no guarantee that the source directory exists. For example, a package
may have a OS-specific file (file_windows.go) which imports another package. On
non-Windows, that other package will not have been fetched during "go get" and
will thus not exist. These packages must be fetched using "go get" or some other
method by hand. Or, running party on a Windows machine (in this case) will
correctly copy those files. These cases are reported when running party.

party performs the equivalent of gofmt on any files affected by the import
rewrite, thus the diff may be somewhat larger than just the import lines.

Third party imports that use local (relative) packages will fail, as those
packages are assumed to be part of the standard library. Those packages will not
be copied.

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
						b, err = fixImportCheck(b, path.Join(relpath, destdir))
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
	if bytes.IndexByte(body[f.Package:pos], '\n') == -1 {
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
