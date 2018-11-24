
//line lex.rl:1
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	
	//"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

var _ = fmt.Println
var _ = strconv.Atoi

func lexSQL(data []byte) error {

	
//line lex.rl:17
	
//line lex.go:22
const scanner_start int = 12
const scanner_first_final int = 12
const scanner_error int = 0

const scanner_en_main int = 12


//line lex.rl:18

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
		numQuote int
		b []byte
	)
	str := func() { s = string(data[mark:p]) }
        _, _, _, _, _, _ = uval, err, isFconst, isUpper, isNotASCII, str

	
//line lex.go:50
	{
	cs = scanner_start
	}

//line lex.go:55
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 12:
		goto st_case_12
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 13:
		goto st_case_13
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 14:
		goto st_case_14
	case 4:
		goto st_case_4
	case 15:
		goto st_case_15
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 16:
		goto st_case_16
	case 8:
		goto st_case_8
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 11:
		goto st_case_11
	case 22:
		goto st_case_22
	}
	goto st_out
tr33:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
	goto st12
tr45:
//line lex.rl:40

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
		
	goto st12
tr55:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
	goto st12
tr67:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st12
tr79:
//line lex.rl:73

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
		
	goto st12
tr89:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:227
		switch data[p] {
		case 32:
			goto st12
		case 34:
			goto tr22
		case 36:
			goto tr23
		case 39:
			goto tr24
		case 43:
			goto tr25
		case 45:
			goto tr26
		case 46:
			goto tr27
		case 48:
			goto tr28
		case 95:
			goto tr31
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 49 <= data[p] && data[p] <= 57 {
						goto tr29
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st12
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr30
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr31
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr32
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line lex.rl:102
 isNotASCII = true 
	goto st1
tr3:
//line lex.rl:116
 numQuote++ 
	goto st1
tr5:
//line lex.rl:116
 numQuote++ 
//line lex.rl:102
 isNotASCII = true 
	goto st1
tr22:
//line lex.rl:37
 mark = p 
	goto st1
tr46:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st1
tr56:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
	goto st1
tr68:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st1
tr80:
//line lex.rl:73

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
		
//line lex.rl:37
 mark = p 
	goto st1
tr90:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:411
		if data[p] == 34 {
			goto st13
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:116
 numQuote++ 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:428
		switch data[p] {
		case 32:
			goto tr33
		case 34:
			goto st2
		case 36:
			goto tr35
		case 39:
			goto tr36
		case 43:
			goto tr37
		case 45:
			goto tr38
		case 46:
			goto tr39
		case 48:
			goto tr40
		case 95:
			goto tr43
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 49 <= data[p] && data[p] <= 57 {
						goto tr41
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr33
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr42
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr43
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr44
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 34 {
			goto tr4
		}
		if data[p] <= 127 {
			goto tr3
		}
		goto tr5
tr23:
//line lex.rl:37
 mark = p 
	goto st3
tr35:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
	goto st3
tr47:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st3
tr57:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
	goto st3
tr69:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st3
tr91:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:611
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 32:
			goto tr45
		case 34:
			goto tr46
		case 36:
			goto tr47
		case 39:
			goto tr48
		case 43:
			goto tr49
		case 45:
			goto tr50
		case 46:
			goto tr51
		case 95:
			goto tr53
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto st14
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr45
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr52
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr53
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr54
tr10:
//line lex.rl:102
 isNotASCII = true 
	goto st4
tr11:
//line lex.rl:150
 numQuote++ 
	goto st4
tr13:
//line lex.rl:150
 numQuote++ 
//line lex.rl:102
 isNotASCII = true 
	goto st4
tr24:
//line lex.rl:37
 mark = p 
	goto st4
tr36:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
	goto st4
tr48:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st4
tr70:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st4
tr82:
//line lex.rl:73

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
		
//line lex.rl:37
 mark = p 
	goto st4
tr92:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line lex.go:798
		if data[p] == 39 {
			goto st15
		}
		if data[p] <= 127 {
			goto st4
		}
		goto tr10
tr12:
//line lex.rl:150
 numQuote++ 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:815
		switch data[p] {
		case 32:
			goto tr55
		case 34:
			goto tr56
		case 36:
			goto tr57
		case 39:
			goto st5
		case 43:
			goto tr59
		case 45:
			goto tr60
		case 46:
			goto tr61
		case 48:
			goto tr62
		case 95:
			goto tr65
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 49 <= data[p] && data[p] <= 57 {
						goto tr63
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr55
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr64
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr65
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr66
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 39 {
			goto tr12
		}
		if data[p] <= 127 {
			goto tr11
		}
		goto tr13
tr25:
//line lex.rl:37
 mark = p 
	goto st6
tr37:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
	goto st6
tr49:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st6
tr59:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
	goto st6
tr71:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st6
tr83:
//line lex.rl:73

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
		
//line lex.rl:37
 mark = p 
	goto st6
tr93:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line lex.go:1023
		if data[p] == 46 {
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
tr14:
//line lex.rl:66

			isFconst = true
		
	goto st7
tr27:
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr39:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr51:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr61:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr73:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr85:
//line lex.rl:73

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
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
tr95:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line lex.go:1199
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
tr78:
//line lex.rl:66

			isFconst = true
		
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:1215
		switch data[p] {
		case 32:
			goto tr67
		case 34:
			goto tr68
		case 36:
			goto tr69
		case 39:
			goto tr70
		case 43:
			goto tr71
		case 45:
			goto tr72
		case 46:
			goto tr73
		case 69:
			goto tr75
		case 95:
			goto tr76
		case 101:
			goto tr75
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto st16
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr67
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr74
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr76
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr77
tr26:
//line lex.rl:37
 mark = p 
	goto st8
tr38:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
	goto st8
tr50:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st8
tr60:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
	goto st8
tr72:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st8
tr84:
//line lex.rl:73

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
		
//line lex.rl:37
 mark = p 
	goto st8
tr94:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:1413
		switch data[p] {
		case 45:
			goto st17
		case 46:
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 10 {
			goto st12
		}
		goto st17
tr15:
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st18
tr29:
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st18
tr41:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st18
tr63:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:1528
		switch data[p] {
		case 32:
			goto tr67
		case 34:
			goto tr68
		case 36:
			goto tr69
		case 39:
			goto tr70
		case 43:
			goto tr71
		case 45:
			goto tr72
		case 46:
			goto tr78
		case 69:
			goto tr75
		case 95:
			goto tr76
		case 101:
			goto tr75
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr15
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr67
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr74
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr76
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr77
tr87:
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr31:
//line lex.rl:37
 mark = p 
	goto st19
tr30:
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr32:
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr43:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
	goto st19
tr42:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr44:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr53:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
	goto st19
tr52:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr54:
//line lex.rl:40

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
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr65:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
	goto st19
tr64:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr66:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr76:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
	goto st19
tr74:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr77:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
tr86:
//line lex.rl:105
 isUpper = true 
	goto st19
tr97:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st19
tr96:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st19
tr98:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:1941
		switch data[p] {
		case 32:
			goto tr79
		case 34:
			goto tr80
		case 36:
			goto st19
		case 39:
			goto tr82
		case 43:
			goto tr83
		case 45:
			goto tr84
		case 46:
			goto tr85
		case 95:
			goto st19
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto st19
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr79
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr86
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st19
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr87
tr75:
//line lex.rl:66

			isFconst = true
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:2012
		switch data[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto tr67
		case 34:
			goto tr68
		case 36:
			goto tr69
		case 39:
			goto tr70
		case 43:
			goto tr71
		case 45:
			goto tr72
		case 46:
			goto tr73
		case 95:
			goto tr76
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto st20
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr67
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr74
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr76
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr77
tr28:
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st21
tr40:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st21
tr62:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:2183
		switch data[p] {
		case 32:
			goto tr67
		case 34:
			goto tr68
		case 36:
			goto tr69
		case 39:
			goto tr70
		case 43:
			goto tr71
		case 45:
			goto tr72
		case 46:
			goto tr78
		case 69:
			goto tr75
		case 95:
			goto tr76
		case 101:
			goto tr75
		case 120:
			goto st11
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr15
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr67
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr74
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr76
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr77
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st22
			}
		default:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 32:
			goto tr89
		case 34:
			goto tr90
		case 36:
			goto tr91
		case 39:
			goto tr92
		case 43:
			goto tr93
		case 45:
			goto tr94
		case 46:
			goto tr95
		case 95:
			goto tr97
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 14:
				switch {
				case data[p] > 8:
					if 9 <= data[p] && data[p] <= 13 {
						goto tr89
					}
				default:
					goto st0
				}
			case data[p] > 47:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 64 {
						goto st0
					}
				case data[p] >= 48:
					goto st22
				}
			default:
				goto st0
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st0
					}
				case data[p] >= 71:
					goto tr96
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 103:
					goto tr97
				}
			default:
				goto st22
			}
		default:
			goto st22
		}
		goto tr98
	st_out:
	_test_eof12: cs = 12; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 14:
//line lex.rl:40

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
		
		case 16, 18, 20, 21:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 22:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
		case 19:
//line lex.rl:73

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
		
		case 13:
//line lex.rl:122

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Ident, string(b))
		
		case 15:
//line lex.rl:156

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = data[mark+1:p-1]
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					return fmt.Errorf("invalid UTF-8 string")
				}
				isNotASCII = false
			}
			emit(Sconst, string(b))
		
//line lex.go:2475
		}
	}

	_out: {}
	}

//line lex.rl:198


	return nil
}
