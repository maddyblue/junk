package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Review struct {
	Name     string
	Rating   int
	Style    string
	Reviewer string
	Link     string
	Price    string
	Date     string
}

func parseReviews(r io.Reader) ([]*Review, error) {
	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.LazyQuotes = true
	// Ignore first row.
	cr.Read()

	// Map name -> style to make sure there aren't multiple mappings.
	styles := map[string]string{}

	var reviews []*Review
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rev := parseRecord(record)
		if rev == nil {
			continue
		}

		if style, ok := styles[rev.Name]; !ok {
			styles[rev.Name] = rev.Style
		} else if rev.Style != style {
			fmt.Printf("%s: saw %s, got %s\n", rev.Name, style, rev.Style)
		}

		reviews = append(reviews, rev)
	}
	return reviews, nil
}

func parseRecord(record []string) *Review {
	rating, err := parseRating(record[4])
	if rating == nil {
		return nil
	}
	price, err := parsePrice(record[6])
	if err != nil {
		return nil
	}
	date, err := parseDate(record[7])
	if err != nil {
		return nil
	}

	return &Review{
		Name:     record[1],
		Reviewer: record[2],
		Link:     record[3],
		Rating:   *rating,
		Style:    clean(record[5]),
		Price:    price,
		Date:     date,
	}
}

var fix = map[string]string{
	"borubon":    "bourbon",
	"highland":   "highlands",
	"lowland":    "lowlands",
	"islands":    "island",
	"cambeltown": "campbeltown",
	"cambletown": "campbeltown",
	"campeltown": "campbeltown",
}

func clean(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	if ns, ok := fix[s]; ok {
		s = ns
	}
	return strings.Title(s)
}

var rating100 = regexp.MustCompile(`^([0-9]+)/100`)

func parseRating(s string) (*int, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" || s == "na" {
		return nil, nil
	}
	// Ratings are from 0 (worst) - 100 (best).

	// Sometimes ratings are 0-100, sometimes 0-10. Let's assume anything
	// in the <= 10 range is not just a very low 0-100 rating.
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		i := int(f)
		// If it's <= 10 then I'm not sure which rating scale it's in, so bail.
		if i <= 10 {
			return nil, nil
		}
		return &i, nil
	}

	if matches := rating100.FindStringSubmatch(s); matches != nil {
		i, err := strconv.Atoi(matches[1])
		return &i, err
	}

	return nil, errors.Errorf("unknown rating: %s", s)
}

var priceMap = map[*regexp.Regexp]string{
	regexp.MustCompile(`^([0-9\.]+)$`):         "$$$1",
	regexp.MustCompile(`^\$([0-9\.]+)$`):       "$0",
	regexp.MustCompile(`^\$([0-9\.]+) usd?`):   "$$$1",
	regexp.MustCompile(`^\$([0-9\.]+) c[an]d`): "$$$1 CAD",
	regexp.MustCompile(`^\$([0-9\.]+) aud`):    "$$$1 AUD",
	regexp.MustCompile(`^([0-9\.]+)€`):         "€$1",
	regexp.MustCompile(`^([0-9\.]+) eu`):       "€$1",
	regexp.MustCompile(`^([0-9\,]+) eu`):       "€$1",
}

func parsePrice(s string) (string, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" || s == "$" {
		return "", nil
	}
	for re, replace := range priceMap {
		if re.MatchString(s) {
			return re.ReplaceAllString(s, replace), nil
		}
	}
	return s, nil
}

func parseDate(s string) (string, error) {
	s = strings.ReplaceAll(s, "//", "/")
	s = strings.ReplaceAll(s, "`", "")
	for _, layout := range []string{
		"1/2/06",
		"1/2/2006",
		"2/1/06",
		"2-1-2006",
		"1-2-2006",
		"1-2-06",
		"2-1-06",
	} {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t.Format("2006-01-02"), nil
		}
	}
	return "", errors.Errorf("unknown date: %s", s)
}
