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
		{
			input:  `12/*34*/56`,
			output: []string{"12", "56"},
		},
		{
			input:  `12/*34*/56/*78*/90`,
			output: []string{"12", "56", "90"},
		},
		{
			input:  `1/*2/*3*/4*/5`,
			output: []string{"1", "5"},
		},
		{
			input:  `1/*2/*3*/4/*5/*6*/7*/8*/9`,
			output: []string{"1", "9"},
		},
		{
			input: `  `,
		},
		{
			input: `/*`,
			err:   "unterminated comment",
		},
		{
			input: `/*/`,
			err:   "unterminated comment",
		},
		{
			input: `/* /* */   * /`,
			err:   "unterminated comment",
		},
		{
			input: `/* */* */`,
			err:   "ERROR at 5 of 9",
		},
		{
			input:  `e'\x0a'`,
			output: []string{"\n"},
		},
		{
			input:  `e'\n'`,
			output: []string{"\n"},
		},
		{
			input:  `e'\u000a'`,
			output: []string{"\n"},
		},
		{
			input:  `e'\U0000000a'`,
			output: []string{"\n"},
		},
		{
			input:  `e'\012'`,
			output: []string{"\n"},
		},
		{
			input:  `e'\x0a\012\n\u000a\U0000000a'`,
			output: []string{"\n\n\n\n\n"},
		},
		{
			input:  `e'\\'`,
			output: []string{"\\"},
		},
		{
			input:  `e'\\n\e'`,
			output: []string{"\\ne"},
		},
		{
			input: `e'\x'`,
			err:   "unterminated string",
		},
		{
			input: `e'\x1'`,
			err:   "unterminated string",
		},
		{
			input: `e'\xx'`,
			err:   "unterminated string",
		},
		{
			input: `e'\01'`,
			err:   "unterminated string",
		},
		{
			input: `e'\u001'`,
			err:   "unterminated string",
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Logf("%s", tc.input)
			out, err := lexSQL([]rune(tc.input))
			if !IsError(err, tc.err) {
				t.Fatalf("unexpected: %v", err)
			}
			if !reflect.DeepEqual(out, tc.output) {
				for i, v := range out {
					t.Logf("%d: %q (%d)\n", i, v, len(v))
				}
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
