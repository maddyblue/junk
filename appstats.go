package appstats

import (
	"appengine"
	"appengine/memcache"
	"appengine/user"
	"appengine_internal"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/gob"
	"net/http"
	"runtime/debug"
	"time"
)

const (
	PROTO_BUF_MAX = 150
)

type Context struct {
	appengine.Context

	req   *http.Request
	stats *RequestStats
}

func (c Context) Call(service, method string, in, out proto.Message, opts *appengine_internal.CallOptions) error {
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

	if len(stat.In) > PROTO_BUF_MAX {
		stat.In = stat.In[:PROTO_BUF_MAX] + "..."
	}
	if len(stat.Out) > PROTO_BUF_MAX {
		stat.Out = stat.Out[:PROTO_BUF_MAX] + "..."
	}

	c.stats.RPCStats = append(c.stats.RPCStats, stat)
	return err
}

func NewContext(req *http.Request) Context {
	c := appengine.NewContext(req)
	u := user.Current(c)
	return Context{
		Context: c,
		req:     req,
		stats: &RequestStats{
			User:   u.String(),
			Admin:  u.Admin,
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
	c.stats.Duration = time.Since(c.stats.Start)

	var buf_part, buf_full bytes.Buffer
	enc_part := gob.NewEncoder(&buf_part)
	part := c.stats
	err := enc_part.Encode(&part)
	if err != nil {
		c.Errorf("appstats Save error: %v", err)
		return
	}
	enc_full := gob.NewEncoder(&buf_full)
	full := stats_full{
		Header: c.req.Header,
		Stats:  c.stats,
	}
	err = enc_full.Encode(&full)
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

	memcache.SetMulti(c.Context, []*memcache.Item{item_part, item_full})
}
