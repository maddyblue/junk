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

	requestStats := make(map[string]*RequestStats)
	byCount := make(map[string]int)
	byRPC := make(map[string]map[string]*StatByName)
	byPath := make(map[string]map[string]*StatByName)
	for k, v := range items {
		var buf bytes.Buffer
		_, _ = buf.Write(v.Value)
		dec := gob.NewDecoder(&buf)
		t := new(RequestStats)
		err := dec.Decode(&t)
		if err != nil {
			continue
		}
		requestStats[k] = t

		for _, r := range t.RPCStats {
			rpc := r.Service + "." + r.Method
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

	statsByRPC := make(map[string][]*StatByName)
	for k, v := range byRPC {
		stats := &StatsByName{}
		for _, s := range v {
			stats.Stats = append(stats.Stats, s)
		}
		sort.Sort(stats)
		statsByRPC[k] = stats.Stats
	}

	allStatsByCount := new(StatsByName)
	for k, v := range byCount {
		allStatsByCount.Stats = append(allStatsByCount.Stats, &StatByName{
			Name:     k,
			Count:    v,
			SubStats: statsByRPC[k],
		})
	}
	sort.Sort(allStatsByCount)

	pathStatsByCount := new(StatsByName)
	for k, v := range byPath {
		stats := &StatsByName{}
		count := 0
		for _, s := range v {
			stats.Stats = append(stats.Stats, s)
			count += s.Count
		}
		sort.Sort(stats)

		pathStatsByCount.Stats = append(pathStatsByCount.Stats, &StatByName{
			Name:     k,
			Count:    count,
			SubStats: stats.Stats,
		})
	}
	sort.Sort(pathStatsByCount)

	v := struct {
		Env              map[string]string
		Requests         map[string]*RequestStats
		AllStatsByCount  []*StatByName
		PathStatsByCount []*StatByName
	}{
		Env: map[string]string{
			"APPLICATION_ID": appengine.AppID(c),
		},
		Requests:         requestStats,
		AllStatsByCount:  allStatsByCount.Stats,
		PathStatsByCount: pathStatsByCount.Stats,
	}

	_ = templates.ExecuteTemplate(w, "main", v)
}

type StatsByName struct {
	Stats []*StatByName
}

func (s *StatsByName) Len() int {
	return len(s.Stats)
}

func (s *StatsByName) Less(i, j int) bool {
	return s.Stats[i].Count < s.Stats[j].Count
}

func (s *StatsByName) Swap(i, j int) {
	s.Stats[i], s.Stats[j] = s.Stats[j], s.Stats[i]
}

type StatByName struct {
	Name          string
	Count         int
	Cost, CostPct int
	SubStats      []*StatByName
	Requests      int
	RecentReqs    []int
}
