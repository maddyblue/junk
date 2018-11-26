
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
const scanner_start int = 73
const scanner_first_final int = 73
const scanner_error int = 0

const scanner_en_main int = 73


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
	case 73:
		goto st_case_73
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 74:
		goto st_case_74
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 75:
		goto st_case_75
	case 4:
		goto st_case_4
	case 76:
		goto st_case_76
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 77:
		goto st_case_77
	case 8:
		goto st_case_8
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 11:
		goto st_case_11
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 88:
		goto st_case_88
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
	case 70:
		goto st_case_70
	case 89:
		goto st_case_89
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	}
	goto st_out
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
		
	goto st73
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
		
	goto st73
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
		
	goto st73
tr216:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
	goto st73
tr231:
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
		
	goto st73
tr242:
//line lex.rl:235

			emit(Bconst, buf.String())
		
	goto st73
tr261:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
	goto st73
tr276:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lex.go:380
		switch data[p] {
		case 32:
			goto st73
		case 34:
			goto tr155
		case 36:
			goto tr156
		case 39:
			goto tr157
		case 43:
			goto tr158
		case 45:
			goto tr159
		case 46:
			goto tr160
		case 48:
			goto tr161
		case 88:
			goto tr164
		case 95:
			goto tr165
		case 98:
			goto tr166
		case 101:
			goto tr167
		case 120:
			goto tr168
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
						goto tr162
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st73
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr163
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr165
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr169
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
tr155:
//line lex.rl:41
 mark = p 
	goto st1
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
	goto st1
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
	goto st1
tr217:
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
tr232:
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
tr243:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st1
tr262:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st1
tr277:
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
//line lex.go:591
		if data[p] == 34 {
			goto st74
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:120
 numQuote++ 
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lex.go:608
		switch data[p] {
		case 32:
			goto tr170
		case 34:
			goto st2
		case 36:
			goto tr172
		case 39:
			goto tr173
		case 43:
			goto tr174
		case 45:
			goto tr175
		case 46:
			goto tr176
		case 48:
			goto tr177
		case 88:
			goto tr180
		case 95:
			goto tr181
		case 98:
			goto tr182
		case 101:
			goto tr183
		case 120:
			goto tr184
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
						goto tr178
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
					goto tr179
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr181
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr185
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
tr156:
//line lex.rl:41
 mark = p 
	goto st3
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
	goto st3
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
	goto st3
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
	goto st3
tr218:
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
tr244:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st3
tr263:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st3
tr278:
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
//line lex.go:818
		if 48 <= data[p] && data[p] <= 57 {
			goto st75
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 32:
			goto tr186
		case 34:
			goto tr187
		case 36:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 88:
			goto tr194
		case 95:
			goto tr195
		case 98:
			goto tr196
		case 101:
			goto tr197
		case 120:
			goto tr198
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
						goto st75
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr186
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr195
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr199
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
tr157:
//line lex.rl:41
 mark = p 
	goto st4
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
	goto st4
tr189:
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
tr219:
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
tr234:
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
tr245:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st4
tr279:
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
//line lex.go:1021
		if data[p] == 39 {
			goto st76
		}
		if data[p] <= 127 {
			goto st4
		}
		goto tr10
tr12:
//line lex.rl:154
 numQuote++ 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lex.go:1038
		switch data[p] {
		case 32:
			goto tr200
		case 34:
			goto tr201
		case 36:
			goto tr202
		case 39:
			goto st5
		case 43:
			goto tr204
		case 45:
			goto tr205
		case 46:
			goto tr206
		case 48:
			goto tr207
		case 88:
			goto tr210
		case 95:
			goto tr211
		case 98:
			goto tr212
		case 101:
			goto tr213
		case 120:
			goto tr214
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
						goto tr208
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr200
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr209
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr211
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr215
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
tr158:
//line lex.rl:41
 mark = p 
	goto st6
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
	goto st6
tr190:
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
tr204:
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
tr220:
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
tr235:
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
tr246:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st6
tr264:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st6
tr280:
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
//line lex.go:1273
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
tr160:
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
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
//line lex.rl:70

			isFconst = true
		
	goto st7
tr192:
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
tr206:
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
tr222:
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
tr237:
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
tr248:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st7
tr266:
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
tr282:
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
//line lex.go:1476
		if 48 <= data[p] && data[p] <= 57 {
			goto st77
		}
		goto st0
tr230:
//line lex.rl:70

			isFconst = true
		
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lex.go:1492
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto tr218
		case 39:
			goto tr219
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr222
		case 69:
			goto tr224
		case 88:
			goto tr225
		case 95:
			goto tr226
		case 98:
			goto tr227
		case 101:
			goto tr224
		case 120:
			goto tr228
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
					goto tr226
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr229
tr159:
//line lex.rl:41
 mark = p 
	goto st8
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
	goto st8
tr191:
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
tr205:
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
tr221:
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
tr236:
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
tr247:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st8
tr265:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st8
tr281:
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
//line lex.go:1715
		switch data[p] {
		case 45:
			goto st78
		case 46:
			goto tr14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr15
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 10 {
			goto st73
		}
		goto st78
tr15:
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st79
tr162:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st79
tr178:
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
		
	goto st79
tr208:
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
		
	goto st79
tr250:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st79
tr268:
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
		
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lex.go:1861
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto tr218
		case 39:
			goto tr219
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr230
		case 69:
			goto tr224
		case 88:
			goto tr225
		case 95:
			goto tr226
		case 98:
			goto tr227
		case 101:
			goto tr224
		case 120:
			goto tr228
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
					goto tr226
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr229
tr239:
//line lex.rl:106
 isNotASCII = true 
	goto st80
tr165:
//line lex.rl:41
 mark = p 
	goto st80
tr163:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st80
tr169:
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st80
tr181:
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
	goto st80
tr179:
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
	goto st80
tr185:
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
	goto st80
tr195:
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
	goto st80
tr193:
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
	goto st80
tr199:
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
	goto st80
tr211:
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
	goto st80
tr209:
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
	goto st80
tr215:
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
	goto st80
tr226:
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
	goto st80
tr223:
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
	goto st80
tr229:
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
	goto st80
tr238:
//line lex.rl:109
 isUpper = true 
	goto st80
tr253:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st80
tr251:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st80
tr257:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st80
tr271:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st80
tr269:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st80
tr275:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st80
tr285:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st80
tr283:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st80
tr287:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lex.go:2345
		switch data[p] {
		case 32:
			goto tr231
		case 34:
			goto tr232
		case 36:
			goto st80
		case 39:
			goto tr234
		case 43:
			goto tr235
		case 45:
			goto tr236
		case 46:
			goto tr237
		case 95:
			goto st80
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
						goto st80
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr231
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr238
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st80
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr239
tr224:
//line lex.rl:70

			isFconst = true
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:2416
		switch data[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st81
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto st81
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto tr218
		case 39:
			goto tr219
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr222
		case 88:
			goto tr225
		case 95:
			goto tr226
		case 98:
			goto tr227
		case 101:
			goto tr240
		case 120:
			goto tr228
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
						goto st81
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
					goto tr226
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr229
tr168:
//line lex.rl:41
 mark = p 
	goto st82
tr164:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st82
tr184:
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
	goto st82
tr180:
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
	goto st82
tr198:
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
	goto st82
tr194:
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
	goto st82
tr214:
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
	goto st82
tr210:
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
	goto st82
tr228:
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
	goto st82
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
//line lex.rl:109
 isUpper = true 
	goto st82
tr256:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st82
tr252:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st82
tr274:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st82
tr270:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st82
tr286:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
	goto st82
tr284:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lex.go:2775
		switch data[p] {
		case 32:
			goto tr231
		case 34:
			goto tr232
		case 36:
			goto st80
		case 39:
			goto tr241
		case 43:
			goto tr235
		case 45:
			goto tr236
		case 46:
			goto tr237
		case 95:
			goto st80
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
						goto st80
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr231
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr238
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st80
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr239
tr241:
//line lex.rl:246
 buf = new(bytes.Buffer) 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line lex.go:2844
		if data[p] == 39 {
			goto st83
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto st0
tr152:
//line lex.rl:252
 buf.WriteByte(ch) 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lex.go:2870
		switch data[p] {
		case 32:
			goto tr242
		case 34:
			goto tr243
		case 36:
			goto tr244
		case 39:
			goto tr245
		case 43:
			goto tr246
		case 45:
			goto tr247
		case 46:
			goto tr248
		case 48:
			goto tr249
		case 88:
			goto tr252
		case 95:
			goto tr253
		case 98:
			goto tr254
		case 101:
			goto tr255
		case 120:
			goto tr256
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
						goto tr250
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr242
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr251
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr253
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr257
tr161:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st84
tr177:
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
		
	goto st84
tr207:
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
		
	goto st84
tr249:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st84
tr267:
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
		
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line lex.go:3058
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 36:
			goto tr218
		case 39:
			goto tr219
		case 43:
			goto tr220
		case 45:
			goto tr221
		case 46:
			goto tr230
		case 69:
			goto tr224
		case 88:
			goto tr225
		case 95:
			goto tr226
		case 98:
			goto tr227
		case 101:
			goto tr224
		case 120:
			goto st70
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
					goto tr226
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr229
tr166:
//line lex.rl:41
 mark = p 
	goto st85
tr182:
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
	goto st85
tr196:
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
	goto st85
tr212:
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
	goto st85
tr227:
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
	goto st85
tr254:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st85
tr272:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lex.go:3247
		switch data[p] {
		case 32:
			goto tr231
		case 34:
			goto tr232
		case 36:
			goto st80
		case 39:
			goto st12
		case 43:
			goto tr235
		case 45:
			goto tr236
		case 46:
			goto tr237
		case 95:
			goto st80
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
						goto st80
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr231
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr238
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st80
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr239
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 39:
			goto tr23
		case 92:
			goto tr24
		}
		goto tr22
tr22:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr25:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr92:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st13
tr106:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr117:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr122:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr125:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr128:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr131:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr134:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr137:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr140:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
tr147:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:3398
		switch data[p] {
		case 39:
			goto st86
		case 92:
			goto st43
		}
		goto tr25
tr23:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st86
tr107:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st86
tr118:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st86
tr123:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st86
tr126:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st86
tr129:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st86
tr132:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st86
tr135:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st86
tr138:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st86
tr141:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st86
tr148:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line lex.go:3455
		switch data[p] {
		case 32:
			goto tr242
		case 34:
			goto tr243
		case 36:
			goto tr244
		case 39:
			goto tr25
		case 43:
			goto tr246
		case 45:
			goto tr247
		case 46:
			goto tr248
		case 48:
			goto tr249
		case 88:
			goto tr252
		case 95:
			goto tr253
		case 98:
			goto tr254
		case 101:
			goto tr255
		case 120:
			goto tr256
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
						goto tr250
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr242
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr251
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr253
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr257
tr167:
//line lex.rl:41
 mark = p 
	goto st87
tr183:
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
	goto st87
tr197:
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
	goto st87
tr213:
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
	goto st87
tr240:
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
	goto st87
tr255:
//line lex.rl:235

			emit(Bconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st87
tr273:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line lex.go:3644
		switch data[p] {
		case 32:
			goto tr231
		case 34:
			goto tr232
		case 36:
			goto st80
		case 39:
			goto st14
		case 43:
			goto tr235
		case 45:
			goto tr236
		case 46:
			goto tr237
		case 95:
			goto st80
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
						goto st80
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr231
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr238
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto st80
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr239
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 39:
			goto tr29
		case 92:
			goto tr30
		}
		goto tr28
tr28:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr31:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr34:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st15
tr48:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr59:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr64:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr67:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr70:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr73:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr76:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr79:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr82:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st15
tr89:
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
//line lex.go:3795
		switch data[p] {
		case 39:
			goto st88
		case 92:
			goto st16
		}
		goto tr31
tr29:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st88
tr49:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st88
tr60:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st88
tr65:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st88
tr68:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st88
tr71:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st88
tr74:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st88
tr77:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st88
tr80:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st88
tr83:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st88
tr90:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line lex.go:3852
		switch data[p] {
		case 32:
			goto tr261
		case 34:
			goto tr262
		case 36:
			goto tr263
		case 39:
			goto tr31
		case 43:
			goto tr264
		case 45:
			goto tr265
		case 46:
			goto tr266
		case 48:
			goto tr267
		case 88:
			goto tr270
		case 95:
			goto tr271
		case 98:
			goto tr272
		case 101:
			goto tr273
		case 120:
			goto tr274
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
						goto tr268
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr261
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr269
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 97:
					goto tr271
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr275
tr30:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st16
tr50:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st16
tr61:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st16
tr66:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st16
tr69:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st16
tr72:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st16
tr75:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st16
tr78:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st16
tr81:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st16
tr84:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st16
tr91:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:3971
		switch data[p] {
		case 85:
			goto tr36
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
			goto tr44
		case 118:
			goto st42
		case 120:
			goto st29
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr35
		}
		goto tr34
tr35:
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
//line lex.go:4011
		if 48 <= data[p] && data[p] <= 55 {
			goto tr46
		}
		goto st0
tr46:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:4025
		if 48 <= data[p] && data[p] <= 55 {
			goto tr47
		}
		goto st0
tr47:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:4039
		switch data[p] {
		case 39:
			goto tr49
		case 92:
			goto tr50
		}
		goto tr48
tr36:
//line lex.rl:208
 rn = 0 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:4056
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
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:4079
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
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line lex.go:4102
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
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:4125
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
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line lex.go:4148
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
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line lex.go:4171
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
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:4194
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
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line lex.go:4217
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr58
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr58
			}
		default:
			goto tr58
		}
		goto st0
tr58:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st28
tr88:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line lex.go:4244
		switch data[p] {
		case 39:
			goto tr60
		case 92:
			goto tr61
		}
		goto tr59
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
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
//line lex.go:4281
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr63
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr63
			}
		default:
			goto tr63
		}
		goto st0
tr63:
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line lex.go:4304
		switch data[p] {
		case 39:
			goto tr65
		case 92:
			goto tr66
		}
		goto tr64
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr68
		case 92:
			goto tr69
		}
		goto tr67
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 39:
			goto tr71
		case 92:
			goto tr72
		}
		goto tr70
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 39:
			goto tr74
		case 92:
			goto tr75
		}
		goto tr73
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 39:
			goto tr77
		case 92:
			goto tr78
		}
		goto tr76
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 39:
			goto tr80
		case 92:
			goto tr81
		}
		goto tr79
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 39:
			goto tr83
		case 92:
			goto tr84
		}
		goto tr82
tr44:
//line lex.rl:208
 rn = 0 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line lex.go:4393
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
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:4416
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
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lex.go:4439
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
tr87:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:4462
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr88
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr88
			}
		default:
			goto tr88
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 39:
			goto tr90
		case 92:
			goto tr91
		}
		goto tr89
tr24:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st43
tr108:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st43
tr119:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st43
tr124:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st43
tr127:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st43
tr130:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st43
tr133:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st43
tr136:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st43
tr139:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st43
tr142:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st43
tr149:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lex.go:4537
		switch data[p] {
		case 85:
			goto tr94
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
			goto tr102
		case 118:
			goto st69
		case 120:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr93
		}
		goto tr92
tr93:
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
//line lex.go:4577
		if 48 <= data[p] && data[p] <= 55 {
			goto tr104
		}
		goto st0
tr104:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:4591
		if 48 <= data[p] && data[p] <= 55 {
			goto tr105
		}
		goto st0
tr105:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lex.go:4605
		switch data[p] {
		case 39:
			goto tr107
		case 92:
			goto tr108
		}
		goto tr106
tr94:
//line lex.rl:208
 rn = 0 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:4622
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
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lex.go:4645
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
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line lex.go:4668
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
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:4691
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
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line lex.go:4714
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
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line lex.go:4737
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
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lex.go:4760
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
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line lex.go:4783
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr116
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr116
			}
		default:
			goto tr116
		}
		goto st0
