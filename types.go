package appstats

import (
	"fmt"
	"net/http"
	"time"
)

const (
	KEY_PREFIX = "__appstats__:"
	KEY_PART   = KEY_PREFIX + "%06d:part"
	KEY_FULL   = KEY_PREFIX + "%v:full"
	DISTANCE   = 100
	MODULUS    = 1000
)

type RequestStats struct {
	Method      string
	Path, Query string
	Status      int
	Start       time.Time
	Duration    time.Duration
	RPCStats    []RPCStat
}

type stats_part RequestStats

type stats_full struct {
	Header http.Header
	Stats  *RequestStats
}

func (r RequestStats) PartKey() string {
	t := (r.Start.Nanosecond() / 1000 / DISTANCE) % MODULUS * DISTANCE
	return fmt.Sprintf(KEY_PART, t)
}

func (r RequestStats) FullKey() string {
	return fmt.Sprintf(KEY_FULL, r.Start.Nanosecond())
}

type RPCStat struct {
	Service, Method string
	Start           time.Time
	Duration        time.Duration
}

type AllRequestStats []*RequestStats

func (s AllRequestStats) Len() int {
	return len(s)
}

func (s AllRequestStats) Less(i, j int) bool {
	return s[i].Start.Sub(s[j].Start) > 0
}

func (s AllRequestStats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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
