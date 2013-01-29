package appstats

import (
	"fmt"
	"time"
)

type RequestStats struct {
	Method      string
	Path, Query string
	Status      int
	Start       time.Time
	Duration    time.Duration
	RPCStats    []RPCStat
}

const (
	KEY_PREFIX = "__appstats__"
	KEY_PART   = ":part"
	TMPL       = KEY_PREFIX + ":%06d"
	DISTANCE   = 100
	MODULUS    = 1000
)

func (r RequestStats) Key() string {
	t := (r.Start.Nanosecond() / 1000 / DISTANCE) % MODULUS * DISTANCE
	return fmt.Sprintf(TMPL, t)
}

type RPCStat struct {
	Service, Method string
	Start           time.Time
	Duration        time.Duration
}
