package main

import (
	"fmt"
	"strconv"
	
	//"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

func lexSQL(data []byte) error {

	%% machine scanner;
	%% write data;

	cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
        
        var (
	        mark int
        	s string
        	uval uint64
        	err error
        	isFconst bool
        	isUpper bool
        	isNotASCII bool
        )
        str := func() { s = string(data[mark:p]) }
        
	%%{
		action mark { mark = p }
		action str { str() }

		action placeholder {
			mark++
			str()
			uval, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				return err
			}
			if err == nil && uval > 1<<63 {
				return fmt.Errorf("integer value out of range: %d", uval)
			}
			emit(Placeholder, s)
		}
		action number {
			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		}
		action markZero {
			if mark == p && data[p] == '0' {
				mark++
			}
		}
		action fconst {
			isFconst = true
		}
		action hex {
			str()
			emit(Iconst, s)
		}
		action ident {
			if isNotASCII {
				str()
				//s = lex.NormalizeName(s)
			} else if isUpper {
				b := make([]byte, p-mark)
				for i, c := range data[mark:p] {
					if c >= 'A' && c <= 'Z' {
						c += 'a' - 'A'
					}
					b[i] = c
				}
				s = string(b)
			} else {
				str()
			}
			isUpper = false
			isNotASCII = false
			emit(Ident, s)
		}
		int = digit+;
		pn = ('-' | '+')?;
		number =
			pn
			(int $markZero ('.' >fconst)? digit* | '.' >fconst int)
			([eE] >fconst pn int)?
			;
		hex = '0x' xdigit+;
		placeholder = '$' int;
		identStart =
			'a'..'z'
			| 'A'..'Z' >{ isUpper = true }
			| '_'
			| 128..255
			;
		ident =
			identStart
			(identStart | digit | '$')*
			;
		top =
			  space
			| /--[^\n]*/
			| hex >mark %hex
			| number >mark %number
			| placeholder >mark %placeholder
			| ident >mark %ident
			#| ';' %{ emitToken(Semicolon) }
			;
		main :=
			top**
			;

		write init;
		write exec;
	}%%

	return nil
}
