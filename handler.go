package appstats

import (
	"appengine"
	"appengine/memcache"
	"bytes"
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"time"
)

var templates *template.Template

func init() {
	templates = template.New("appstats").Funcs(funcs)
	templates.Parse(HTML_BASE)
	templates.Parse(HTML_MAIN)
	templates.Parse(HTML_DETAILS)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/details") {
		Details(w, r)
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
	byRequest := make(map[int]map[string]*StatByName)
	for i, v := range ars {
		idx := i + 1
		requestById[idx] = v
		idByRequest[v] = idx
		requests[idx] = &StatByName{
			RequestStats: v,
		}
		byRequest[idx] = make(map[string]*StatByName)
	}

	requestByPath := make(map[string][]int)
	byCount := make(map[string]int)
	byRPC := make(map[string]map[string]*StatByName)
	byPath := make(map[string]map[string]*StatByName)
	for _, t := range ars {
		id := idByRequest[t]

		if _, present := requestByPath[t.Path]; !present {
			requestByPath[t.Path] = make([]int, 0)
		}
		requestByPath[t.Path] = append(requestByPath[t.Path], id)

		for _, r := range t.RPCStats {
			rpc := r.Name()

			// byRequest
			if _, present := byRequest[id][rpc]; !present {
				byRequest[id][rpc] = &StatByName{
					Name: rpc,
				}
			}
			byRequest[id][rpc].Count++

			// byCount
			if _, present := byCount[rpc]; !present {
				byCount[rpc] = 0
			}
			byCount[rpc] += 1

			// byRPC
			if _, present := byRPC[rpc]; !present {
				byRPC[rpc] = make(map[string]*StatByName)
			}
			if _, present := byRPC[rpc][t.Path]; !present {
				byRPC[rpc][t.Path] = &StatByName{
					Name: t.Path,
				}
			}
			byRPC[rpc][t.Path].Count++

			// byPath
			if _, present := byPath[t.Path]; !present {
				byPath[t.Path] = make(map[string]*StatByName)
			}
			if _, present := byPath[t.Path][rpc]; !present {
				byPath[t.Path][rpc] = &StatByName{
					Name: rpc,
				}
			}
			byPath[t.Path][rpc].Count++
		}
	}

	for k, v := range byRequest {
		stats := StatsByName{}
		for _, s := range v {
			stats = append(stats, s)
		}
		sort.Sort(stats)
		requests[k].SubStats = stats
	}

	statsByRPC := make(map[string]StatsByName)
	for k, v := range byRPC {
		stats := StatsByName{}
		for _, s := range v {
			stats = append(stats, s)
		}
		sort.Sort(stats)
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
	sort.Sort(allStatsByCount)

	pathStatsByCount := StatsByName{}
	for k, v := range byPath {
		stats := StatsByName{}
		count := 0
		for _, s := range v {
			stats = append(stats, s)
			count += s.Count
		}
		sort.Sort(stats)

		pathStatsByCount = append(pathStatsByCount, &StatByName{
			Name:       k,
			Count:      count,
			SubStats:   stats,
			Requests:   len(requestByPath[k]),
			RecentReqs: requestByPath[k],
		})
	}
	sort.Sort(pathStatsByCount)

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
