package main

import "fmt"

type Tok int

func emit(tok Tok, val string) {
	fmt.Printf("%v: %d: %q\n", tok, len(val), val)
}

func emitToken(tok Tok) {
	fmt.Printf("TOKEN: %v (%[1]c)\n", tok)
}

func main() {
	fmt.Println(lexSQL([]byte(`iddent bleeh as as of not between not in not like with time as not with`)))
	fmt.Println(lexSQL([]byte(`B'' B'0' B'1' B'01' B'1010110100011'`)))
	fmt.Println(lexSQL([]byte(`; . .. + < << <<= <= << <<=`)))
	fmt.Println(lexSQL([]byte(`x'636174'`)))
	fmt.Println(lexSQL([]byte(`e e'\xc3\xb1' e'\u000a\U0000000A\x0a\X0A\n\012\\n' e`)))
	fmt.Println(lexSQL([]byte(`b b'\u000a\U0000000A\x0a\X0A\n\012\\n' b`)))
	fmt.Println(lexSQL([]byte("\"\xc3\x28\"")))
	fmt.Println(lexSQL([]byte(`"☃" "" "blah" "yo""hello" """start" "end""" """" "double """" done"`)))
	fmt.Println(lexSQL([]byte("a b a$1 aBbC C 01 1e0 .01 01. 0x1 0X2 0x0a $12 34\n--4555 \n 67$78 $1 2")))
	fmt.Println(lexSQL([]byte(`'''' '''''' '''start' 'end''' 'mid''dle'`)))
}
