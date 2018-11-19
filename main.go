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
)

func emit(typ Typ, val string) {
	fmt.Printf("%s: %q\n", typ, val)
}

func emitToken(typ Typ) {
	fmt.Printf("%s\n", typ)
}

func main() {
	fmt.Println(lexSQL([]byte("a b a$1 aBbC C 01 1e0 .01 01. 0x1 0x0a $12 34\n--4555 \n 67$78 $1 2")))
}
