package main

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		input  string
		output []string
		err    string
	}{
		{
			input:  ``,
			output: nil,
		},
		{
			input:  `'testing'`,
			output: []string{`testing`},
		},
		{
			input:  `'testing' `,
			output: []string{`testing`},
		},
		{
			input:  "'testing'\n",
			output: []string{`testing`},
		},
		{
			input:  "'testing'\n ",
			output: []string{`testing`},
		},
		{
			input:  "'testing' \n ",
			output: []string{`testing`},
		},
		{
			input:  "'testing' \n",
			output: []string{`testing`},
		},
		{
			input: `'test'
			'ing'`,
			output: []string{`testing`},
		},
		{
			input: `'before' 'test'
			'ing' 'after'`,
			output: []string{"before", "testing", "after"},
		},
		{
			input: `'before' 'test'
			'ing'`,
			output: []string{"before", "testing"},
		},
		{
			input: `'test'
			'ing' 'after'`,
			output: []string{"testing", "after"},
		},
		{
			input: `'`,
			err:   "unterminated string",
		},
		{
			input:  `''`,
			output: []string{""},
		},
		{
			input: `'''`,
			err:   "unterminated string",
		},
		{
			input:  `''''`,
			output: []string{`'`},
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Logf("%q", tc.input)
			out, err := lexSQL([]byte(tc.input))
			if !IsError(err, tc.err) {
				t.Fatalf("unexpected: %v", err)
			}
			if !reflect.DeepEqual(out, tc.output) {
				t.Fatalf("got: %v, expected %v", out, tc.output)
			}
		})
	}
}

// IsError returns true if the error string matches the supplied regex.
// An empty regex is interpreted to mean that a nil error is expected.
func IsError(err error, re string) bool {
	if err == nil && re == "" {
		return true
	}
	if err == nil || re == "" {
		return false
	}
	matched, merr := regexp.MatchString(re, err.Error())
	if merr != nil {
		return false
	}
	return matched
}
