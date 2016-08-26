package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"
)

func main() {
	dir, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	files, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	var rs []Recipe
	for _, f := range files {
		p := filepath.Join(dir.Name(), f.Name())
		d, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(d), "\n")
		r := Recipe{}
		dirs := false
		for i, line := range lines {
			if i < 2 {
				continue
			}
			line = strings.TrimSpace(line)
			if i == 2 {
				r.Name = strings.Title(strings.ToLower(line))
			} else if i == 3 {
				m := summaryRE.FindStringSubmatch(line)
				if len(m) == 5 {
					r.By = strings.TrimSpace(m[1])
					r.Size = strings.TrimSpace(m[2])
					r.Prep = strings.TrimSpace(m[3])
					r.Categories = strings.TrimSpace(m[4])
					if r.Prep == "0:00" {
						r.Prep = ""
					}
				}
			}
			if i < 5 || strings.Count(line, "-") > 10 {
				continue
			}
			if !dirs && unicode.IsUpper(rune(line[0])) {
				dirs = true
			}
			if dirs {
				r.Directions += line + "\n"
			} else {
				m := ingredRE.FindAllStringSubmatch(line, -1)
				if len(m) > 0 {
					for _, ingred := range m {
						r.Ingreds = append(r.Ingreds, strings.TrimSpace(ingred[0]))
					}
				} else {
					r.Directions += line + "\n"
				}
			}
		}
		r.Directions = strings.TrimSpace(r.Directions)
		rs = append(rs, r)
	}
	fb := bytes.NewBufferString(xml.Header)
	if err := ftmpl.Execute(fb, &rs); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("feed.xml", fb.Bytes(), 0666)
	fmt.Println("done")
}

var summaryRE = regexp.MustCompile("Recipe By : (.*)Serving Size : (.*)Preparation Time : (.*)Categories : (.*)")
var ingredRE = regexp.MustCompile(`(\d+)(/\d+)? (\D+)`)

type Recipe struct {
	Name       string
	By         string
	Size       string
	Prep       string
	Categories string
	Directions string
	Ingreds    []string
}

var ftmpl, rtmpl *template.Template

func init() {
	var err error
	rtmpl, err = template.New("r").Parse(RTMPL)
	if err != nil {
		log.Fatal(err)
	}
	ftmpl, err = template.New("f").Parse(FTMPL)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Recipe) String() string {
	b := &bytes.Buffer{}
	if err := rtmpl.Execute(b, &r); err != nil {
		log.Fatal(err)
	}
	s := strings.TrimSpace(b.String())
	s = strings.Replace(s, "\n", "", -1)
	bs := &bytes.Buffer{}
	if err := xml.EscapeText(bs, []byte(s)); err != nil {
		log.Fatal(err)
	}
	return bs.String()
}

var di = 1

func (r *Recipe) Date() string {
	d := time.Date(2013, time.January, di, 0, 0, 0, 0, time.UTC)
	di++
	return d.Format("2006-01-02T15:04:05Z")
}

const RTMPL = `
{{if .By}}Recipe By: {{.By}}<br />{{end}}
{{if .Size}}Serving Size: {{.Size}}<br />{{end}}
{{if .Prep}}Prep Time: {{.Prep}}<br />{{end}}
{{if .Prep}}Categories: {{.Categories}}<br />{{end}}
{{if .Ingreds}}<ul>{{range .Ingreds}}<li>{{.}}</li>{{end}}</ul>{{end}}
{{.Directions}}
`

const FTMPL = `<ns0:feed xmlns:ns0="http://www.w3.org/2005/Atom">
<ns0:title type="html">My Blog</ns0:title>
<ns0:link href="http://mjibson.wordpress.com" rel="self" type="application/atom+xml" />
<ns0:link href="http://mjibson.wordpress.com" rel="self" type="application/atom+xml" />
<ns0:updated>2013-12-29T00:53:50Z</ns0:updated>
<ns0:generator>Blogger</ns0:generator>
<ns0:link href="http://mjibson.wordpress.com" rel="alternate" type="text/html" />
<ns0:link href="http://mjibson.wordpress.com" rel="alternate" type="text/html" />
{{range $i, $e := .}}<ns0:entry>
<ns0:category scheme="http://schemas.google.com/g/2005#kind" term="http://schemas.google.com/blogger/2008/kind#post" />
<ns0:id>post-{{$i}}</ns0:id>
<ns0:author>
<ns0:name>ljibson</ns0:name>
</ns0:author>
<ns0:content type="html">{{.String}}</ns0:content>
<ns0:published>{{.Date}}</ns0:published>
<ns0:title type="html">{{.Name}}</ns0:title>
<ns0:link href="http://mjibson.wordpress.com/post-{{$i}}/" rel="self" type="application/atom+xml" />
<ns0:link href="http://mjibson.wordpress.com/post-{{$i}}/" rel="alternate" type="text/html" />
</ns0:entry>
{{end}}</ns0:feed>`
