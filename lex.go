
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
const scanner_start int = 70
const scanner_first_final int = 70
const scanner_error int = 0

const scanner_en_main int = 70


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
	case 70:
		goto st_case_70
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 71:
		goto st_case_71
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 72:
		goto st_case_72
	case 4:
		goto st_case_4
	case 73:
		goto st_case_73
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 74:
		goto st_case_74
	case 8:
		goto st_case_8
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 13:
		goto st_case_13
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 84:
		goto st_case_84
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
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	}
	goto st_out
tr163:
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
		
	goto st70
tr177:
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
		
	goto st70
tr189:
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
		
	goto st70
tr203:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st70
tr216:
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
		
	goto st70
tr227:
//line lex.rl:235

			emit(Bconst, buf.String())
		
	goto st70
tr241:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
	goto st70
tr252:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line lex.go:370
		switch data[p] {
		case 32:
			goto st70
		case 34:
			goto tr150
		case 36:
			goto tr151
		case 39:
			goto tr152
		case 43:
			goto tr153
		case 45:
			goto tr154
		case 46:
			goto tr155
		case 48:
			goto tr156
		case 95:
			goto tr159
		case 98:
			goto tr160
		case 101:
			goto tr161
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
						goto tr157
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st70
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr158
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr159
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr162
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
tr150:
//line lex.rl:41
 mark = p 
	goto st1
tr178:
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
tr190:
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
tr204:
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
tr217:
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
tr228:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st1
tr242:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st1
tr253:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:577
		if data[p] == 34 {
			goto st71
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:120
 numQuote++ 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lex.go:594
		switch data[p] {
		case 32:
			goto tr163
		case 34:
			goto st2
		case 36:
			goto tr165
		case 39:
			goto tr166
		case 43:
			goto tr167
		case 45:
			goto tr168
		case 46:
			goto tr169
		case 48:
			goto tr170
		case 95:
			goto tr173
		case 98:
			goto tr174
		case 101:
			goto tr175
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
						goto tr171
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr163
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr172
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr173
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr176
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
tr151:
//line lex.rl:41
 mark = p 
	goto st3
tr165:
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
tr179:
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
tr191:
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
tr205:
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
tr229:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st3
tr243:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st3
tr254:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:800
		if 48 <= data[p] && data[p] <= 57 {
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 32:
			goto tr177
		case 34:
			goto tr178
		case 36:
			goto tr179
		case 39:
			goto tr180
		case 43:
			goto tr181
		case 45:
			goto tr182
		case 46:
			goto tr183
		case 95:
			goto tr185
		case 98:
			goto tr186
		case 101:
			goto tr187
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
						goto st72
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr177
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr185
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr188
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
tr152:
//line lex.rl:41
 mark = p 
	goto st4
tr166:
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
tr180:
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
tr206:
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
tr219:
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
tr244:
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
//line lex.go:991
		if data[p] == 39 {
			goto st73
		}
		if data[p] <= 127 {
			goto st4
		}
		goto tr10
tr12:
//line lex.rl:154
 numQuote++ 
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lex.go:1008
		switch data[p] {
		case 32:
			goto tr189
		case 34:
			goto tr190
		case 36:
			goto tr191
		case 39:
			goto st5
		case 43:
			goto tr193
		case 45:
			goto tr194
		case 46:
			goto tr195
		case 48:
			goto tr196
		case 95:
			goto tr199
		case 98:
			goto tr200
		case 101:
			goto tr201
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
						goto tr197
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr189
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr198
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr199
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr202
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
tr153:
//line lex.rl:41
 mark = p 
	goto st6
tr167:
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
tr181:
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
tr193:
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
tr207:
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
tr220:
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
tr230:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st6
tr245:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st6
tr255:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line lex.go:1239
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
tr155:
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr169:
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
tr183:
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
tr195:
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
tr209:
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
tr222:
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
tr232:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr247:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr257:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
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
//line lex.go:1442
		if 48 <= data[p] && data[p] <= 57 {
			goto st74
		}
		goto st0
tr215:
//line lex.rl:70

			isFconst = true
		
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lex.go:1458
		switch data[p] {
		case 32:
			goto tr203
		case 34:
			goto tr204
		case 36:
			goto tr205
		case 39:
			goto tr206
		case 43:
			goto tr207
		case 45:
			goto tr208
		case 46:
			goto tr209
		case 69:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr213
		case 101:
			goto tr211
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
						goto st74
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr203
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr210
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr212
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr214
tr154:
//line lex.rl:41
 mark = p 
	goto st8
tr168:
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
tr182:
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
tr194:
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
tr208:
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
tr221:
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
tr231:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st8
tr246:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st8
tr256:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:1677
		switch data[p] {
		case 45:
			goto st75
		case 46:
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 10 {
			goto st70
		}
		goto st75
tr15:
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st76
tr157:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st76
tr171:
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
		
	goto st76
tr197:
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
		
	goto st76
tr234:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st76
tr259:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lex.go:1823
		switch data[p] {
		case 32:
			goto tr203
		case 34:
			goto tr204
		case 36:
			goto tr205
		case 39:
			goto tr206
		case 43:
			goto tr207
		case 45:
			goto tr208
		case 46:
			goto tr215
		case 69:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr213
		case 101:
			goto tr211
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
				goto tr203
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr210
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr212
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr214
tr224:
//line lex.rl:106
 isNotASCII = true 
	goto st77
tr159:
//line lex.rl:41
 mark = p 
	goto st77
tr158:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st77
tr162:
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st77
tr173:
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
	goto st77
tr172:
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
	goto st77
tr176:
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
	goto st77
tr185:
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
	goto st77
tr184:
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
	goto st77
tr188:
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
	goto st77
tr199:
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
	goto st77
tr198:
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
	goto st77
tr202:
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
	goto st77
tr212:
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
	goto st77
tr210:
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
	goto st77
tr214:
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
	goto st77
tr223:
//line lex.rl:109
 isUpper = true 
	goto st77
tr236:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st77
tr235:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st77
tr239:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st77
tr249:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st77
tr248:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st77
tr250:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st77
tr261:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st77
tr260:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st77
tr264:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lex.go:2303
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto st77
		case 39:
			goto tr219
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr222
		case 95:
			goto st77
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
						goto st77
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr216
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr223
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st77
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr224
tr211:
//line lex.rl:70

			isFconst = true
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:2374
		switch data[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st78
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto st78
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 32:
			goto tr203
		case 34:
			goto tr204
		case 36:
			goto tr205
		case 39:
			goto tr206
		case 43:
			goto tr207
		case 45:
			goto tr208
		case 46:
			goto tr209
		case 95:
			goto tr212
		case 98:
			goto tr213
		case 101:
			goto tr225
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
						goto st78
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr203
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr210
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr212
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr214
tr160:
//line lex.rl:41
 mark = p 
	goto st79
tr174:
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
	goto st79
tr186:
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
	goto st79
tr200:
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
	goto st79
tr213:
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
	goto st79
tr237:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st79
tr262:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lex.go:2581
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto st77
		case 39:
			goto st11
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr222
		case 95:
			goto st77
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
						goto st77
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr216
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr223
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st77
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr224
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
		goto tr20
tr20:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr23:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr91:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st12
tr105:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr116:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr121:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr124:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr127:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr130:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr133:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr136:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr139:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
tr146:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:2732
		switch data[p] {
		case 39:
			goto st80
		case 92:
			goto st43
		}
		goto tr23
tr21:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st80
tr106:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st80
tr117:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st80
tr122:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st80
tr125:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st80
tr128:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st80
tr131:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st80
tr134:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st80
tr137:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st80
tr140:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st80
tr147:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lex.go:2789
		switch data[p] {
		case 32:
			goto tr227
		case 34:
			goto tr228
		case 36:
			goto tr229
		case 39:
			goto tr23
		case 43:
			goto tr230
		case 45:
			goto tr231
		case 46:
			goto tr232
		case 48:
			goto tr233
		case 95:
			goto tr236
		case 98:
			goto tr237
		case 101:
			goto tr238
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
						goto tr234
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr227
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr235
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr236
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr239
tr156:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st81
tr170:
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
		
	goto st81
tr196:
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
		
	goto st81
tr233:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st81
tr258:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lex.go:2973
		switch data[p] {
		case 32:
			goto tr203
		case 34:
			goto tr204
		case 36:
			goto tr205
		case 39:
			goto tr206
		case 43:
			goto tr207
		case 45:
			goto tr208
		case 46:
			goto tr215
		case 69:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr213
		case 101:
			goto tr211
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
				goto tr203
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr210
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr212
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr214
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st82
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st82
			}
		default:
			goto st82
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto tr241
		case 34:
			goto tr242
		case 36:
			goto tr243
		case 39:
			goto tr244
		case 43:
			goto tr245
		case 45:
			goto tr246
		case 46:
			goto tr247
		case 95:
			goto tr249
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 14:
				switch {
				case data[p] > 8:
					if 9 <= data[p] && data[p] <= 13 {
						goto tr241
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
					goto st82
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
					goto tr248
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 103:
					goto tr249
				}
			default:
				goto st82
			}
		default:
			goto st82
		}
		goto tr250
tr161:
//line lex.rl:41
 mark = p 
	goto st83
tr175:
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
	goto st83
tr187:
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
	goto st83
tr201:
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
	goto st83
tr225:
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
	goto st83
tr238:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st83
tr263:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lex.go:3252
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto st77
		case 39:
			goto st14
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr222
		case 95:
			goto st77
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
						goto st77
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr216
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr223
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st77
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr224
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 39:
			goto tr28
		case 92:
			goto tr29
		}
		goto tr27
tr27:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr30:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr33:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st15
tr47:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr58:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr63:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr66:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr69:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr72:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr75:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr78:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr81:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr88:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:3403
		switch data[p] {
		case 39:
			goto st84
		case 92:
			goto st16
		}
		goto tr30
tr28:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st84
tr48:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st84
tr59:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st84
tr64:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st84
tr67:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st84
tr70:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st84
tr73:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st84
tr76:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st84
tr79:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st84
tr82:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st84
tr89:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line lex.go:3460
		switch data[p] {
		case 32:
			goto tr252
		case 34:
			goto tr253
		case 36:
			goto tr254
		case 39:
			goto tr30
		case 43:
			goto tr255
		case 45:
			goto tr256
		case 46:
			goto tr257
		case 48:
			goto tr258
		case 95:
			goto tr261
		case 98:
			goto tr262
		case 101:
			goto tr263
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
						goto tr259
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr252
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr260
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr261
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr264
tr29:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st16
tr49:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st16
tr60:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st16
tr65:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st16
tr68:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st16
tr71:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st16
tr74:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st16
tr77:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st16
tr80:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st16
tr83:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st16
tr90:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:3575
		switch data[p] {
		case 85:
			goto tr35
		case 88:
			goto st29
		case 97:
			goto st32
		case 98:
			goto st33
		case 102:
			goto st34
		case 110:
			goto st35
		case 114:
			goto st36
		case 116:
			goto st37
		case 117:
			goto tr43
		case 118:
			goto st42
		case 120:
			goto st29
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr34
		}
		goto tr33
tr34:
//line lex.rl:213
 ch = 0 
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line lex.go:3615
		if 48 <= data[p] && data[p] <= 55 {
			goto tr45
		}
		goto st0
tr45:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:3629
		if 48 <= data[p] && data[p] <= 55 {
			goto tr46
		}
		goto st0
tr46:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:3643
		switch data[p] {
		case 39:
			goto tr48
		case 92:
			goto tr49
		}
		goto tr47
tr35:
//line lex.rl:208
 rn = 0 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:3660
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
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:3683
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
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line lex.go:3706
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
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:3729
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
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line lex.go:3752
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
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line lex.go:3775
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr55
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr55
			}
		default:
			goto tr55
		}
		goto st0
tr55:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:3798
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr56
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr56:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line lex.go:3821
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr57
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr57
			}
		default:
			goto tr57
		}
		goto st0
tr57:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st28
tr87:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line lex.go:3848
		switch data[p] {
		case 39:
			goto tr59
		case 92:
			goto tr60
		}
		goto tr58
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr61
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr61
			}
		default:
			goto tr61
		}
		goto st0
tr61:
//line lex.rl:196
 ch = 0 
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line lex.go:3885
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr62
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr62
			}
		default:
			goto tr62
		}
		goto st0
tr62:
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line lex.go:3908
		switch data[p] {
		case 39:
			goto tr64
		case 92:
			goto tr65
		}
		goto tr63
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr67
		case 92:
			goto tr68
		}
		goto tr66
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 39:
			goto tr70
		case 92:
			goto tr71
		}
		goto tr69
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 39:
			goto tr73
		case 92:
			goto tr74
		}
		goto tr72
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 39:
			goto tr76
		case 92:
			goto tr77
		}
		goto tr75
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 39:
			goto tr79
		case 92:
			goto tr80
		}
		goto tr78
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 39:
			goto tr82
		case 92:
			goto tr83
		}
		goto tr81
