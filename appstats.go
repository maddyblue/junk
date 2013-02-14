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

// Package appstats is a library for profiling the RPC performance of Google
// App Enigne applications.
//
// Installation instructions: https://github.com/mjibson/appstats
//
// Reference: https://developers.google.com/appengine/docs/python/tools/appstats
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
)

func DefaultShouldRecord(r *http.Request) bool {
	if RECORD_FRACTION >= 1.0 {
		return true
	}

	return rand.Float64() < RECORD_FRACTION
}

type Context struct {
	appengine.Context

	req   *http.Request
	stats *RequestStats
}

func (c Context) Call(service, method string, in, out appengine_internal.ProtoMessage, opts *appengine_internal.CallOptions) error {
	c.stats.wg.Add(1)
	defer c.stats.wg.Done()

	stat := RPCStat{
		Service:   service,
		Method:    method,
		Start:     time.Now(),
		Offset:    time.Since(c.stats.Start),
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

	c.stats.lock.Lock()
	c.stats.RPCStats = append(c.stats.RPCStats, stat)
	c.stats.Cost += stat.Cost
	c.stats.lock.Unlock()
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
		stats: &RequestStats{
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
		stats:   c.stats,
	}
}

func (c Context) Save() {
	c.stats.wg.Wait()
	c.stats.Duration = time.Since(c.stats.Start)

	var buf_part, buf_full bytes.Buffer
	enc_full := gob.NewEncoder(&buf_full)
	full := stats_full{
		Header: c.req.Header,
		Stats:  c.stats,
	}
	err := enc_full.Encode(&full)
	if err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	}
	enc_part := gob.NewEncoder(&buf_part)
	part := stats_part(*c.stats)
	for i := range part.RPCStats {
		part.RPCStats[i].StackData = ""
		part.RPCStats[i].In = ""
		part.RPCStats[i].Out = ""
	}
	err = enc_part.Encode(&part)
	if err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	}

	item_part := &memcache.Item{
		Key:   c.stats.PartKey(),
		Value: buf_part.Bytes(),
	}

	item_full := &memcache.Item{
		Key:   c.stats.FullKey(),
		Value: buf_full.Bytes(),
	}

	u := url.URL{
		Scheme:   "http",
		Host:     c.req.Host,
		Path:     "/_ah/stats/details",
		RawQuery: fmt.Sprintf("time=%v", c.stats.Start.Nanosecond()),
	}

	c.Infof("Saved; %s: %s, %s: %s, link: %v",
		item_part.Key,
		ByteSize(len(item_part.Value)),
		item_full.Key,
		ByteSize(len(item_full.Value)),
		u.String(),
	)

	memcache.SetMulti(c.Context, []*memcache.Item{item_part, item_full})
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
	if r.c.stats.Status == 0 {
		r.WriteHeader(http.StatusOK)
	}

	return r.ResponseWriter.Write(b)
}

func (r responseWriter) WriteHeader(i int) {
	r.c.stats.Status = i
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
