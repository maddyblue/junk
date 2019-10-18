package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/lib/pq"
	"github.com/mjibson/sr"
)

var (
	flagParseTSV = flag.String("parse-tsv", "", "TSV file to parse")
)

type Specification struct {
	Addr string `default:"localhost:4001"`
	DB   string `default:"postgres://root@localhost:26257/sr?sslmode=disable"`
}

func main() {
	flag.Parse()

	var spec Specification
	err := envconfig.Process("sr", &spec)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbURL, err := url.Parse(spec.DB)
	if err != nil {
		log.Fatal(err)
	}

	db := mustInitDB(dbURL.String())
	defer db.Close()
	fmt.Println("inited", dbURL)

	s := &cf.SRContext{
		DB: db,
		X:  sqlx.NewDb(db, "postgres"),
	}

	if *flagParseTSV != "" {
		f, err := os.Open(*flagParseTSV)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		revs, err := parseReviews(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("importing", len(revs), "reviews")

		if _, err := s.DB.Exec(`
			DROP TABLE IF EXISTS reviews;

			CREATE TABLE reviews (
				spirit   STRING NOT NULL,
				name     STRING NOT NULL,
				rating   INT8 NOT NULL,
				style    STRING NOT NULL,
				reviewer STRING,
				link     STRING,
				price    STRING,
				date     DATE
			);
		`); err != nil {
			log.Fatal(err)
		}

		txn, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		stmt, err := txn.Prepare(pq.CopyIn("reviews", "spirit", "name", "reviewer", "link", "rating", "style", "price", "date"))
		if err != nil {
			log.Fatal(err)
		}

		nullEmpty := func(s string) *string {
			r := strings.TrimSpace(s)
			if r == "" {
				return nil
			}
			return &r
		}

		for _, rev := range revs {
			_, err = stmt.Exec(
				"whiskey",
				rev.Name,
				nullEmpty(rev.Reviewer),
				nullEmpty(rev.Link),
				rev.Rating,
				nullEmpty(rev.Style),
				nullEmpty(rev.Price),
				nullEmpty(rev.Date),
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		err = stmt.Close()
		if err != nil {
			log.Fatal(err)
		}

		err = txn.Commit()
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	mux := http.NewServeMux()
	mux.Handle("/api/Avg", s.Wrap(s.Avg))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("HTTP listen on addr:", spec.Addr)
	log.Fatal(http.ListenAndServe(spec.Addr, mux))
}
