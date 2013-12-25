package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	feed := Feed{}
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
		feed.Entry = append(feed.Entry, &Entry{
			Title:    r.Name,
			Content:  r.Atom(),
			Category: r.Category(),
		})
	}
	b, _ := xml.MarshalIndent(&feed, "", "  ")
	ioutil.WriteFile("feed.xml", b, 0666)
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

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("r").Parse(RTMPL)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Recipe) Atom() *Text {
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, &r); err != nil {
		log.Fatal(err)
	}
	return &Text{
		Body: strings.TrimSpace(b.String()),
	}
}

func (r *Recipe) Category() []*Category {
	if r.Categories == "" {
		return nil
	}
	return []*Category{
		&Category{
			Term: r.Categories,
		},
	}
}

const RTMPL = `
{{if .By}}
	<p>Recipe By: {{.By}}</p>
{{end}}
{{if .Size}}
	<p>Serving Size: {{.Size}}</p>
{{end}}
{{if .Prep}}
	<p>Prep Time: {{.Prep}}</p>
{{end}}
{{if .Ingreds}}
<ul>
{{range .Ingreds}}
	<li>{{.}}</li>
{{end}}
</ul>
{{end}}
{{.Directions}}
`

type Feed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Entry   []*Entry `xml:"entry"`
}

type Entry struct {
	Title    string      `xml:"title"`
	Content  *Text       `xml:"content"`
	Category []*Category `xml:"category"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

type Text struct {
	Type string `xml:"type,attr"`
	Body string `xml:",chardata"`
}
