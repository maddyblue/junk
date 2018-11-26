
//line lex.rl:1
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
	
	//"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

var _ = fmt.Println
var _ = strconv.Atoi

func lexSQL(data []byte) error {

	
//line lex.rl:18
	
//line lex.go:23
const scanner_start int = 41
const scanner_first_final int = 41
const scanner_error int = 0

const scanner_en_main int = 41


//line lex.rl:19

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
		ch byte
		rn rune
		buf *bytes.Buffer
	)
	str := func() { s = string(data[mark:p]) }
        _, _, _, _, _, _, _ = uval, err, isFconst, isUpper, isNotASCII, str, buf

	
//line lex.go:54
	{
	cs = scanner_start
	}

//line lex.go:59
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 41:
		goto st_case_41
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 42:
		goto st_case_42
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 43:
		goto st_case_43
	case 4:
		goto st_case_4
	case 44:
		goto st_case_44
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 45:
		goto st_case_45
	case 8:
		goto st_case_8
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 13:
		goto st_case_13
	case 53:
		goto st_case_53
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	}
	goto st_out
tr110:
//line lex.rl:126

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
		
	goto st41
tr123:
//line lex.rl:44

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
		
	goto st41
tr134:
//line lex.rl:160

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
		
	goto st41
tr147:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st41
tr160:
//line lex.rl:77

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
		
	goto st41
tr170:
//line lex.rl:235

			emit(Bconst, buf.String())
		
	goto st41
