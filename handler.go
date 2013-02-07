/*
 * Copyright (c) 2013 Matt Jibson <matt.jibson@gmail.com>
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

package appstats

import (
	"appengine"
	"appengine/memcache"
	"bytes"
	"encoding/gob"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

var templates *template.Template
var staticFiles map[string][]byte

func init() {
	templates = template.New("appstats").Funcs(funcs)
	templates.Parse(HTML_BASE)
	templates.Parse(HTML_MAIN)
	templates.Parse(HTML_DETAILS)
	templates.Parse(HTML_FILE)

	staticFiles = map[string][]byte{
		"app_engine_logo_sm.gif": static_app_engine_logo_sm_gif,
		"appstats_css.css":       static_appstats_css_css,
		"appstats_js.js":         static_appstats_js_js,
		"gantt.js":               static_gantt_js,
		"minus.gif":              static_minus_gif,
		"pix.gif":                static_pix_gif,
		"plus.gif":               static_plus_gif,
	}
}

func AppstatsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/details") {
		Details(w, r)
	} else if strings.HasSuffix(r.URL.Path, "/file") {
		File(w, r)
	} else if strings.Contains(r.URL.Path, "/static/") {
		Static(w, r)
	} else {
		Index(w, r)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, MODULUS)
	for i := range keys {
		keys[i] = fmt.Sprintf(KEY_PART, i*DISTANCE)
	}

	c := appengine.NewContext(r)
	items, err := memcache.GetMulti(c, keys)
	if err != nil {
		return
	}

	ars := AllRequestStats{}
	for _, v := range items {
		var buf bytes.Buffer
		_, _ = buf.Write(v.Value)
		dec := gob.NewDecoder(&buf)
		t := stats_part{}
		err := dec.Decode(&t)
		if err != nil {
			continue
		}
		r := RequestStats(t)
		ars = append(ars, &r)
	}
	sort.Sort(ars)

	requestById := make(map[int]*RequestStats, len(ars))
	idByRequest := make(map[*RequestStats]int, len(ars))
	requests := make(map[int]*StatByName)
	byRequest := make(map[int]map[string]int)
	for i, v := range ars {
		idx := i + 1
		requestById[idx] = v
		idByRequest[v] = idx
		requests[idx] = &StatByName{
			RequestStats: v,
		}
		byRequest[idx] = make(map[string]int)
	}

	requestByPath := make(map[string][]int)
	byCount := make(map[string]int)
	byRPC := make(map[string]map[string]int)
	byPath := make(map[string]map[string]int)
	for _, t := range ars {
		id := idByRequest[t]

		requestByPath[t.Path] = append(requestByPath[t.Path], id)

		for _, r := range t.RPCStats {
			rpc := r.Name()

			byRequest[id][rpc]++
			byCount[rpc]++

			if _, present := byRPC[rpc]; !present {
				byRPC[rpc] = make(map[string]int)
			}
			byRPC[rpc][t.Path]++

			if _, present := byPath[t.Path]; !present {
				byPath[t.Path] = make(map[string]int)
			}
			byPath[t.Path][rpc]++
		}
	}

	for k, v := range byRequest {
		stats := StatsByName{}
		for rpc, count := range v {
			stats = append(stats, &StatByName{
				Name:  rpc,
				Count: count,
			})
		}
		sort.Sort(Reverse{stats})
		requests[k].SubStats = stats
	}

	statsByRPC := make(map[string]StatsByName)
	for k, v := range byRPC {
		stats := StatsByName{}
		for path, count := range v {
			stats = append(stats, &StatByName{
				Name:  path,
				Count: count,
			})
		}
		sort.Sort(Reverse{stats})
		statsByRPC[k] = stats
	}

	allStatsByCount := StatsByName{}
	for k, v := range byCount {
		allStatsByCount = append(allStatsByCount, &StatByName{
			Name:     k,
			Count:    v,
			SubStats: statsByRPC[k],
		})
	}
	sort.Sort(Reverse{allStatsByCount})

	pathStatsByCount := StatsByName{}
	for k, v := range byPath {
		stats := StatsByName{}
		total := 0
		for name, count := range v {
			stats = append(stats, &StatByName{
				Name:  name,
				Count: count,
			})
			total += count
		}
		sort.Sort(Reverse{stats})

		pathStatsByCount = append(pathStatsByCount, &StatByName{
			Name:       k,
			Count:      total,
			SubStats:   stats,
			Requests:   len(requestByPath[k]),
			RecentReqs: requestByPath[k],
		})
	}
	sort.Sort(Reverse{pathStatsByCount})

	v := struct {
		Env                 map[string]string
		Requests            map[int]*StatByName
		RequestStatsByCount map[int]*StatByName
		AllStatsByCount     StatsByName
		PathStatsByCount    StatsByName
	}{
		Env: map[string]string{
			"APPLICATION_ID": appengine.AppID(c),
		},
		Requests:         requests,
		AllStatsByCount:  allStatsByCount,
		PathStatsByCount: pathStatsByCount,
	}

	_ = templates.ExecuteTemplate(w, "main", v)
}

func Details(w http.ResponseWriter, r *http.Request) {
	qtime := r.URL.Query().Get("time")
	key := fmt.Sprintf(KEY_FULL, qtime)

	c := appengine.NewContext(r)
	item, err := memcache.Get(c, key)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	_, _ = buf.Write(item.Value)
	dec := gob.NewDecoder(&buf)
	full := stats_full{}
	err = dec.Decode(&full)
	if err != nil {
		// todo: send down an empty request
		return
	}

	byCount := make(map[string]int)
	durationCount := make(map[string]time.Duration)
	var _real, _api time.Duration
	for _, r := range full.Stats.RPCStats {
		rpc := r.Name()

		// byCount
		if _, present := byCount[rpc]; !present {
			byCount[rpc] = 0
			durationCount[rpc] = 0
		}
		byCount[rpc] += 1
		durationCount[rpc] += r.Duration
		_real += r.Duration
	}

	allStatsByCount := StatsByName{}
	for k, v := range byCount {
		allStatsByCount = append(allStatsByCount, &StatByName{
			Name:     k,
			Count:    v,
			Duration: durationCount[k],
		})
	}
	sort.Sort(allStatsByCount)

	v := struct {
		Env             map[string]string
		Record          *RequestStats
		Header          http.Header
		AllStatsByCount StatsByName
		Real, Api       time.Duration
	}{
		Env: map[string]string{
			"APPLICATION_ID": appengine.AppID(c),
		},
		Record:          full.Stats,
		Header:          full.Header,
		AllStatsByCount: allStatsByCount,
		Real:            _real,
		Api:             _api,
	}

	_ = templates.ExecuteTemplate(w, "details", v)
}

func File(w http.ResponseWriter, r *http.Request) {
	fname := r.URL.Query().Get("f")
	n := r.URL.Query().Get("n")
	lineno, _ := strconv.Atoi(n)
	c := appengine.NewContext(r)

	f, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	fp := make(map[int]string)
	for k, v := range strings.Split(string(f), "\n") {
		fp[k+1] = v
	}

	v := struct {
		Env      map[string]string
		Filename string
		Lineno   int
		Fp       map[int]string
	}{
		Env: map[string]string{
			"APPLICATION_ID": appengine.AppID(c),
		},
		Filename: fname,
		Lineno:   lineno,
		Fp:       fp,
	}

	_ = templates.ExecuteTemplate(w, "file", v)
}

func Static(w http.ResponseWriter, r *http.Request) {
	fname := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if v, present := staticFiles[fname]; present {
		h := w.Header()

		if strings.HasSuffix(r.URL.Path, ".css") {
			h.Set("Content-type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			h.Set("Content-type", "text/javascript")
		}

		h.Set("Cache-Control", "public, max-age=expiry")
		expires := time.Now().Add(time.Hour)
		h.Set("Expires", expires.Format(time.RFC1123))

		w.Write(v)
	}
}
