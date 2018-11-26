package main

import "fmt"

type tok struct {
	typ Typ
	val string
}

//go:generate stringer -type Typ

type Typ int

const (
	_ Typ = iota
	Placeholder
	Ident
	Semicolon
	Fconst
	Iconst
	Sconst
	Bconst
)

func emit(typ Typ, val string) {
	fmt.Printf("%s: %d: %q\n", typ, len(val), val)
}

func emitToken(typ Typ) {
	fmt.Printf("%s\n", typ)
}

func main() {
	fmt.Println(lexSQL([]byte(`x'636174'`)))
	fmt.Println(lexSQL([]byte(`e e'\xc3\xb1' e'\u000a\U0000000A\x0a\X0A\n\012\\n' e`)))
	fmt.Println(lexSQL([]byte(`b b'\u000a\U0000000A\x0a\X0A\n\012\\n' b`)))
	fmt.Println(lexSQL([]byte("\"\xc3\x28\"")))
	fmt.Println(lexSQL([]byte(`"☃" "" "blah" "yo""hello" """start" "end""" """" "double """" done"`)))
	fmt.Println(lexSQL([]byte("a b a$1 aBbC C 01 1e0 .01 01. 0x1 0x0a $12 34\n--4555 \n 67$78 $1 2")))
	fmt.Println(lexSQL([]byte(`'''' '''''' '''start' 'end''' 'mid''dle'`)))
}
