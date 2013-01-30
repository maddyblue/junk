package appstats

import (
	"appengine"
	"appengine/memcache"
	"appengine_internal"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/gob"
	"net/http"
	"time"
)

type Context struct {
	appengine.Context

	req   *http.Request
	stats *RequestStats
}

func (c Context) Call(service, method string, in, out proto.Message, opts *appengine_internal.CallOptions) error {
	stat := RPCStat{
		Service: service,
		Method:  method,
		Start:   time.Now(),
	}
	err := c.Context.Call(service, method, in, out, opts)
	stat.Duration = time.Since(stat.Start)
	c.stats.RPCStats = append(c.stats.RPCStats, stat)
	return err
}

func NewContext(req *http.Request) Context {
	return Context{
		Context: appengine.NewContext(req),
		req:     req,
		stats: &RequestStats{
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
	}
}

func (c Context) Save() {
	c.stats.Duration = time.Since(c.stats.Start)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(c.stats)
	if err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	}

	item := &memcache.Item{
		Key:   c.stats.Key() + KEY_PART,
		Value: buf.Bytes(),
	}

	memcache.Set(c.Context, item)
}
