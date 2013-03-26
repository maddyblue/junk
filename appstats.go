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
	"appengine/user"
	"appengine_internal"
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"
)

var (
	// RECORD_FRACTION is the fraction of requests to record.
	// Set to a number between 0.0 (none) and 1.0 (all).
	RECORD_FRACTION float64 = 1.0

	// ShouldRecord is the function used to determine if recording will occur
	// for a given request. The default is to use RECORD_FRACTION.
	ShouldRecord func(*http.Request) bool = DefaultShouldRecord

	// PROTO_BUF_MAX is the amount of protobuf data to record.
	// Data after this is truncated.
	PROTO_BUF_MAX int = 150

	// NAMESPACE is the memcache namespace under which to store appstats data.
	NAMESPACE string = "__appstats__"
)

const (
	URL         = "/_ah/stats/"
	URL_DETAILS = URL + "details"
	URL_FILE    = URL + "file"
	URL_STATIC  = URL + "static/"
)

func init() {
	http.HandleFunc(URL, AppstatsHandler)
}

func DefaultShouldRecord(r *http.Request) bool {
	if RECORD_FRACTION >= 1.0 {
		return true
	}

	return rand.Float64() < RECORD_FRACTION
}

type Context struct {
	appengine.Context

	req   *http.Request
	Stats *RequestStats
}

func (c Context) Call(service, method string, in, out appengine_internal.ProtoMessage, opts *appengine_internal.CallOptions) error {
	c.Stats.wg.Add(1)
	defer c.Stats.wg.Done()

	if service == "__go__" {
		return c.Context.Call(service, method, in, out, opts)
	}

	stat := RPCStat{
		Service:   service,
		Method:    method,
		Start:     time.Now(),
		Offset:    time.Since(c.Stats.Start),
		StackData: string(debug.Stack()),
	}
	err := c.Context.Call(service, method, in, out, opts)
	stat.Duration = time.Since(stat.Start)
	stat.In = in.String()
	stat.Out = out.String()
	stat.Cost = GetCost(out)

	if len(stat.In) > PROTO_BUF_MAX {
		stat.In = stat.In[:PROTO_BUF_MAX] + "..."
	}
	if len(stat.Out) > PROTO_BUF_MAX {
		stat.Out = stat.Out[:PROTO_BUF_MAX] + "..."
	}

	c.Stats.lock.Lock()
	c.Stats.RPCStats = append(c.Stats.RPCStats, stat)
	c.Stats.Cost += stat.Cost
	c.Stats.lock.Unlock()
	return err
}

func NewContext(req *http.Request) Context {
	c := appengine.NewContext(req)
	var uname string
	var admin bool
	if u := user.Current(c); u != nil {
		uname = u.String()
		admin = u.Admin
	}
	return Context{
		Context: c,
		req:     req,
		Stats: &RequestStats{
			User:   uname,
			Admin:  admin,
			Method: req.Method,
			Path:   req.URL.Path,
			Query:  req.URL.RawQuery,
			Start:  time.Now(),
		},
	}
}

func (c Context) FromContext(ctx appengine.Context) Context {
	return Context{
		Context: ctx,
		req:     c.req,
		Stats:   c.Stats,
	}
}

const bufMaxLen = 1000000

func (c Context) Save() {
	c.Stats.wg.Wait()
	c.Stats.Duration = time.Since(c.Stats.Start)

	var buf_part, buf_full bytes.Buffer
	full := stats_full{
		Header: c.req.Header,
		Stats:  c.Stats,
	}
	if err := gob.NewEncoder(&buf_full).Encode(&full); err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	} else if buf_full.Len() > bufMaxLen {
		// first try clearing stack traces
		for i := range full.Stats.RPCStats {
			full.Stats.RPCStats[i].StackData = ""
		}
		buf_full.Truncate(0)
		gob.NewEncoder(&buf_full).Encode(&full)
	}
	part := stats_part(*c.Stats)
	for i := range part.RPCStats {
		part.RPCStats[i].StackData = ""
		part.RPCStats[i].In = ""
		part.RPCStats[i].Out = ""
	}
	if err := gob.NewEncoder(&buf_part).Encode(&part); err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	}

	item_part := &memcache.Item{
		Key:   c.Stats.PartKey(),
		Value: buf_part.Bytes(),
	}

	item_full := &memcache.Item{
		Key:   c.Stats.FullKey(),
		Value: buf_full.Bytes(),
	}

	c.Infof("Saved; %s: %s, %s: %s, link: %v",
		item_part.Key,
		ByteSize(len(item_part.Value)),
		item_full.Key,
		ByteSize(len(item_full.Value)),
		c.URL(),
	)

	nc := context(c.req)
	memcache.SetMulti(nc, []*memcache.Item{item_part, item_full})
}

func (c Context) URL() string {
	u := url.URL{
		Scheme:   "http",
		Host:     c.req.Host,
		Path:     URL_DETAILS,
		RawQuery: fmt.Sprintf("time=%v", c.Stats.Start.Nanosecond()),
	}

	return u.String()
}

func context(r *http.Request) appengine.Context {
	c := appengine.NewContext(r)
	nc, _ := appengine.Namespace(c, NAMESPACE)
	return nc
}

type Handler struct {
	f func(appengine.Context, http.ResponseWriter, *http.Request)
}

func NewHandler(f func(appengine.Context, http.ResponseWriter, *http.Request)) Handler {
	return Handler{
		f: f,
	}
}

type responseWriter struct {
	http.ResponseWriter

	c Context
}

func (r responseWriter) Write(b []byte) (int, error) {
	// Emulate the behavior of http.ResponseWriter.Write since it doesn't
	// call our WriteHeader implementation.
	if r.c.Stats.Status == 0 {
		r.WriteHeader(http.StatusOK)
	}

	return r.ResponseWriter.Write(b)
}

func (r responseWriter) WriteHeader(i int) {
	r.c.Stats.Status = i
	r.ResponseWriter.WriteHeader(i)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ShouldRecord(r) {
		c := NewContext(r)
		rw := responseWriter{
			ResponseWriter: w,
			c:              c,
		}
		h.f(c, rw, r)
		c.Save()
	} else {
		c := appengine.NewContext(r)
		h.f(c, w, r)
	}
}
