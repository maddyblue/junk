
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
const scanner_start int = 10
const scanner_first_final int = 10
const scanner_error int = 0

const scanner_en_main int = 10


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
	case 10:
		goto st_case_10
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 11:
		goto st_case_11
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 12:
		goto st_case_12
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 13:
		goto st_case_13
	case 6:
		goto st_case_6
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 9:
		goto st_case_9
	case 19:
		goto st_case_19
	}
	goto st_out
tr26:
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
		
	goto st10
tr37:
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
		
	goto st10
tr46:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st10
tr57:
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
		
	goto st10
tr66:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line lex.go:193
		switch data[p] {
		case 32:
			goto st10
		case 34:
			goto tr16
		case 36:
			goto tr17
		case 43:
			goto tr18
		case 45:
			goto tr19
		case 46:
			goto tr20
		case 48:
			goto tr21
		case 95:
			goto tr24
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
						goto tr22
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st10
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr23
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr24
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr25
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
tr16:
//line lex.rl:37
 mark = p 
	goto st1
tr38:
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
tr47:
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
tr58:
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
tr67:
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
//line lex.go:345
		if data[p] == 34 {
			goto st11
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:116
 numQuote++ 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line lex.go:362
		switch data[p] {
		case 32:
			goto tr26
		case 34:
			goto st2
		case 36:
			goto tr28
		case 43:
			goto tr29
		case 45:
			goto tr30
		case 46:
			goto tr31
		case 48:
			goto tr32
		case 95:
			goto tr35
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
						goto tr33
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr26
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr34
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr35
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr36
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
tr17:
//line lex.rl:37
 mark = p 
	goto st3
tr28:
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
tr39:
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
tr48:
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
tr68:
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
//line lex.go:513
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 32:
			goto tr37
		case 34:
			goto tr38
		case 36:
			goto tr39
		case 43:
			goto tr40
		case 45:
			goto tr41
		case 46:
			goto tr42
		case 95:
			goto tr44
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
						goto st12
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr37
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr43
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr44
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr45
tr18:
//line lex.rl:37
 mark = p 
	goto st4
tr29:
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
tr40:
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
tr49:
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
tr60:
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
tr69:
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
//line lex.go:684
		if data[p] == 46 {
			goto tr8
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
tr8:
//line lex.rl:66

			isFconst = true
		
	goto st5
tr20:
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st5
tr31:
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
		
	goto st5
tr42:
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
		
	goto st5
tr51:
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
		
	goto st5
tr62:
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
		
	goto st5
tr71:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:66

			isFconst = true
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line lex.go:826
		if 48 <= data[p] && data[p] <= 57 {
			goto st13
		}
		goto st0
tr56:
//line lex.rl:66

			isFconst = true
		
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:842
		switch data[p] {
		case 32:
			goto tr46
		case 34:
			goto tr47
		case 36:
			goto tr48
		case 43:
			goto tr49
		case 45:
			goto tr50
		case 46:
			goto tr51
		case 69:
			goto tr53
		case 95:
			goto tr54
		case 101:
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
						goto st13
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr46
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
					goto tr54
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr55
tr19:
//line lex.rl:37
 mark = p 
	goto st6
tr30:
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
tr41:
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
tr50:
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
tr61:
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
tr70:
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
//line lex.go:1008
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto tr8
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 10 {
			goto st10
		}
		goto st14
tr9:
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st15
tr22:
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st15
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
		
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:1087
		switch data[p] {
		case 32:
			goto tr46
		case 34:
			goto tr47
		case 36:
			goto tr48
		case 43:
			goto tr49
		case 45:
			goto tr50
		case 46:
			goto tr56
		case 69:
			goto tr53
		case 95:
			goto tr54
		case 101:
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
						goto tr9
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr46
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
					goto tr54
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr55
tr64:
//line lex.rl:102
 isNotASCII = true 
	goto st16
tr24:
//line lex.rl:37
 mark = p 
	goto st16
tr23:
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st16
tr25:
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st16
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
	goto st16
tr34:
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
	goto st16
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
//line lex.rl:102
 isNotASCII = true 
	goto st16
tr44:
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
	goto st16
tr43:
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
	goto st16
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
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st16
tr54:
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
	goto st16
tr52:
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
	goto st16
tr55:
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
	goto st16
tr63:
//line lex.rl:105
 isUpper = true 
	goto st16
tr73:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
	goto st16
tr72:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:105
 isUpper = true 
	goto st16
tr74:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
//line lex.rl:37
 mark = p 
//line lex.rl:102
 isNotASCII = true 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:1404
		switch data[p] {
		case 32:
			goto tr57
		case 34:
			goto tr58
		case 36:
			goto st16
		case 43:
			goto tr60
		case 45:
			goto tr61
		case 46:
			goto tr62
		case 95:
			goto st16
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
				goto tr57
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr63
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st16
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr64
tr53:
//line lex.rl:66

			isFconst = true
		
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line lex.go:1473
		switch data[p] {
		case 43:
			goto st8
		case 45:
			goto st8
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 32:
			goto tr46
		case 34:
			goto tr47
		case 36:
			goto tr48
		case 43:
			goto tr49
		case 45:
			goto tr50
		case 46:
			goto tr51
		case 95:
			goto tr54
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
						goto st17
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr46
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
					goto tr54
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr55
tr21:
//line lex.rl:37
 mark = p 
//line lex.rl:61

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st18
tr32:
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
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:1606
		switch data[p] {
		case 32:
			goto tr46
		case 34:
			goto tr47
		case 36:
			goto tr48
		case 43:
			goto tr49
		case 45:
			goto tr50
		case 46:
			goto tr56
		case 69:
			goto tr53
		case 95:
			goto tr54
		case 101:
			goto tr53
		case 120:
			goto st9
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
						goto tr9
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr46
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
					goto tr54
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr55
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st19
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st19
			}
		default:
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 32:
			goto tr66
		case 34:
			goto tr67
		case 36:
			goto tr68
		case 43:
			goto tr69
		case 45:
			goto tr70
		case 46:
			goto tr71
		case 95:
			goto tr73
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 14:
				switch {
				case data[p] > 8:
					if 9 <= data[p] && data[p] <= 13 {
						goto tr66
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
					goto st19
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
					goto tr72
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto st19
			}
		default:
			goto st19
		}
		goto tr74
	st_out:
	_test_eof10: cs = 10; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 12:
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
		
		case 13, 15, 17, 18:
//line lex.rl:52

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 19:
//line lex.rl:69

			str()
			emit(Iconst, s)
		
		case 16:
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
		
		case 11:
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
		
//line lex.go:1864
		}
	}

	_out: {}
	}

//line lex.rl:163


	return nil
}
