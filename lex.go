
//line lex.rl:1
package main

import (
	"fmt"
	"strconv"
	
	//"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

func lexSQL(data []byte) error {

	
//line lex.rl:13
	
//line lex.go:18
const scanner_start int = 8
const scanner_first_final int = 8
const scanner_error int = 0

const scanner_en_main int = 8


//line lex.rl:14

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
        
	
//line lex.go:43
	{
	cs = scanner_start
	}

//line lex.go:48
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 8:
		goto st_case_8
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 9:
		goto st_case_9
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 10:
		goto st_case_10
	case 4:
		goto st_case_4
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 7:
		goto st_case_7
	case 16:
		goto st_case_16
	}
	goto st_out
tr18:
//line lex.rl:33

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
		
	goto st8
tr25:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st8
tr34:
//line lex.rl:66

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
		
	goto st8
tr41:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:152
		switch data[p] {
		case 32:
			goto st8
		case 36:
			goto tr10
		case 43:
			goto tr11
		case 45:
			goto tr12
		case 46:
			goto tr13
		case 48:
			goto tr14
		case 96:
			goto st0
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto st8
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr16
			}
		default:
			goto tr15
		}
		goto tr17
st_case_0:
	st0:
		cs = 0
		goto _out
tr10:
//line lex.rl:30
 mark = p 
	goto st1
tr19:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
	goto st1
tr26:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
	goto st1
tr42:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:258
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 32:
			goto tr18
		case 36:
			goto tr19
		case 43:
			goto tr20
		case 45:
			goto tr21
		case 46:
			goto tr22
		case 96:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto tr18
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr23
			}
		default:
			goto st9
		}
		goto tr24
tr11:
//line lex.rl:30
 mark = p 
	goto st2
tr20:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
	goto st2
tr27:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
	goto st2
tr36:
//line lex.rl:66

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
		
//line lex.rl:30
 mark = p 
	goto st2
tr43:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line lex.go:392
		if data[p] == 46 {
			goto tr2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
tr2:
//line lex.rl:59

			isFconst = true
		
	goto st3
tr13:
//line lex.rl:30
 mark = p 
//line lex.rl:59

			isFconst = true
		
	goto st3
tr22:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
//line lex.rl:59

			isFconst = true
		
	goto st3
tr29:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
//line lex.rl:59

			isFconst = true
		
	goto st3
tr38:
//line lex.rl:66

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
		
//line lex.rl:30
 mark = p 
//line lex.rl:59

			isFconst = true
		
	goto st3
tr45:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
//line lex.rl:59

			isFconst = true
		
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:500
		if 48 <= data[p] && data[p] <= 57 {
			goto st10
		}
		goto st0
tr33:
//line lex.rl:59

			isFconst = true
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line lex.go:516
		switch data[p] {
		case 32:
			goto tr25
		case 36:
			goto tr26
		case 43:
			goto tr27
		case 45:
			goto tr28
		case 46:
			goto tr29
		case 69:
			goto tr31
		case 96:
			goto st0
		case 101:
			goto tr31
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr30
			}
		default:
			goto st10
		}
		goto tr32
tr12:
//line lex.rl:30
 mark = p 
	goto st4
tr21:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
	goto st4
tr28:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
	goto st4
tr37:
//line lex.rl:66

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
		
//line lex.rl:30
 mark = p 
	goto st4
tr44:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line lex.go:645
		switch data[p] {
		case 45:
			goto st11
		case 46:
			goto tr2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 10 {
			goto st8
		}
		goto st11
tr3:
//line lex.rl:54

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st12
tr15:
//line lex.rl:30
 mark = p 
//line lex.rl:54

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:688
		switch data[p] {
		case 32:
			goto tr25
		case 36:
			goto tr26
		case 43:
			goto tr27
		case 45:
			goto tr28
		case 46:
			goto tr33
		case 69:
			goto tr31
		case 96:
			goto st0
		case 101:
			goto tr31
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr30
			}
		default:
			goto tr3
		}
		goto tr32
tr17:
//line lex.rl:30
 mark = p 
	goto st13
tr16:
//line lex.rl:30
 mark = p 
//line lex.rl:97
 isUpper = true 
	goto st13
tr24:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
	goto st13
tr23:
//line lex.rl:33

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
		
//line lex.rl:30
 mark = p 
//line lex.rl:97
 isUpper = true 
	goto st13
tr32:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
	goto st13
tr30:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:30
 mark = p 
//line lex.rl:97
 isUpper = true 
	goto st13
tr39:
//line lex.rl:97
 isUpper = true 
	goto st13
tr47:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
	goto st13
tr46:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
//line lex.rl:30
 mark = p 
//line lex.rl:97
 isUpper = true 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:848
		switch data[p] {
		case 32:
			goto tr34
		case 43:
			goto tr36
		case 45:
			goto tr37
		case 46:
			goto tr38
		case 96:
			goto st0
		}
		switch {
		case data[p] < 37:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 35 {
					goto st0
				}
			default:
				goto tr34
			}
		case data[p] > 47:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr39
			}
		default:
			goto st0
		}
		goto st13
tr31:
//line lex.rl:59

			isFconst = true
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line lex.go:908
		switch data[p] {
		case 43:
			goto st6
		case 45:
			goto st6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
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
			goto tr25
		case 36:
			goto tr26
		case 43:
			goto tr27
		case 45:
			goto tr28
		case 46:
			goto tr29
		case 96:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr30
			}
		default:
			goto st14
		}
		goto tr32
tr14:
//line lex.rl:30
 mark = p 
//line lex.rl:54

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:998
		switch data[p] {
		case 32:
			goto tr25
		case 36:
			goto tr26
		case 43:
			goto tr27
		case 45:
			goto tr28
		case 46:
			goto tr33
		case 69:
			goto tr31
		case 96:
			goto st0
		case 101:
			goto tr31
		case 120:
			goto st7
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto st0
				}
			case data[p] > 13:
				if 14 <= data[p] && data[p] <= 47 {
					goto st0
				}
			default:
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] > 90:
				switch {
				case data[p] > 94:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 91:
					goto st0
				}
			default:
				goto tr30
			}
		default:
			goto tr3
		}
		goto tr32
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st16
			}
		default:
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 32:
			goto tr41
		case 36:
			goto tr42
		case 43:
			goto tr43
		case 45:
			goto tr44
		case 46:
			goto tr45
		case 96:
			goto st0
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
				goto tr41
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 70:
					if 71 <= data[p] && data[p] <= 90 {
						goto tr46
					}
				case data[p] >= 65:
					goto st16
				}
			case data[p] > 94:
				switch {
				case data[p] > 102:
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
		goto tr47
	st_out:
	_test_eof8: cs = 8; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 9:
//line lex.rl:33

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
		
		case 10, 12, 14, 15:
//line lex.rl:45

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 16:
//line lex.rl:62

			str()
			emit(Iconst, s)
		
		case 13:
//line lex.rl:66

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
		
//line lex.go:1212
		}
	}

	_out: {}
	}

//line lex.rl:120


	return nil
}