tr183:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:299
		switch data[p] {
		case 32:
			goto st41
		case 34:
			goto tr98
		case 36:
			goto tr99
		case 39:
			goto tr100
		case 43:
			goto tr101
		case 45:
			goto tr102
		case 46:
			goto tr103
		case 48:
			goto tr104
		case 95:
			goto tr107
		case 98:
			goto tr108
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
						goto tr105
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st41
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr106
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr107
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr109
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line lex.rl:106
 isNotASCII = true 
	goto st1
tr3:
//line lex.rl:120
 numQuote++ 
	goto st1
tr5:
//line lex.rl:120
 numQuote++ 
//line lex.rl:106
 isNotASCII = true 
	goto st1
tr98:
//line lex.rl:41
 mark = p 
	goto st1
tr124:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st1
tr135:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st1
tr148:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st1
tr161:
//line lex.rl:77

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
		
//line lex.rl:41
 mark = p 
	goto st1
tr171:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st1
tr184:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:493
		if data[p] == 34 {
			goto st42
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:120
 numQuote++ 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line lex.go:510
		switch data[p] {
		case 32:
			goto tr110
		case 34:
			goto st2
		case 36:
			goto tr112
		case 39:
			goto tr113
		case 43:
			goto tr114
		case 45:
			goto tr115
		case 46:
			goto tr116
		case 48:
			goto tr117
		case 95:
			goto tr120
		case 98:
			goto tr121
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
						goto tr118
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr110
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr119
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr120
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr122
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
tr99:
//line lex.rl:41
 mark = p 
	goto st3
tr112:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st3
tr125:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st3
tr136:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st3
tr149:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st3
tr172:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st3
tr185:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:703
		if 48 <= data[p] && data[p] <= 57 {
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr123
		case 34:
			goto tr124
		case 36:
			goto tr125
		case 39:
			goto tr126
		case 43:
			goto tr127
		case 45:
			goto tr128
		case 46:
			goto tr129
		case 95:
			goto tr131
		case 98:
			goto tr132
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
						goto st43
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr123
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr130
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr131
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr133
tr10:
//line lex.rl:106
 isNotASCII = true 
	goto st4
tr11:
//line lex.rl:154
 numQuote++ 
	goto st4
tr13:
//line lex.rl:154
 numQuote++ 
//line lex.rl:106
 isNotASCII = true 
	goto st4
tr100:
//line lex.rl:41
 mark = p 
	goto st4
tr113:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st4
tr126:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st4
tr150:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st4
tr163:
//line lex.rl:77

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
		
//line lex.rl:41
 mark = p 
	goto st4
tr186:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line lex.go:892
		if data[p] == 39 {
			goto st44
		}
		if data[p] <= 127 {
			goto st4
		}
		goto tr10
tr12:
//line lex.rl:154
 numQuote++ 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lex.go:909
		switch data[p] {
		case 32:
			goto tr134
		case 34:
			goto tr135
		case 36:
			goto tr136
		case 39:
			goto st5
		case 43:
			goto tr138
		case 45:
			goto tr139
		case 46:
			goto tr140
		case 48:
			goto tr141
		case 95:
			goto tr144
		case 98:
			goto tr145
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
						goto tr142
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr134
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr143
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr144
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr146
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
tr101:
//line lex.rl:41
 mark = p 
	goto st6
tr114:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st6
tr127:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st6
tr138:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st6
tr151:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st6
tr164:
//line lex.rl:77

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
		
//line lex.rl:41
 mark = p 
	goto st6
tr173:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st6
tr187:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line lex.go:1127
		if data[p] == 46 {
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
tr14:
//line lex.rl:70

			isFconst = true
		
	goto st7
tr103:
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr116:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr129:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr140:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr153:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr166:
//line lex.rl:77

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr175:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr189:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line lex.go:1315
		if 48 <= data[p] && data[p] <= 57 {
			goto st45
		}
		goto st0
tr159:
//line lex.rl:70

			isFconst = true
		
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:1331
		switch data[p] {
		case 32:
			goto tr147
		case 34:
			goto tr148
		case 36:
			goto tr149
		case 39:
			goto tr150
		case 43:
			goto tr151
		case 45:
			goto tr152
		case 46:
			goto tr153
		case 69:
			goto tr155
		case 95:
			goto tr156
		case 98:
			goto tr157
		case 101:
			goto tr155
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
						goto st45
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr147
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr154
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr156
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr158
tr102:
//line lex.rl:41
 mark = p 
	goto st8
tr115:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st8
tr128:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st8
tr139:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st8
tr152:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st8
tr165:
//line lex.rl:77

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
		
//line lex.rl:41
 mark = p 
	goto st8
tr174:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st8
tr188:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:1539
		switch data[p] {
		case 45:
			goto st46
		case 46:
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 10 {
			goto st41
		}
		goto st46
tr15:
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st47
tr105:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st47
tr118:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st47
tr142:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st47
tr177:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:1668
		switch data[p] {
		case 32:
			goto tr147
		case 34:
			goto tr148
		case 36:
			goto tr149
		case 39:
			goto tr150
		case 43:
			goto tr151
		case 45:
			goto tr152
		case 46:
			goto tr159
		case 69:
			goto tr155
		case 95:
			goto tr156
		case 98:
			goto tr157
		case 101:
			goto tr155
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
				goto tr147
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr154
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr156
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr158
tr168:
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr107:
//line lex.rl:41
 mark = p 
	goto st48
tr106:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr109:
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr120:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st48
tr119:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr122:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr131:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st48
tr130:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr133:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr144:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st48
tr143:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr146:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr156:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st48
tr154:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr158:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr167:
//line lex.rl:109
 isUpper = true 
	goto st48
tr179:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st48
tr178:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr181:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
tr191:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st48
tr190:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st48
tr192:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lex.go:2111
		switch data[p] {
		case 32:
			goto tr160
		case 34:
			goto tr161
		case 36:
			goto st48
		case 39:
			goto tr163
		case 43:
			goto tr164
		case 45:
			goto tr165
		case 46:
			goto tr166
		case 95:
			goto st48
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
						goto st48
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr160
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr167
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr168
tr155:
//line lex.rl:70

			isFconst = true
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:2182
		switch data[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st49
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 32:
			goto tr147
		case 34:
			goto tr148
		case 36:
			goto tr149
		case 39:
			goto tr150
		case 43:
			goto tr151
		case 45:
			goto tr152
		case 46:
			goto tr153
		case 95:
			goto tr156
		case 98:
			goto tr157
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
						goto st49
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr147
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr154
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr156
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr158
tr108:
//line lex.rl:41
 mark = p 
	goto st50
tr121:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
	goto st50
tr132:
//line lex.rl:44

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
		
//line lex.rl:41
 mark = p 
	goto st50
tr145:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
	goto st50
tr157:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st50
tr180:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:2376
		switch data[p] {
		case 32:
			goto tr160
		case 34:
			goto tr161
		case 36:
			goto st48
		case 39:
			goto st11
		case 43:
			goto tr164
		case 45:
			goto tr165
		case 46:
			goto tr166
		case 95:
			goto st48
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
						goto st48
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr160
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr167
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr168
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 39:
			goto tr21
		case 92:
			goto tr22
		}
		if data[p] <= 127 {
			goto tr20
		}
		goto tr23
tr20:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr23:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr24:
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr27:
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr29:
//line lex.rl:230
 buf.WriteByte(data[p]) 
	goto st12
tr43:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr46:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr55:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr58:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr61:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr64:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr65:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr68:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr69:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr72:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr73:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr76:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr77:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr80:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr81:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr84:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr85:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr88:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr93:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
tr96:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:106
 isNotASCII = true 
//line lex.rl:224
 buf.WriteByte(data[p]) 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:2624
		switch data[p] {
		case 39:
			goto st51
		case 92:
			goto st14
		}
		if data[p] <= 127 {
			goto tr24
		}
		goto tr27
tr21:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st51
tr44:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st51
tr56:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st51
tr62:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st51
tr66:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st51
tr70:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st51
tr74:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st51
tr78:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st51
tr82:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st51
tr86:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st51
tr94:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line lex.go:2684
		switch data[p] {
		case 32:
			goto tr170
		case 34:
			goto tr171
		case 36:
			goto tr172
		case 39:
			goto tr24
		case 43:
			goto tr173
		case 45:
			goto tr174
		case 46:
			goto tr175
		case 48:
			goto tr176
		case 95:
			goto tr179
		case 98:
			goto tr180
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
						goto tr177
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr170
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr178
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr179
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr181
tr104:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st52
tr117:
//line lex.rl:126

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st52
tr141:
//line lex.rl:160

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
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st52
tr176:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line lex.go:2849
		switch data[p] {
		case 32:
			goto tr147
		case 34:
			goto tr148
		case 36:
			goto tr149
		case 39:
			goto tr150
		case 43:
			goto tr151
		case 45:
			goto tr152
		case 46:
			goto tr159
		case 69:
			goto tr155
		case 95:
			goto tr156
		case 98:
			goto tr157
		case 101:
			goto tr155
		case 120:
			goto st13
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
				goto tr147
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr154
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr156
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr158
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st53
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st53
			}
		default:
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 32:
			goto tr183
		case 34:
			goto tr184
		case 36:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 95:
			goto tr191
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 14:
				switch {
				case data[p] > 8:
					if 9 <= data[p] && data[p] <= 13 {
						goto tr183
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
					goto st53
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
					goto tr190
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 103:
					goto tr191
				}
			default:
				goto st53
			}
		default:
			goto st53
		}
		goto tr192
tr22:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st14
tr45:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st14
tr57:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st14
tr63:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st14
tr67:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st14
tr71:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st14
tr75:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st14
tr79:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st14
tr83:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st14
tr87:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st14
tr95:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line lex.go:3058
		switch data[p] {
		case 85:
			goto tr31
		case 88:
			goto st27
		case 97:
			goto st30
		case 98:
			goto st31
		case 102:
			goto st32
		case 110:
			goto st33
		case 114:
			goto st34
		case 116:
			goto st35
		case 117:
			goto tr39
		case 118:
			goto st40
		case 120:
			goto st27
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr30
		}
		goto tr29
tr30:
//line lex.rl:213
 ch = 0 
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:3098
		if 48 <= data[p] && data[p] <= 55 {
			goto tr41
		}
		goto st0
tr41:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:3112
		if 48 <= data[p] && data[p] <= 55 {
			goto tr42
		}
		goto st0
tr42:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line lex.go:3126
		switch data[p] {
		case 39:
			goto tr44
		case 92:
			goto tr45
		}
		if data[p] <= 127 {
			goto tr43
		}
		goto tr46
tr31:
//line lex.rl:208
 rn = 0 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:3146
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr47
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr47
			}
		default:
			goto tr47
		}
		goto st0
tr47:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:3169
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr48
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr48
			}
		default:
			goto tr48
		}
		goto st0
tr48:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:3192
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr49
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr49
			}
		default:
			goto tr49
		}
		goto st0
tr49:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:3215
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr50
			}
		default:
			goto tr50
		}
		goto st0
tr50:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line lex.go:3238
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr51
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
tr51:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:3261
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto st0
tr52:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line lex.go:3284
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr53
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr53
			}
		default:
			goto tr53
		}
		goto st0
tr53:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line lex.go:3307
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr54
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto st0
tr54:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st26
tr92:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:3334
		switch data[p] {
		case 39:
			goto tr56
		case 92:
			goto tr57
		}
		if data[p] <= 127 {
			goto tr55
		}
		goto tr58
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr59
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr59
			}
		default:
			goto tr59
		}
		goto st0