tr43:
//line lex.rl:208
 rn = 0 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line lex.go:3997
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto st0
tr84:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:4020
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr85
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr85
			}
		default:
			goto tr85
		}
		goto st0
tr85:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lex.go:4043
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr86
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr86
			}
		default:
			goto tr86
		}
		goto st0
tr86:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:4066
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr87
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr87
			}
		default:
			goto tr87
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 39:
			goto tr89
		case 92:
			goto tr90
		}
		goto tr88
tr22:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st43
tr107:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st43
tr118:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st43
tr123:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st43
tr126:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st43
tr129:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st43
tr132:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st43
tr135:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st43
tr138:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st43
tr141:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st43
tr148:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lex.go:4141
		switch data[p] {
		case 85:
			goto tr93
		case 88:
			goto st56
		case 97:
			goto st59
		case 98:
			goto st60
		case 102:
			goto st61
		case 110:
			goto st62
		case 114:
			goto st63
		case 116:
			goto st64
		case 117:
			goto tr101
		case 118:
			goto st69
		case 120:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr92
		}
		goto tr91
tr92:
//line lex.rl:213
 ch = 0 
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lex.go:4181
		if 48 <= data[p] && data[p] <= 55 {
			goto tr103
		}
		goto st0
tr103:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:4195
		if 48 <= data[p] && data[p] <= 55 {
			goto tr104
		}
		goto st0
