package cf

import (
	"bytes"
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	servertiming "github.com/mitchellh/go-server-timing"
	"github.com/pkg/errors"
)

func init() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	fmt.Println("Files:")
	for _, f := range files {
		fmt.Printf("\t%v\n", f.Name())
	}

	addr := os.Getenv("DB_ADDR")
	if addr == "" {
		return
	}
	db, err := sql.Open("postgres", addr)
	if err != nil {
		panic(err)
	}
	srctx = &SRContext{
		DB: db,
		X:  sqlx.NewDb(db, "postgres"),
	}
	Avg = srctx.Wrap(srctx.Avg)
}

var (
	srctx *SRContext
	Avg   func(http.ResponseWriter, *http.Request)
)

type SRContext struct {
	DB *sql.DB
	X  *sqlx.DB
}

func (s *SRContext) Wrap(
	f func(context.Context, *http.Request) (interface{}, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*60)
		defer cancel()
		var sh servertiming.Header
		ctx = servertiming.NewContext(ctx, &sh)
		if v, err := url.ParseQuery(r.URL.RawQuery); err == nil {
			r.URL.RawQuery = v.Encode()
		}
		url := r.URL.String()
		start := time.Now()
		defer func() { fmt.Printf("%s: %s\n", url, time.Since(start)) }()
		tm := servertiming.FromContext(ctx).NewMetric("req").Start()
		res, err := f(ctx, r)
		tm.Stop()
		if len(sh.Metrics) > 0 {
			if len(sh.Metrics) > 10 {
				sh.Metrics = sh.Metrics[:10]
			}
			w.Header().Add(servertiming.HeaderKey, sh.String())
		}
		if err != nil {
			log.Printf("%s: %+v", url, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, gzip, err := resultToBytes(res)
		if err != nil {
			log.Printf("%s: %v", url, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeDataGzip(w, r, data, gzip)
	}
}

func (s *SRContext) Avg(ctx context.Context, r *http.Request) (interface{}, error) {
	var avgs []struct {
		Name    string
		Average float64
		Count   int
		Styles  pq.StringArray
	}

	err := s.X.SelectContext(ctx, &avgs, `
		SELECT
			name,
			avg(rating) AS average,
			count(*),
			array_agg(DISTINCT style) AS styles
		FROM
			reviews
		GROUP BY
			name
		HAVING
			count(*) > 5;
	`)
	return avgs, err
}

func resultToBytes(res interface{}) (data, gzipped []byte, err error) {
	data, err = json.Marshal(res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "json marshal")
	}
	var gz bytes.Buffer
	gzw, _ := gzip.NewWriterLevel(&gz, gzip.BestCompression)
	if _, err := gzw.Write(data); err != nil {
		return nil, nil, errors.Wrap(err, "gzip")
	}
	if err := gzw.Close(); err != nil {
		return nil, nil, errors.Wrap(err, "gzip close")
	}
	return data, gz.Bytes(), nil
}

func writeDataGzip(w http.ResponseWriter, r *http.Request, data, gzip []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "max-age=3600")
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Add("Content-Encoding", "gzip")
		w.Write(gzip)
	} else {
		w.Write(data)
	}
}