tr116:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st55
tr146:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line lex.go:4810
		switch data[p] {
		case 39:
			goto tr118
		case 92:
			goto tr119
		}
		goto tr117
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
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
//line lex.go:4847
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr121
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr121
			}
		default:
			goto tr121
		}
		goto st0
tr121:
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line lex.go:4870
		switch data[p] {
		case 39:
			goto tr123
		case 92:
			goto tr124
		}
		goto tr122
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 39:
			goto tr126
		case 92:
			goto tr127
		}
		goto tr125
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 39:
			goto tr129
		case 92:
			goto tr130
		}
		goto tr128
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 39:
			goto tr132
		case 92:
			goto tr133
		}
		goto tr131
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 39:
			goto tr135
		case 92:
			goto tr136
		}
		goto tr134
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 39:
			goto tr138
		case 92:
			goto tr139
		}
		goto tr137
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 39:
			goto tr141
		case 92:
			goto tr142
		}
		goto tr140
tr102:
//line lex.rl:208
 rn = 0 
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lex.go:4959
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
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lex.go:4982
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
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lex.go:5005
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
tr145:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lex.go:5028
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr146
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr146
			}
		default:
			goto tr146
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 39:
			goto tr148
		case 92:
			goto tr149
		}
		goto tr147
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st89
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st89
			}
		default:
			goto st89
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto tr276
		case 34:
			goto tr277
		case 36:
			goto tr278
		case 39:
			goto tr279
		case 43:
			goto tr280
		case 45:
			goto tr281
		case 46:
			goto tr282
		case 88:
			goto tr284
		case 95:
			goto tr285
		case 120:
			goto tr286
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 14:
				switch {
				case data[p] > 8:
					if 9 <= data[p] && data[p] <= 13 {
						goto tr276
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
					goto st89
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
					goto tr283
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 127 {
						goto st0
					}
				case data[p] >= 103:
					goto tr285
				}
			default:
				goto st89
			}
		default:
			goto st89
		}
		goto tr287
tr21:
//line lex.rl:250
 ch = 0 
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st71
tr153:
//line lex.rl:252
 buf.WriteByte(ch) 
//line lex.rl:250
 ch = 0 
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lex.go:5169
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr151
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr151
			}
		default:
			goto tr151
		}
		goto st0
tr151:
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lex.go:5192
		if data[p] == 39 {
			goto tr152
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr153
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr153
			}
		default:
			goto tr153
		}
		goto st0
	st_out:
	_test_eof73: cs = 73; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
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
	_test_eof70: cs = 70; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 75:
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
		
		case 77, 79, 81, 84:
//line lex.rl:56

			str()
			if isFconst {
				emit(Fconst, s)
			} else {
				emit(Iconst, s)
			}
			isFconst = false
		
		case 89:
//line lex.rl:73

			str()
			emit(Iconst, s)
		
		case 80, 82, 85, 87:
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
		
		case 74:
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
		
		case 76:
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
		
		case 83, 86:
//line lex.rl:235

			emit(Bconst, buf.String())
		
		case 88:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(Sconst, buf.String())
		
//line lex.go:5423
		}
	}

	_out: {}
	}

//line lex.rl:276


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