tr104:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lex.go:4209
		switch data[p] {
		case 39:
			goto tr106
		case 92:
			goto tr107
		}
		goto tr105
tr93:
//line lex.rl:208
 rn = 0 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:4226
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr108
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr108
			}
		default:
			goto tr108
		}
		goto st0
tr108:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lex.go:4249
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr109
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr109
			}
		default:
			goto tr109
		}
		goto st0
tr109:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line lex.go:4272
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr110
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr110
			}
		default:
			goto tr110
		}
		goto st0
tr110:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:4295
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
tr111:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line lex.go:4318
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr112
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr112
			}
		default:
			goto tr112
		}
		goto st0
tr112:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line lex.go:4341
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr113
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr113
			}
		default:
			goto tr113
		}
		goto st0
tr113:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lex.go:4364
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr114
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr114
			}
		default:
			goto tr114
		}
		goto st0
tr114:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line lex.go:4387
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr115
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
tr115:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st55
tr145:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line lex.go:4414
		switch data[p] {
		case 39:
			goto tr117
		case 92:
			goto tr118
		}
		goto tr116
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr119
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr119
			}
		default:
			goto tr119
		}
		goto st0
tr119:
//line lex.rl:196
 ch = 0 
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line lex.go:4451
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr120
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr120
			}
		default:
			goto tr120
		}
		goto st0