tr59:
//line lex.rl:196
 ch = 0 
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line lex.go:3374
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr60
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr60
			}
		default:
			goto tr60
		}
		goto st0
tr60:
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line lex.go:3397
		switch data[p] {
		case 39:
			goto tr62
		case 92:
			goto tr63
		}
		if data[p] <= 127 {
			goto tr61
		}
		goto tr64
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 39:
			goto tr66
		case 92:
			goto tr67
		}
		if data[p] <= 127 {
			goto tr65
		}
		goto tr68
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 39:
			goto tr70
		case 92:
			goto tr71
		}
		if data[p] <= 127 {
			goto tr69
		}
		goto tr72
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr74
		case 92:
			goto tr75
		}
		if data[p] <= 127 {
			goto tr73
		}
		goto tr76
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 39:
			goto tr78
		case 92:
			goto tr79
		}
		if data[p] <= 127 {
			goto tr77
		}
		goto tr80
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 39:
			goto tr82
		case 92:
			goto tr83
		}
		if data[p] <= 127 {
			goto tr81
		}
		goto tr84
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 39:
			goto tr86
		case 92:
			goto tr87
		}
		if data[p] <= 127 {
			goto tr85
		}
		goto tr88
tr39:
//line lex.rl:208
 rn = 0 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line lex.go:3507
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr89
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr89
			}
		default:
			goto tr89
		}
		goto st0
tr89:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line lex.go:3530
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr90
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr90
			}
		default:
			goto tr90
		}
		goto st0
tr90:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line lex.go:3553
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr91
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr91
			}
		default:
			goto tr91
		}
		goto st0
tr91:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:3576
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr92
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr92
			}
		default:
			goto tr92
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 39:
			goto tr94
		case 92:
			goto tr95
		}
		if data[p] <= 127 {
			goto tr93
		}
		goto tr96
	st_out:
	_test_eof41: cs = 41; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 43:
//line lex.rl:44

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
		
		case 45, 47, 49, 52:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 53:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
		case 48, 50:
//line lex.rl:77

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
		
		case 42:
//line lex.rl:126

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
		
		case 44:
//line lex.rl:160

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
		
		case 51:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.go:3775
		}
	}

	_out: {}
	}

//line lex.rl:256


	return nil
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
