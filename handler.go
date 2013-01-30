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
)

var templates *template.Template

func init() {
	templates = template.New("appstats")
	templates.Parse(HTML_BASE)
	templates.Parse(HTML_MAIN)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	Index(w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, MODULUS)
	for i := range keys {
		keys[i] = fmt.Sprintf(TMPL+KEY_PART, i*DISTANCE)
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
		t := new(RequestStats)
		err := dec.Decode(&t)
		if err != nil {
			continue
		}
		ars = append(ars, t)
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
			rpc := r.Service + "." + r.Method

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

	statsByRPC := make(map[string][]*StatByName)
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
		AllStatsByCount     []*StatByName
		PathStatsByCount    []*StatByName
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

type StatsByName []*StatByName

func (s StatsByName) Len() int {
	return len(s)
}

func (s StatsByName) Less(i, j int) bool {
	return s[i].Count < s[j].Count
}

func (s StatsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type StatByName struct {
	Name          string
	Count         int
	Cost, CostPct int
	SubStats      []*StatByName
	Requests      int
	RecentReqs    []int
	RequestStats  *RequestStats
}