tr120:
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line lex.go:4474
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto tr123
		}
		goto tr121
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 39:
			goto tr125
		case 92:
			goto tr126
		}
		goto tr124
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 39:
			goto tr128
		case 92:
			goto tr129
		}
		goto tr127
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 39:
			goto tr131
		case 92:
			goto tr132
		}
		goto tr130
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 39:
			goto tr134
		case 92:
			goto tr135
		}
		goto tr133
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 39:
			goto tr137
		case 92:
			goto tr138
		}
		goto tr136
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 39:
			goto tr140
		case 92:
			goto tr141
		}
		goto tr139
tr101:
//line lex.rl:208
 rn = 0 
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lex.go:4563
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr142
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr142
			}
		default:
			goto tr142
		}
		goto st0
tr142:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lex.go:4586
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr143
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr143
			}
		default:
			goto tr143
		}
		goto st0
tr143:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lex.go:4609
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr144
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr144
			}
		default:
			goto tr144
		}
		goto st0
tr144:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lex.go:4632
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr145
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr145
			}
		default:
			goto tr145
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 39:
			goto tr147
		case 92:
			goto tr148
		}
		goto tr146
	st_out:
	_test_eof70: cs = 70; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
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
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 72:
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
		
		case 74, 76, 78, 81:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 82:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
		case 77, 79, 83:
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
		
		case 71:
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
		
		case 73:
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
		
		case 80:
//line lex.rl:235

			emit(Bconst, buf.String())
		
		case 84:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.go:4867
		}
	}

	_out: {}
	}

//line lex.rl:264


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
