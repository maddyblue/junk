
//line lex.rl:1
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
	
	"github.com/cockroachdb/cockroach/pkg/sql/lex"
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
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 1:
		goto st_case_1
	case 73:
		goto st_case_73
	case 2:
		goto st_case_2
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 3:
		goto st_case_3
	case 80:
		goto st_case_80
	case 4:
		goto st_case_4
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 5:
		goto st_case_5
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 6:
		goto st_case_6
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 113:
		goto st_case_113
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
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
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 122:
		goto st_case_122
	case 69:
		goto st_case_69
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	}
	goto st_out
tr179:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st70
tr209:
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
			emit(lex.IDENT, string(b))
		
	goto st70
tr243:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st70
tr272:
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
			emit(lex.SCONST, string(b))
		
	goto st70
tr305:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st70
tr333:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st70
tr363:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st70
tr395:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st70
tr424:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st70
tr457:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st70
tr487:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st70
tr518:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st70
tr550:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st70
tr580:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st70
tr609:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st70
tr632:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st70
tr661:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st70
tr680:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st70
tr711:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st70
tr739:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st70
tr768:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st70
tr797:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st70
tr826:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st70
tr856:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st70
tr885:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st70
tr914:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st70
tr943:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st70
tr973:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st70
tr998:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st70
tr1027:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st70
tr1056:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st70
tr1083:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st70
tr1112:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st70
tr1141:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st70
tr1170:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st70
tr1199:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st70
tr1229:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line lex.go:603
		switch data[p] {
		case 32:
			goto st70
		case 33:
			goto st71
		case 34:
			goto tr152
		case 35:
			goto st74
		case 36:
			goto tr154
		case 38:
			goto st77
		case 39:
			goto tr157
		case 43:
			goto tr158
		case 45:
			goto tr159
		case 46:
			goto tr160
		case 47:
			goto st87
		case 48:
			goto tr162
		case 58:
			goto st91
		case 60:
			goto st94
		case 62:
			goto st97
		case 63:
			goto st99
		case 64:
			goto st101
		case 88:
			goto tr170
		case 95:
			goto tr171
		case 98:
			goto tr173
		case 101:
			goto tr174
		case 120:
			goto tr175
		case 124:
			goto st104
		case 126:
			goto st108
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto st76
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st70
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr169
					}
				case data[p] >= 59:
					goto st76
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto st76
					}
				case data[p] >= 97:
					goto tr172
				}
			default:
				goto st76
			}
		default:
			goto tr163
		}
		goto tr178
st_case_0:
	st0:
		cs = 0
		goto _out
tr180:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st71
tr210:
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
			emit(lex.IDENT, string(b))
		
	goto st71
tr244:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st71
tr273:
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
			emit(lex.SCONST, string(b))
		
	goto st71
tr306:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st71
tr334:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st71
tr364:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st71
tr396:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st71
tr425:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st71
tr458:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st71
tr488:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st71
tr519:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st71
tr551:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st71
tr581:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st71
tr610:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st71
tr633:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st71
tr662:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st71
tr681:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st71
tr712:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st71
tr740:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st71
tr769:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st71
tr798:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st71
tr827:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st71
tr857:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st71
tr886:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st71
tr915:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st71
tr944:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st71
tr974:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st71
tr999:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st71
tr1028:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st71
tr1057:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st71
tr1084:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st71
tr1113:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st71
tr1142:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st71
tr1171:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st71
tr1200:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st71
tr1230:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lex.go:976
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 59:
			goto tr184
		case 60:
			goto tr194
		case 61:
			goto st130
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto st131
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr199
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
tr152:
//line lex.rl:41
 mark = p 
	goto st72
tr181:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st72
tr245:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st72
tr274:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st72
tr307:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st72
tr335:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st72
tr365:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st72
tr397:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st72
tr426:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st72
tr459:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st72
tr489:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr520:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr552:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st72
tr582:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st72
tr611:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:41
 mark = p 
	goto st72
tr634:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st72
tr663:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st72
tr682:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st72
tr713:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st72
tr741:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st72
tr770:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st72
tr799:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st72
tr828:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st72
tr858:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr887:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr916:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr945:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st72
tr975:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st72
tr1000:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st72
tr1029:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st72
tr1058:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st72
tr1085:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st72
tr1114:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st72
tr1143:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st72
tr1172:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st72
tr1201:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st72
tr1231:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lex.go:1392
		if data[p] == 34 {
			goto st73
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
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
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:1419
		if data[p] == 34 {
			goto st73
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:120
 numQuote++ 
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lex.go:1436
		switch data[p] {
		case 32:
			goto tr209
		case 33:
			goto tr210
		case 34:
			goto st2
		case 35:
			goto tr212
		case 36:
			goto tr213
		case 38:
			goto tr215
		case 39:
			goto tr216
		case 43:
			goto tr217
		case 45:
			goto tr218
		case 46:
			goto tr219
		case 47:
			goto tr220
		case 48:
			goto tr221
		case 58:
			goto tr223
		case 60:
			goto tr224
		case 62:
			goto tr225
		case 63:
			goto tr226
		case 64:
			goto tr227
		case 88:
			goto tr229
		case 95:
			goto tr230
		case 98:
			goto tr232
		case 101:
			goto tr233
		case 120:
			goto tr234
		case 124:
			goto tr235
		case 126:
			goto tr236
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr214
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr209
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr228
					}
				case data[p] >= 59:
					goto tr214
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr214
					}
				case data[p] >= 97:
					goto tr231
				}
			default:
				goto tr214
			}
		default:
			goto tr222
		}
		goto tr237
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
tr182:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st74
tr212:
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
			emit(lex.IDENT, string(b))
		
	goto st74
tr246:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st74
tr275:
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
			emit(lex.SCONST, string(b))
		
	goto st74
tr308:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st74
tr336:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st74
tr366:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st74
tr398:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st74
tr427:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st74
tr460:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st74
tr490:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st74
tr521:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st74
tr553:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st74
tr583:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st74
tr612:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st74
tr635:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st74
tr664:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st74
tr683:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st74
tr714:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st74
tr742:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st74
tr771:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st74
tr800:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st74
tr829:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st74
tr859:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st74
tr888:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st74
tr917:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st74
tr946:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st74
tr976:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st74
tr1001:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st74
tr1030:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st74
tr1059:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st74
tr1086:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st74
tr1115:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st74
tr1144:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st74
tr1173:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st74
tr1202:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st74
tr1232:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lex.go:1817
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto st127
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto st128
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
tr154:
//line lex.rl:41
 mark = p 
	goto st75
tr183:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st75
tr213:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st75
tr247:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st75
tr276:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st75
tr309:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st75
tr337:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st75
tr367:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st75
tr399:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st75
tr428:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st75
tr461:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st75
tr491:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr522:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr554:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st75
tr584:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st75
tr636:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st75
tr684:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st75
tr715:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st75
tr743:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st75
tr772:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st75
tr801:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st75
tr830:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st75
tr860:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr889:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr918:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr947:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st75
tr977:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st75
tr1002:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st75
tr1031:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st75
tr1060:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st75
tr1087:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st75
tr1116:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st75
tr1145:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st75
tr1174:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st75
tr1203:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st75
tr1233:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lex.go:2212
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto st126
		}
		goto tr208
tr184:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st76
tr214:
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
			emit(lex.IDENT, string(b))
		
	goto st76
tr248:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st76
tr277:
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
			emit(lex.SCONST, string(b))
		
	goto st76
tr310:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st76
tr338:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st76
tr368:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st76
tr400:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st76
tr429:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st76
tr462:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st76
tr492:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st76
tr523:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st76
tr555:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st76
tr585:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st76
tr614:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st76
tr637:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st76
tr665:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st76
tr685:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st76
tr716:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st76
tr744:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st76
tr773:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st76
tr802:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st76
tr831:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st76
tr861:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st76
tr890:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st76
tr919:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st76
tr948:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st76
tr978:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st76
tr1003:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st76
tr1032:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st76
tr1061:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st76
tr1088:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st76
tr1117:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st76
tr1146:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st76
tr1175:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st76
tr1204:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st76
tr1234:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lex.go:2579
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
tr185:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st77
tr215:
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
			emit(lex.IDENT, string(b))
		
	goto st77
tr249:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st77
tr278:
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
			emit(lex.SCONST, string(b))
		
	goto st77
tr311:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st77
tr339:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st77
tr369:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st77
tr401:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st77
tr430:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st77
tr463:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st77
tr493:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st77
tr524:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st77
tr556:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st77
tr586:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st77
tr615:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st77
tr638:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st77
tr666:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st77
tr686:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st77
tr717:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st77
tr745:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st77
tr774:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st77
tr803:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st77
tr832:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st77
tr862:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st77
tr891:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st77
tr920:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st77
tr949:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st77
tr979:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st77
tr1004:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st77
tr1033:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st77
tr1062:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st77
tr1089:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st77
tr1118:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st77
tr1147:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st77
tr1176:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st77
tr1205:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st77
tr1235:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lex.go:2948
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto st78
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 32:
			goto tr243
		case 33:
			goto tr244
		case 34:
			goto tr245
		case 35:
			goto tr246
		case 36:
			goto tr247
		case 38:
			goto tr249
		case 39:
			goto tr250
		case 43:
			goto tr251
		case 45:
			goto tr252
		case 46:
			goto tr253
		case 47:
			goto tr254
		case 48:
			goto tr255
		case 58:
			goto tr257
		case 60:
			goto tr258
		case 62:
			goto tr259
		case 63:
			goto tr260
		case 64:
			goto tr261
		case 88:
			goto tr263
		case 95:
			goto tr264
		case 98:
			goto tr266
		case 101:
			goto tr267
		case 120:
			goto tr268
		case 124:
			goto tr269
		case 126:
			goto tr270
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr248
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr243
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr262
					}
				case data[p] >= 59:
					goto tr248
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr248
					}
				case data[p] >= 97:
					goto tr265
				}
			default:
				goto tr248
			}
		default:
			goto tr256
		}
		goto tr271
tr157:
//line lex.rl:41
 mark = p 
	goto st79
tr186:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st79
tr216:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st79
tr250:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st79
tr312:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st79
tr340:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st79
tr370:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st79
tr402:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st79
tr431:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st79
tr464:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st79
tr494:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr525:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr557:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st79
tr587:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st79
tr616:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:41
 mark = p 
	goto st79
tr639:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st79
tr667:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st79
tr687:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st79
tr746:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st79
tr775:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st79
tr804:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st79
tr833:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st79
tr863:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr892:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr921:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr950:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st79
tr980:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st79
tr1005:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st79
tr1034:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st79
tr1063:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st79
tr1090:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st79
tr1119:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st79
tr1148:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st79
tr1177:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st79
tr1206:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st79
tr1236:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lex.go:3457
		if data[p] == 39 {
			goto st80
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr8:
//line lex.rl:106
 isNotASCII = true 
	goto st3
tr9:
//line lex.rl:154
 numQuote++ 
	goto st3
tr11:
//line lex.rl:154
 numQuote++ 
//line lex.rl:106
 isNotASCII = true 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:3484
		if data[p] == 39 {
			goto st80
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr10:
//line lex.rl:154
 numQuote++ 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lex.go:3501
		switch data[p] {
		case 32:
			goto tr272
		case 33:
			goto tr273
		case 34:
			goto tr274
		case 35:
			goto tr275
		case 36:
			goto tr276
		case 38:
			goto tr278
		case 39:
			goto st4
		case 43:
			goto tr280
		case 45:
			goto tr281
		case 46:
			goto tr282
		case 47:
			goto tr283
		case 48:
			goto tr284
		case 58:
			goto tr286
		case 60:
			goto tr287
		case 62:
			goto tr288
		case 63:
			goto tr289
		case 64:
			goto tr290
		case 88:
			goto tr292
		case 95:
			goto tr293
		case 98:
			goto tr295
		case 101:
			goto tr296
		case 120:
			goto tr297
		case 124:
			goto tr298
		case 126:
			goto tr299
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr277
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr272
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr291
					}
				case data[p] >= 59:
					goto tr277
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr277
					}
				case data[p] >= 97:
					goto tr294
				}
			default:
				goto tr277
			}
		default:
			goto tr285
		}
		goto tr300
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 39 {
			goto tr10
		}
		if data[p] <= 127 {
			goto tr9
		}
		goto tr11
tr158:
//line lex.rl:41
 mark = p 
	goto st81
tr187:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st81
tr217:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st81
tr251:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st81
tr280:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st81
tr313:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st81
tr341:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st81
tr371:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st81
tr403:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st81
tr432:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st81
tr465:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st81
tr495:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr526:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr558:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st81
tr588:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st81
tr617:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:41
 mark = p 
	goto st81
tr640:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st81
tr668:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st81
tr688:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st81
tr718:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st81
tr747:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st81
tr776:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st81
tr805:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st81
tr834:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st81
tr864:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr893:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr922:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr951:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st81
tr981:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st81
tr1006:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st81
tr1035:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st81
tr1064:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st81
tr1091:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st81
tr1120:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st81
tr1149:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st81
tr1178:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st81
tr1208:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st81
tr1237:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lex.go:3960
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr301
		case 47:
			goto tr190
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr302
		}
		goto tr208
tr159:
//line lex.rl:41
 mark = p 
	goto st82
tr188:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st82
tr218:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st82
tr252:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st82
tr281:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st82
tr314:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st82
tr342:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st82
tr372:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st82
tr404:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st82
tr433:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st82
tr466:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st82
tr496:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr527:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr559:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st82
tr589:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st82
tr618:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:41
 mark = p 
	goto st82
tr641:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st82
tr669:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st82
tr689:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st82
tr719:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st82
tr748:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st82
tr777:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st82
tr806:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st82
tr835:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st82
tr865:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr894:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr923:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr952:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st82
tr982:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st82
tr1007:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st82
tr1036:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st82
tr1065:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st82
tr1092:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st82
tr1121:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st82
tr1150:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st82
tr1179:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st82
tr1209:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st82
tr1238:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lex.go:4405
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto st83
		case 46:
			goto tr301
		case 47:
			goto tr190
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto st124
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr302
		}
		goto tr208
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 10 {
			goto st70
		}
		goto st83
tr301:
//line lex.rl:70

			isFconst = true
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line lex.go:4522
		if 48 <= data[p] && data[p] <= 57 {
			goto st84
		}
		goto st0
tr392:
//line lex.rl:70

			isFconst = true
		
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line lex.go:4538
		switch data[p] {
		case 32:
			goto tr305
		case 33:
			goto tr306
		case 34:
			goto tr307
		case 35:
			goto tr308
		case 36:
			goto tr309
		case 38:
			goto tr311
		case 39:
			goto tr312
		case 43:
			goto tr313
		case 45:
			goto tr314
		case 46:
			goto tr315
		case 47:
			goto tr316
		case 58:
			goto tr317
		case 60:
			goto tr318
		case 62:
			goto tr319
		case 63:
			goto tr320
		case 64:
			goto tr321
		case 69:
			goto tr323
		case 88:
			goto tr324
		case 95:
			goto tr325
		case 98:
			goto tr327
		case 101:
			goto tr323
		case 120:
			goto tr328
		case 124:
			goto tr329
		case 126:
			goto tr330
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr310
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr305
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr322
					}
				case data[p] >= 59:
					goto tr310
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr310
					}
				case data[p] >= 97:
					goto tr326
				}
			default:
				goto tr310
			}
		default:
			goto st84
		}
		goto tr331
tr160:
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr189:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr219:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr253:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr282:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr315:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr343:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr373:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr405:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr434:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr467:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr497:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr528:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr560:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr590:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr619:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr642:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr670:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr690:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr720:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr749:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr778:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr807:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr836:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr866:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr895:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr924:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr953:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr983:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1008:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1037:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1066:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1093:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1122:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1151:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1180:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1210:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
tr1239:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:70

			isFconst = true
		
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lex.go:5137
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto st86
		case 47:
			goto tr190
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto st84
		}
		goto tr208
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 32:
			goto tr333
		case 33:
			goto tr334
		case 34:
			goto tr335
		case 35:
			goto tr336
		case 36:
			goto tr337
		case 38:
			goto tr339
		case 39:
			goto tr340
		case 43:
			goto tr341
		case 45:
			goto tr342
		case 46:
			goto tr343
		case 47:
			goto tr344
		case 48:
			goto tr345
		case 58:
			goto tr347
		case 60:
			goto tr348
		case 62:
			goto tr349
		case 63:
			goto tr350
		case 64:
			goto tr351
		case 88:
			goto tr353
		case 95:
			goto tr354
		case 98:
			goto tr356
		case 101:
			goto tr357
		case 120:
			goto tr358
		case 124:
			goto tr359
		case 126:
			goto tr360
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr338
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr333
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr352
					}
				case data[p] >= 59:
					goto tr338
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr338
					}
				case data[p] >= 97:
					goto tr355
				}
			default:
				goto tr338
			}
		default:
			goto tr346
		}
		goto tr361
tr190:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st87
tr220:
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
			emit(lex.IDENT, string(b))
		
	goto st87
tr254:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st87
tr283:
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
			emit(lex.SCONST, string(b))
		
	goto st87
tr316:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st87
tr344:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st87
tr374:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st87
tr406:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st87
tr435:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st87
tr468:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st87
tr498:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st87
tr529:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st87
tr561:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st87
tr591:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st87
tr620:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st87
tr643:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st87
tr671:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st87
tr691:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st87
tr721:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st87
tr750:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st87
tr779:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st87
tr808:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st87
tr837:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st87
tr867:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st87
tr896:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st87
tr925:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st87
tr954:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st87
tr984:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st87
tr1009:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st87
tr1038:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st87
tr1067:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st87
tr1094:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st87
tr1123:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st87
tr1152:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st87
tr1181:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st87
tr1211:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st87
tr1240:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line lex.go:5607
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto st88
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto tr363
		case 33:
			goto tr364
		case 34:
			goto tr365
		case 35:
			goto tr366
		case 36:
			goto tr367
		case 38:
			goto tr369
		case 39:
			goto tr370
		case 43:
			goto tr371
		case 45:
			goto tr372
		case 46:
			goto tr373
		case 47:
			goto tr374
		case 48:
			goto tr375
		case 58:
			goto tr377
		case 60:
			goto tr378
		case 62:
			goto tr379
		case 63:
			goto tr380
		case 64:
			goto tr381
		case 88:
			goto tr383
		case 95:
			goto tr384
		case 98:
			goto tr386
		case 101:
			goto tr387
		case 120:
			goto tr388
		case 124:
			goto tr389
		case 126:
			goto tr390
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr368
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr363
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr382
					}
				case data[p] >= 59:
					goto tr368
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr368
					}
				case data[p] >= 97:
					goto tr385
				}
			default:
				goto tr368
			}
		default:
			goto tr376
		}
		goto tr391
tr162:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr191:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr221:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr255:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr284:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr345:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr375:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr407:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr436:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr469:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr499:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr530:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr562:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr592:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr644:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr692:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr722:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr751:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr780:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr809:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr838:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr868:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr897:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr926:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr955:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1010:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1039:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1095:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1124:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1153:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1182:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1212:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
tr1241:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line lex.go:6263
		switch data[p] {
		case 32:
			goto tr305
		case 33:
			goto tr306
		case 34:
			goto tr307
		case 35:
			goto tr308
		case 36:
			goto tr309
		case 38:
			goto tr311
		case 39:
			goto tr312
		case 43:
			goto tr313
		case 45:
			goto tr314
		case 46:
			goto tr392
		case 47:
			goto tr316
		case 58:
			goto tr317
		case 60:
			goto tr318
		case 62:
			goto tr319
		case 63:
			goto tr320
		case 64:
			goto tr321
		case 69:
			goto tr323
		case 88:
			goto tr324
		case 95:
			goto tr325
		case 98:
			goto tr327
		case 101:
			goto tr323
		case 120:
			goto st69
		case 124:
			goto tr329
		case 126:
			goto tr330
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr310
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr305
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr322
					}
				case data[p] >= 59:
					goto tr310
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr310
					}
				case data[p] >= 97:
					goto tr326
				}
			default:
				goto tr310
			}
		default:
			goto tr302
		}
		goto tr331
tr163:
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr192:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr222:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr256:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr285:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr302:
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr346:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr376:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr408:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr437:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr470:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr500:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr531:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr563:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr593:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr645:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr693:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr723:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr752:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr781:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr810:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr839:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr869:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr898:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr927:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr956:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1011:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1040:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1096:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1125:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1154:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1183:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1213:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1242:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:65

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line lex.go:6824
		switch data[p] {
		case 32:
			goto tr305
		case 33:
			goto tr306
		case 34:
			goto tr307
		case 35:
			goto tr308
		case 36:
			goto tr309
		case 38:
			goto tr311
		case 39:
			goto tr312
		case 43:
			goto tr313
		case 45:
			goto tr314
		case 46:
			goto tr392
		case 47:
			goto tr316
		case 58:
			goto tr317
		case 60:
			goto tr318
		case 62:
			goto tr319
		case 63:
			goto tr320
		case 64:
			goto tr321
		case 69:
			goto tr323
		case 88:
			goto tr324
		case 95:
			goto tr325
		case 98:
			goto tr327
		case 101:
			goto tr323
		case 120:
			goto tr328
		case 124:
			goto tr329
		case 126:
			goto tr330
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr310
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr305
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr322
					}
				case data[p] >= 59:
					goto tr310
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr310
					}
				case data[p] >= 97:
					goto tr326
				}
			default:
				goto tr310
			}
		default:
			goto tr302
		}
		goto tr331
tr193:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st91
tr223:
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
			emit(lex.IDENT, string(b))
		
	goto st91
tr257:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st91
tr286:
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
			emit(lex.SCONST, string(b))
		
	goto st91
tr317:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st91
tr347:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st91
tr377:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st91
tr438:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st91
tr471:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st91
tr501:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st91
tr532:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st91
tr564:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st91
tr594:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st91
tr621:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st91
tr646:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st91
tr672:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st91
tr694:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st91
tr724:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st91
tr753:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st91
tr782:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st91
tr811:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st91
tr840:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st91
tr870:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st91
tr899:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st91
tr928:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st91
tr957:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st91
tr985:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st91
tr1012:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st91
tr1041:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st91
tr1068:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st91
tr1097:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st91
tr1126:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st91
tr1155:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st91
tr1184:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st91
tr1214:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st91
tr1243:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line lex.go:7189
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto st92
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto tr395
		case 33:
			goto tr396
		case 34:
			goto tr397
		case 35:
			goto tr398
		case 36:
			goto tr399
		case 38:
			goto tr401
		case 39:
			goto tr402
		case 43:
			goto tr403
		case 45:
			goto tr404
		case 46:
			goto tr405
		case 47:
			goto tr406
		case 48:
			goto tr407
		case 58:
			goto st93
		case 60:
			goto tr410
		case 62:
			goto tr411
		case 63:
			goto tr412
		case 64:
			goto tr413
		case 88:
			goto tr415
		case 95:
			goto tr416
		case 98:
			goto tr418
		case 101:
			goto tr419
		case 120:
			goto tr420
		case 124:
			goto tr421
		case 126:
			goto tr422
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr400
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr395
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr414
					}
				case data[p] >= 59:
					goto tr400
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr400
					}
				case data[p] >= 97:
					goto tr417
				}
			default:
				goto tr400
			}
		default:
			goto tr408
		}
		goto tr423
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto tr424
		case 33:
			goto tr425
		case 34:
			goto tr426
		case 35:
			goto tr427
		case 36:
			goto tr428
		case 38:
			goto tr430
		case 39:
			goto tr431
		case 43:
			goto tr432
		case 45:
			goto tr433
		case 46:
			goto tr434
		case 47:
			goto tr435
		case 48:
			goto tr436
		case 58:
			goto tr438
		case 60:
			goto tr439
		case 62:
			goto tr440
		case 63:
			goto tr441
		case 64:
			goto tr442
		case 88:
			goto tr444
		case 95:
			goto tr445
		case 98:
			goto tr447
		case 101:
			goto tr448
		case 120:
			goto tr449
		case 124:
			goto tr450
		case 126:
			goto tr451
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr429
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr424
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr443
					}
				case data[p] >= 59:
					goto tr429
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr429
					}
				case data[p] >= 97:
					goto tr446
				}
			default:
				goto tr429
			}
		default:
			goto tr437
		}
		goto tr452
tr194:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st94
tr224:
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
			emit(lex.IDENT, string(b))
		
	goto st94
tr258:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st94
tr287:
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
			emit(lex.SCONST, string(b))
		
	goto st94
tr318:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st94
tr348:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st94
tr378:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st94
tr410:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st94
tr439:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st94
tr472:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st94
tr502:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st94
tr533:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st94
tr565:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st94
tr595:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st94
tr622:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st94
tr647:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st94
tr673:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st94
tr695:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st94
tr725:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st94
tr754:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st94
tr783:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st94
tr812:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st94
tr841:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st94
tr871:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st94
tr900:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st94
tr929:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st94
tr958:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st94
tr986:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st94
tr1013:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st94
tr1042:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st94
tr1069:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st94
tr1098:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st94
tr1127:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st94
tr1156:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st94
tr1185:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st94
tr1215:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st94
tr1244:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line lex.go:7764
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 59:
			goto tr184
		case 60:
			goto st95
		case 61:
			goto st119
		case 62:
			goto st120
		case 63:
			goto tr197
		case 64:
			goto st121
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr199
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto tr457
		case 33:
			goto tr458
		case 34:
			goto tr459
		case 35:
			goto tr460
		case 36:
			goto tr461
		case 38:
			goto tr463
		case 39:
			goto tr464
		case 43:
			goto tr465
		case 45:
			goto tr466
		case 46:
			goto tr467
		case 47:
			goto tr468
		case 48:
			goto tr469
		case 58:
			goto tr471
		case 59:
			goto tr462
		case 60:
			goto tr472
		case 61:
			goto st96
		case 62:
			goto tr474
		case 63:
			goto tr475
		case 64:
			goto tr476
		case 88:
			goto tr478
		case 95:
			goto tr479
		case 98:
			goto tr481
		case 101:
			goto tr482
		case 120:
			goto tr483
		case 124:
			goto tr484
		case 126:
			goto tr485
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr462
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr457
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr477
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr462
					}
				case data[p] >= 97:
					goto tr480
				}
			default:
				goto tr462
			}
		default:
			goto tr470
		}
		goto tr486
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto tr487
		case 33:
			goto tr488
		case 34:
			goto tr489
		case 35:
			goto tr490
		case 36:
			goto tr491
		case 38:
			goto tr493
		case 39:
			goto tr494
		case 43:
			goto tr495
		case 45:
			goto tr496
		case 46:
			goto tr497
		case 47:
			goto tr498
		case 48:
			goto tr499
		case 58:
			goto tr501
		case 60:
			goto tr502
		case 62:
			goto tr503
		case 63:
			goto tr504
		case 64:
			goto tr505
		case 88:
			goto tr507
		case 95:
			goto tr508
		case 98:
			goto tr510
		case 101:
			goto tr511
		case 120:
			goto tr512
		case 124:
			goto tr513
		case 126:
			goto tr514
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr492
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr487
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr506
					}
				case data[p] >= 59:
					goto tr492
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr492
					}
				case data[p] >= 97:
					goto tr509
				}
			default:
				goto tr492
			}
		default:
			goto tr500
		}
		goto tr515
tr196:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st97
tr225:
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
			emit(lex.IDENT, string(b))
		
	goto st97
tr259:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st97
tr288:
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
			emit(lex.SCONST, string(b))
		
	goto st97
tr319:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st97
tr349:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st97
tr379:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st97
tr411:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st97
tr440:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st97
tr474:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st97
tr503:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st97
tr534:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st97
tr566:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st97
tr596:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st97
tr623:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st97
tr648:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st97
tr674:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st97
tr696:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st97
tr726:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st97
tr755:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st97
tr784:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st97
tr813:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st97
tr843:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st97
tr872:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st97
tr901:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st97
tr930:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st97
tr959:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st97
tr987:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st97
tr1043:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st97
tr1070:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st97
tr1099:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st97
tr1157:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st97
tr1186:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st97
tr1216:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st97
tr1245:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line lex.go:8329
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 59:
			goto tr184
		case 60:
			goto tr194
		case 61:
			goto st98
		case 62:
			goto st117
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr199
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 32:
			goto tr518
		case 33:
			goto tr519
		case 34:
			goto tr520
		case 35:
			goto tr521
		case 36:
			goto tr522
		case 38:
			goto tr524
		case 39:
			goto tr525
		case 43:
			goto tr526
		case 45:
			goto tr527
		case 46:
			goto tr528
		case 47:
			goto tr529
		case 48:
			goto tr530
		case 58:
			goto tr532
		case 60:
			goto tr533
		case 62:
			goto tr534
		case 63:
			goto tr535
		case 64:
			goto tr536
		case 88:
			goto tr538
		case 95:
			goto tr539
		case 98:
			goto tr541
		case 101:
			goto tr542
		case 120:
			goto tr543
		case 124:
			goto tr544
		case 126:
			goto tr545
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr523
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr518
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr537
					}
				case data[p] >= 59:
					goto tr523
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr523
					}
				case data[p] >= 97:
					goto tr540
				}
			default:
				goto tr523
			}
		default:
			goto tr531
		}
		goto tr546
tr197:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st99
tr226:
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
			emit(lex.IDENT, string(b))
		
	goto st99
tr260:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st99
tr289:
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
			emit(lex.SCONST, string(b))
		
	goto st99
tr320:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st99
tr350:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st99
tr380:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st99
tr412:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st99
tr441:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st99
tr475:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st99
tr504:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st99
tr535:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st99
tr567:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st99
tr597:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st99
tr624:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st99
tr649:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st99
tr675:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st99
tr697:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st99
tr727:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st99
tr756:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st99
tr785:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st99
tr814:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st99
tr844:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st99
tr873:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st99
tr902:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st99
tr931:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st99
tr960:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st99
tr988:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st99
tr1015:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st99
tr1044:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st99
tr1071:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st99
tr1100:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st99
tr1129:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st99
tr1158:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st99
tr1187:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st99
tr1217:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st99
tr1246:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line lex.go:8800
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto st100
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto st115
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto st116
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto tr550
		case 33:
			goto tr551
		case 34:
			goto tr552
		case 35:
			goto tr553
		case 36:
			goto tr554
		case 38:
			goto tr556
		case 39:
			goto tr557
		case 43:
			goto tr558
		case 45:
			goto tr559
		case 46:
			goto tr560
		case 47:
			goto tr561
		case 48:
			goto tr562
		case 58:
			goto tr564
		case 60:
			goto tr565
		case 62:
			goto tr566
		case 63:
			goto tr567
		case 64:
			goto tr568
		case 88:
			goto tr570
		case 95:
			goto tr571
		case 98:
			goto tr573
		case 101:
			goto tr574
		case 120:
			goto tr575
		case 124:
			goto tr576
		case 126:
			goto tr577
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr555
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr550
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr569
					}
				case data[p] >= 59:
					goto tr555
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr555
					}
				case data[p] >= 97:
					goto tr572
				}
			default:
				goto tr555
			}
		default:
			goto tr563
		}
		goto tr578
tr198:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st101
tr227:
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
			emit(lex.IDENT, string(b))
		
	goto st101
tr261:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st101
tr290:
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
			emit(lex.SCONST, string(b))
		
	goto st101
tr321:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st101
tr351:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st101
tr381:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st101
tr413:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st101
tr442:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st101
tr476:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st101
tr505:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st101
tr536:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st101
tr568:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st101
tr598:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st101
tr625:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st101
tr650:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st101
tr676:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st101
tr698:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st101
tr728:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st101
tr757:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st101
tr786:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st101
tr815:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st101
tr845:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st101
tr874:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st101
tr903:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st101
tr932:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st101
tr961:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st101
tr989:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st101
tr1016:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st101
tr1045:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st101
tr1072:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st101
tr1101:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st101
tr1130:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st101
tr1159:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st101
tr1188:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st101
tr1218:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st101
tr1247:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line lex.go:9272
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto st102
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 32:
			goto tr580
		case 33:
			goto tr581
		case 34:
			goto tr582
		case 35:
			goto tr583
		case 36:
			goto tr584
		case 38:
			goto tr586
		case 39:
			goto tr587
		case 43:
			goto tr588
		case 45:
			goto tr589
		case 46:
			goto tr590
		case 47:
			goto tr591
		case 48:
			goto tr592
		case 58:
			goto tr594
		case 60:
			goto tr595
		case 62:
			goto tr596
		case 63:
			goto tr597
		case 64:
			goto tr598
		case 88:
			goto tr600
		case 95:
			goto tr601
		case 98:
			goto tr603
		case 101:
			goto tr604
		case 120:
			goto tr605
		case 124:
			goto tr606
		case 126:
			goto tr607
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr585
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr580
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr599
					}
				case data[p] >= 59:
					goto tr585
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr585
					}
				case data[p] >= 97:
					goto tr602
				}
			default:
				goto tr585
			}
		default:
			goto tr593
		}
		goto tr608
tr629:
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr172:
//line lex.rl:41
 mark = p 
	goto st103
tr169:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr178:
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr202:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st103
tr199:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr208:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr231:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st103
tr228:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr237:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr265:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st103
tr262:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr271:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr294:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st103
tr291:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr300:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr326:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st103
tr322:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr331:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr355:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st103
tr352:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr361:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr385:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st103
tr382:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr391:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr417:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st103
tr414:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr423:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr446:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st103
tr443:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr452:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr480:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st103
tr477:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr486:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr509:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr506:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr515:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr540:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr537:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr546:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr572:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st103
tr569:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr578:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr602:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st103
tr599:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr608:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr626:
//line lex.rl:109
 isUpper = true 
	goto st103
tr654:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st103
tr651:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr660:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr702:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st103
tr699:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr708:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr732:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st103
tr729:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr738:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr761:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st103
tr758:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr767:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr790:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st103
tr787:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr796:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr819:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st103
tr816:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr825:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr849:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st103
tr846:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr855:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr878:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr875:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr884:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr907:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr904:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr913:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr936:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr933:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr942:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr965:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st103
tr962:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr971:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr993:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st103
tr990:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr997:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1020:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st103
tr1017:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1026:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1049:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st103
tr1046:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1055:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1076:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st103
tr1073:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1082:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1105:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st103
tr1102:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1111:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1134:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st103
tr1131:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1140:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1163:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st103
tr1160:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1169:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1192:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st103
tr1189:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1198:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1222:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st103
tr1219:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1228:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
tr1251:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st103
tr1248:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st103
tr1257:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:106
 isNotASCII = true 
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line lex.go:10504
		switch data[p] {
		case 32:
			goto tr609
		case 33:
			goto tr610
		case 34:
			goto tr611
		case 35:
			goto tr612
		case 36:
			goto st103
		case 38:
			goto tr615
		case 39:
			goto tr616
		case 43:
			goto tr617
		case 45:
			goto tr618
		case 46:
			goto tr619
		case 47:
			goto tr620
		case 58:
			goto tr621
		case 60:
			goto tr622
		case 62:
			goto tr623
		case 63:
			goto tr624
		case 64:
			goto tr625
		case 95:
			goto st103
		case 124:
			goto tr627
		case 126:
			goto tr628
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr614
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr609
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr626
					}
				case data[p] >= 59:
					goto tr614
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr614
					}
				case data[p] >= 97:
					goto st103
				}
			default:
				goto tr614
			}
		default:
			goto st103
		}
		goto tr629
tr206:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st104
tr235:
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
			emit(lex.IDENT, string(b))
		
	goto st104
tr269:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st104
tr298:
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
			emit(lex.SCONST, string(b))
		
	goto st104
tr329:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st104
tr359:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st104
tr389:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st104
tr421:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st104
tr450:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st104
tr484:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st104
tr513:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st104
tr544:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st104
tr576:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st104
tr606:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st104
tr627:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st104
tr658:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st104
tr677:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st104
tr706:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st104
tr736:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st104
tr765:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st104
tr794:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st104
tr823:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st104
tr853:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st104
tr882:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st104
tr911:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st104
tr940:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st104
tr969:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st104
tr995:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st104
tr1024:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st104
tr1053:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st104
tr1080:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st104
tr1109:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st104
tr1138:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st104
tr1167:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st104
tr1196:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st104
tr1226:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st104
tr1255:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line lex.go:10863
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto st114
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
tr175:
//line lex.rl:41
 mark = p 
	goto st105
tr170:
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr205:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st105
tr200:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr234:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st105
tr229:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr268:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st105
tr263:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr297:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st105
tr292:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr328:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st105
tr324:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr358:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st105
tr353:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr388:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st105
tr383:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr420:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st105
tr415:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr449:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st105
tr444:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr483:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st105
tr478:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr512:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr507:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr543:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr538:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr575:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st105
tr570:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr605:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st105
tr600:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr657:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st105
tr652:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr705:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st105
tr700:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr735:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st105
tr730:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr764:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st105
tr759:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr793:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st105
tr788:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr822:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st105
tr817:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr852:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st105
tr847:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr881:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr876:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr910:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr905:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr939:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr934:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr968:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st105
tr963:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr994:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st105
tr991:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1023:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st105
tr1018:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1052:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st105
tr1047:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1079:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st105
tr1074:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1108:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st105
tr1103:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1137:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st105
tr1132:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1166:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st105
tr1161:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1195:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st105
tr1190:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1225:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st105
tr1220:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
tr1254:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st105
tr1249:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
//line lex.rl:109
 isUpper = true 
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line lex.go:11621
		switch data[p] {
		case 32:
			goto tr609
		case 33:
			goto tr610
		case 34:
			goto tr611
		case 35:
			goto tr612
		case 36:
			goto st103
		case 38:
			goto tr615
		case 39:
			goto tr631
		case 43:
			goto tr617
		case 45:
			goto tr618
		case 46:
			goto tr619
		case 47:
			goto tr620
		case 58:
			goto tr621
		case 60:
			goto tr622
		case 62:
			goto tr623
		case 63:
			goto tr624
		case 64:
			goto tr625
		case 95:
			goto st103
		case 124:
			goto tr627
		case 126:
			goto tr628
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr614
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr609
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr626
					}
				case data[p] >= 59:
					goto tr614
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr614
					}
				case data[p] >= 97:
					goto st103
				}
			default:
				goto tr614
			}
		default:
			goto st103
		}
		goto tr629
tr631:
//line lex.rl:246
 buf = new(bytes.Buffer) 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line lex.go:11719
		if data[p] == 39 {
			goto st106
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr145:
//line lex.rl:252
 buf.WriteByte(ch) 
	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line lex.go:11745
		switch data[p] {
		case 32:
			goto tr632
		case 33:
			goto tr633
		case 34:
			goto tr634
		case 35:
			goto tr635
		case 36:
			goto tr636
		case 38:
			goto tr638
		case 39:
			goto tr639
		case 43:
			goto tr640
		case 45:
			goto tr641
		case 46:
			goto tr642
		case 47:
			goto tr643
		case 48:
			goto tr644
		case 58:
			goto tr646
		case 60:
			goto tr647
		case 62:
			goto tr648
		case 63:
			goto tr649
		case 64:
			goto tr650
		case 88:
			goto tr652
		case 95:
			goto tr653
		case 98:
			goto tr655
		case 101:
			goto tr656
		case 120:
			goto tr657
		case 124:
			goto tr658
		case 126:
			goto tr659
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr637
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr632
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr651
					}
				case data[p] >= 59:
					goto tr637
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr637
					}
				case data[p] >= 97:
					goto tr654
				}
			default:
				goto tr637
			}
		default:
			goto tr645
		}
		goto tr660
tr171:
//line lex.rl:41
 mark = p 
	goto st107
tr201:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st107
tr230:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st107
tr264:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st107
tr293:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st107
tr325:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st107
tr354:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st107
tr384:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st107
tr416:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st107
tr445:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st107
tr479:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st107
tr508:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr539:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr571:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st107
tr601:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st107
tr653:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st107
tr701:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st107
tr731:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st107
tr760:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st107
tr789:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st107
tr818:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st107
tr848:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st107
tr877:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr906:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr935:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr964:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st107
tr992:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:41
 mark = p 
	goto st107
tr1019:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st107
tr1048:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st107
tr1075:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st107
tr1104:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st107
tr1133:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st107
tr1162:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st107
tr1191:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st107
tr1221:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st107
tr1250:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line lex.go:12140
		switch data[p] {
		case 32:
			goto tr661
		case 33:
			goto tr662
		case 34:
			goto tr663
		case 35:
			goto tr664
		case 36:
			goto st103
		case 38:
			goto tr666
		case 39:
			goto tr667
		case 43:
			goto tr668
		case 45:
			goto tr669
		case 46:
			goto tr670
		case 47:
			goto tr671
		case 58:
			goto tr672
		case 60:
			goto tr673
		case 62:
			goto tr674
		case 63:
			goto tr675
		case 64:
			goto tr676
		case 95:
			goto st103
		case 124:
			goto tr677
		case 126:
			goto tr678
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr665
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr661
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr626
					}
				case data[p] >= 59:
					goto tr665
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr665
					}
				case data[p] >= 97:
					goto st103
				}
			default:
				goto tr665
			}
		default:
			goto st103
		}
		goto tr629
tr240:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st108
tr236:
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
			emit(lex.IDENT, string(b))
		
	goto st108
tr270:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st108
tr299:
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
			emit(lex.SCONST, string(b))
		
	goto st108
tr330:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st108
tr360:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
	goto st108
tr390:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
	goto st108
tr422:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
	goto st108
tr451:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
	goto st108
tr485:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
	goto st108
tr514:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st108
tr545:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
	goto st108
tr577:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st108
tr607:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
	goto st108
tr628:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
	goto st108
tr659:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
	goto st108
tr678:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
	goto st108
tr707:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
	goto st108
tr737:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st108
tr766:
//line lex.rl:294
 emitToken(lex.CONCAT) 
	goto st108
tr795:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
	goto st108
tr824:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st108
tr854:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
	goto st108
tr883:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st108
tr912:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
	goto st108
tr941:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
	goto st108
tr970:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
	goto st108
tr996:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
	goto st108
tr1025:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
	goto st108
tr1054:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
	goto st108
tr1081:
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
			emit(lex.PLACEHOLDER, s)
		
	goto st108
tr1110:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
	goto st108
tr1139:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
	goto st108
tr1168:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
	goto st108
tr1197:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
	goto st108
tr1227:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
	goto st108
tr1256:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line lex.go:12499
		switch data[p] {
		case 32:
			goto tr179
		case 33:
			goto tr180
		case 34:
			goto tr181
		case 35:
			goto tr182
		case 36:
			goto tr183
		case 38:
			goto tr185
		case 39:
			goto tr186
		case 42:
			goto st109
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto tr190
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 60:
			goto tr194
		case 62:
			goto tr196
		case 63:
			goto tr197
		case 64:
			goto tr198
		case 88:
			goto tr200
		case 95:
			goto tr201
		case 98:
			goto tr203
		case 101:
			goto tr204
		case 120:
			goto tr205
		case 124:
			goto tr206
		case 126:
			goto tr240
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr184
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr199
					}
				case data[p] >= 59:
					goto tr184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr184
					}
				case data[p] >= 97:
					goto tr202
				}
			default:
				goto tr184
			}
		default:
			goto tr192
		}
		goto tr208
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 32:
			goto tr680
		case 33:
			goto tr681
		case 34:
			goto tr682
		case 35:
			goto tr683
		case 36:
			goto tr684
		case 38:
			goto tr686
		case 39:
			goto tr687
		case 43:
			goto tr688
		case 45:
			goto tr689
		case 46:
			goto tr690
		case 47:
			goto tr691
		case 48:
			goto tr692
		case 58:
			goto tr694
		case 60:
			goto tr695
		case 62:
			goto tr696
		case 63:
			goto tr697
		case 64:
			goto tr698
		case 88:
			goto tr700
		case 95:
			goto tr701
		case 98:
			goto tr703
		case 101:
			goto tr704
		case 120:
			goto tr705
		case 124:
			goto tr706
		case 126:
			goto tr707
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr685
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr680
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr699
					}
				case data[p] >= 59:
					goto tr685
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr685
					}
				case data[p] >= 97:
					goto tr702
				}
			default:
				goto tr685
			}
		default:
			goto tr693
		}
		goto tr708
tr173:
//line lex.rl:41
 mark = p 
	goto st110
tr203:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st110
tr232:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st110
tr266:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st110
tr295:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st110
tr327:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st110
tr356:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st110
tr386:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st110
tr418:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st110
tr447:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st110
tr481:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st110
tr510:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr541:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr573:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st110
tr603:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st110
tr655:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st110
tr703:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st110
tr733:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st110
tr762:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st110
tr791:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st110
tr820:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st110
tr850:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st110
tr879:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr908:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr937:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr966:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st110
tr1021:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st110
tr1050:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st110
tr1077:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st110
tr1106:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st110
tr1135:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st110
tr1164:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st110
tr1193:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st110
tr1223:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st110
tr1252:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line lex.go:12990
		switch data[p] {
		case 32:
			goto tr609
		case 33:
			goto tr610
		case 34:
			goto tr611
		case 35:
			goto tr612
		case 36:
			goto st103
		case 38:
			goto tr615
		case 39:
			goto st7
		case 43:
			goto tr617
		case 45:
			goto tr618
		case 46:
			goto tr619
		case 47:
			goto tr620
		case 58:
			goto tr621
		case 60:
			goto tr622
		case 62:
			goto tr623
		case 63:
			goto tr624
		case 64:
			goto tr625
		case 95:
			goto st103
		case 124:
			goto tr627
		case 126:
			goto tr628
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr614
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr609
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr626
					}
				case data[p] >= 59:
					goto tr614
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr614
					}
				case data[p] >= 97:
					goto st103
				}
			default:
				goto tr614
			}
		default:
			goto st103
		}
		goto tr629
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 39:
			goto tr17
		case 92:
			goto tr18
		}
		goto tr16
tr16:
//line lex.rl:218
 buf = new(bytes.Buffer) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr19:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr86:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st8
tr100:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr111:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr116:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr119:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr122:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr125:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr128:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr131:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr134:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
tr141:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:13170
		switch data[p] {
		case 39:
			goto st111
		case 92:
			goto st38
		}
		goto tr19
tr17:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st111
tr101:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st111
tr112:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st111
tr117:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st111
tr120:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st111
tr123:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st111
tr126:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st111
tr129:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st111
tr132:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st111
tr135:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st111
tr142:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line lex.go:13227
		switch data[p] {
		case 32:
			goto tr632
		case 33:
			goto tr633
		case 34:
			goto tr634
		case 35:
			goto tr635
		case 36:
			goto tr636
		case 38:
			goto tr638
		case 39:
			goto tr19
		case 43:
			goto tr640
		case 45:
			goto tr641
		case 46:
			goto tr642
		case 47:
			goto tr643
		case 48:
			goto tr644
		case 58:
			goto tr646
		case 60:
			goto tr647
		case 62:
			goto tr648
		case 63:
			goto tr649
		case 64:
			goto tr650
		case 88:
			goto tr652
		case 95:
			goto tr653
		case 98:
			goto tr655
		case 101:
			goto tr656
		case 120:
			goto tr657
		case 124:
			goto tr658
		case 126:
			goto tr659
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr637
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr632
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr651
					}
				case data[p] >= 59:
					goto tr637
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr637
					}
				case data[p] >= 97:
					goto tr654
				}
			default:
				goto tr637
			}
		default:
			goto tr645
		}
		goto tr660
tr174:
//line lex.rl:41
 mark = p 
	goto st112
tr204:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.rl:41
 mark = p 
	goto st112
tr233:
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
			emit(lex.IDENT, string(b))
		
//line lex.rl:41
 mark = p 
	goto st112
tr267:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st112
tr296:
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
			emit(lex.SCONST, string(b))
		
//line lex.rl:41
 mark = p 
	goto st112
tr972:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:41
 mark = p 
	goto st112
tr357:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
//line lex.rl:41
 mark = p 
	goto st112
tr387:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
//line lex.rl:41
 mark = p 
	goto st112
tr419:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
//line lex.rl:41
 mark = p 
	goto st112
tr448:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:41
 mark = p 
	goto st112
tr482:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
//line lex.rl:41
 mark = p 
	goto st112
tr511:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr542:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr574:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st112
tr604:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
//line lex.rl:41
 mark = p 
	goto st112
tr656:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st112
tr704:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st112
tr734:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:41
 mark = p 
	goto st112
tr763:
//line lex.rl:294
 emitToken(lex.CONCAT) 
//line lex.rl:41
 mark = p 
	goto st112
tr792:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
//line lex.rl:41
 mark = p 
	goto st112
tr821:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:41
 mark = p 
	goto st112
tr851:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
//line lex.rl:41
 mark = p 
	goto st112
tr880:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr909:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr938:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr967:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:41
 mark = p 
	goto st112
tr1022:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
//line lex.rl:41
 mark = p 
	goto st112
tr1051:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
//line lex.rl:41
 mark = p 
	goto st112
tr1078:
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
			emit(lex.PLACEHOLDER, s)
		
//line lex.rl:41
 mark = p 
	goto st112
tr1107:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:41
 mark = p 
	goto st112
tr1136:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:41
 mark = p 
	goto st112
tr1165:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:41
 mark = p 
	goto st112
tr1194:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:41
 mark = p 
	goto st112
tr1224:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:41
 mark = p 
	goto st112
tr1253:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:41
 mark = p 
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line lex.go:13613
		switch data[p] {
		case 32:
			goto tr609
		case 33:
			goto tr610
		case 34:
			goto tr611
		case 35:
			goto tr612
		case 36:
			goto st103
		case 38:
			goto tr615
		case 39:
			goto st9
		case 43:
			goto tr617
		case 45:
			goto tr618
		case 46:
			goto tr619
		case 47:
			goto tr620
		case 58:
			goto tr621
		case 60:
			goto tr622
		case 62:
			goto tr623
		case 63:
			goto tr624
		case 64:
			goto tr625
		case 95:
			goto st103
		case 124:
			goto tr627
		case 126:
			goto tr628
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr614
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr609
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr626
					}
				case data[p] >= 59:
					goto tr614
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr614
					}
				case data[p] >= 97:
					goto st103
				}
			default:
				goto tr614
			}
		default:
			goto st103
		}
		goto tr629
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
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
	goto st10
tr25:
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr28:
//line lex.rl:229
 buf.WriteByte(data[p]) 
	goto st10
tr42:
//line lex.rl:215
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr53:
//line lex.rl:209
 buf.WriteRune(rn) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr58:
//line lex.rl:198
 buf.WriteByte(ch) 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr61:
//line lex.rl:186
 buf.WriteByte('\a') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr64:
//line lex.rl:187
 buf.WriteByte('\b') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr67:
//line lex.rl:188
 buf.WriteByte('\f') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr70:
//line lex.rl:189
 buf.WriteByte('\n') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr73:
//line lex.rl:190
 buf.WriteByte('\r') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr76:
//line lex.rl:191
 buf.WriteByte('\t') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
tr83:
//line lex.rl:192
 buf.WriteByte('\v') 
//line lex.rl:223
 buf.WriteByte(data[p]) 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line lex.go:13793
		switch data[p] {
		case 39:
			goto st113
		case 92:
			goto st11
		}
		goto tr25
tr23:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st113
tr43:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st113
tr54:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st113
tr59:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st113
tr62:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st113
tr65:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st113
tr68:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st113
tr71:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st113
tr74:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st113
tr77:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st113
tr84:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line lex.go:13850
		switch data[p] {
		case 32:
			goto tr711
		case 33:
			goto tr712
		case 34:
			goto tr713
		case 35:
			goto tr714
		case 36:
			goto tr715
		case 38:
			goto tr717
		case 39:
			goto tr25
		case 43:
			goto tr718
		case 45:
			goto tr719
		case 46:
			goto tr720
		case 47:
			goto tr721
		case 48:
			goto tr722
		case 58:
			goto tr724
		case 60:
			goto tr725
		case 62:
			goto tr726
		case 63:
			goto tr727
		case 64:
			goto tr728
		case 88:
			goto tr730
		case 95:
			goto tr731
		case 98:
			goto tr733
		case 101:
			goto tr734
		case 120:
			goto tr735
		case 124:
			goto tr736
		case 126:
			goto tr737
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr716
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr711
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr729
					}
				case data[p] >= 59:
					goto tr716
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr716
					}
				case data[p] >= 97:
					goto tr732
				}
			default:
				goto tr716
			}
		default:
			goto tr723
		}
		goto tr738
tr24:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st11
tr44:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st11
tr55:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st11
tr60:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st11
tr63:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st11
tr66:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st11
tr69:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st11
tr72:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st11
tr75:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st11
tr78:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st11
tr85:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line lex.go:13998
		switch data[p] {
		case 85:
			goto tr30
		case 88:
			goto st24
		case 97:
			goto st27
		case 98:
			goto st28
		case 102:
			goto st29
		case 110:
			goto st30
		case 114:
			goto st31
		case 116:
			goto st32
		case 117:
			goto tr38
		case 118:
			goto st37
		case 120:
			goto st24
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr29
		}
		goto tr28
tr29:
//line lex.rl:213
 ch = 0 
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:14038
		if 48 <= data[p] && data[p] <= 55 {
			goto tr40
		}
		goto st0
tr40:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:14052
		if 48 <= data[p] && data[p] <= 55 {
			goto tr41
		}
		goto st0
tr41:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line lex.go:14066
		switch data[p] {
		case 39:
			goto tr43
		case 92:
			goto tr44
		}
		goto tr42
tr30:
//line lex.rl:208
 rn = 0 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:14083
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr45
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr45
			}
		default:
			goto tr45
		}
		goto st0
tr45:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:14106
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr46
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr46:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line lex.go:14129
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
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:14152
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
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:14175
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
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:14198
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
//line lex.go:14221
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
//line lex.go:14244
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
tr82:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:14271
		switch data[p] {
		case 39:
			goto tr54
		case 92:
			goto tr55
		}
		goto tr53
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
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
//line lex.rl:196
 ch = 0 
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line lex.go:14308
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
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:14331
		switch data[p] {
		case 39:
			goto tr59
		case 92:
			goto tr60
		}
		goto tr58
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 39:
			goto tr62
		case 92:
			goto tr63
		}
		goto tr61
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 39:
			goto tr65
		case 92:
			goto tr66
		}
		goto tr64
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 39:
			goto tr68
		case 92:
			goto tr69
		}
		goto tr67
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 39:
			goto tr71
		case 92:
			goto tr72
		}
		goto tr70
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 39:
			goto tr74
		case 92:
			goto tr75
		}
		goto tr73
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr77
		case 92:
			goto tr78
		}
		goto tr76
tr38:
//line lex.rl:208
 rn = 0 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line lex.go:14420
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr79
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr79
			}
		default:
			goto tr79
		}
		goto st0
tr79:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line lex.go:14443
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr80
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr80
			}
		default:
			goto tr80
		}
		goto st0
tr80:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line lex.go:14466
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr81
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr81
			}
		default:
			goto tr81
		}
		goto st0
tr81:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line lex.go:14489
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr82
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr82
			}
		default:
			goto tr82
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 39:
			goto tr84
		case 92:
			goto tr85
		}
		goto tr83
tr18:
//line lex.rl:218
 buf = new(bytes.Buffer) 
	goto st38
tr102:
//line lex.rl:215
 buf.WriteByte(ch) 
	goto st38
tr113:
//line lex.rl:209
 buf.WriteRune(rn) 
	goto st38
tr118:
//line lex.rl:198
 buf.WriteByte(ch) 
	goto st38
tr121:
//line lex.rl:186
 buf.WriteByte('\a') 
	goto st38
tr124:
//line lex.rl:187
 buf.WriteByte('\b') 
	goto st38
tr127:
//line lex.rl:188
 buf.WriteByte('\f') 
	goto st38
tr130:
//line lex.rl:189
 buf.WriteByte('\n') 
	goto st38
tr133:
//line lex.rl:190
 buf.WriteByte('\r') 
	goto st38
tr136:
//line lex.rl:191
 buf.WriteByte('\t') 
	goto st38
tr143:
//line lex.rl:192
 buf.WriteByte('\v') 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line lex.go:14564
		switch data[p] {
		case 85:
			goto tr88
		case 88:
			goto st51
		case 97:
			goto st54
		case 98:
			goto st55
		case 102:
			goto st56
		case 110:
			goto st57
		case 114:
			goto st58
		case 116:
			goto st59
		case 117:
			goto tr96
		case 118:
			goto st64
		case 120:
			goto st51
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr87
		}
		goto tr86
tr87:
//line lex.rl:213
 ch = 0 
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:14604
		if 48 <= data[p] && data[p] <= 55 {
			goto tr98
		}
		goto st0
tr98:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lex.go:14618
		if 48 <= data[p] && data[p] <= 55 {
			goto tr99
		}
		goto st0
tr99:
//line lex.rl:214
 ch = (ch << 3) | data[p] - '0' 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:14632
		switch data[p] {
		case 39:
			goto tr101
		case 92:
			goto tr102
		}
		goto tr100
tr88:
//line lex.rl:208
 rn = 0 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line lex.go:14649
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr103
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr103
			}
		default:
			goto tr103
		}
		goto st0
tr103:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lex.go:14672
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr104
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto st0
tr104:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lex.go:14695
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr105
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr105
			}
		default:
			goto tr105
		}
		goto st0
tr105:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:14718
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr106
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
tr106:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lex.go:14741
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr107
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr107
			}
		default:
			goto tr107
		}
		goto st0
tr107:
//line lex.rl:206
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:14764
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
//line lex.go:14787
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
//line lex.go:14810
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
tr140:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:14837
		switch data[p] {
		case 39:
			goto tr112
		case 92:
			goto tr113
		}
		goto tr111
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
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
//line lex.rl:196
 ch = 0 
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line lex.go:14874
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
//line lex.rl:197
 ch = (ch << 4) | unhex(data[p]) 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lex.go:14897
		switch data[p] {
		case 39:
			goto tr117
		case 92:
			goto tr118
		}
		goto tr116
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 39:
			goto tr120
		case 92:
			goto tr121
		}
		goto tr119
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 39:
			goto tr123
		case 92:
			goto tr124
		}
		goto tr122
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 39:
			goto tr126
		case 92:
			goto tr127
		}
		goto tr125
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 39:
			goto tr129
		case 92:
			goto tr130
		}
		goto tr128
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 39:
			goto tr132
		case 92:
			goto tr133
		}
		goto tr131
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 39:
			goto tr135
		case 92:
			goto tr136
		}
		goto tr134
tr96:
//line lex.rl:208
 rn = 0 
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line lex.go:14986
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr137
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr137
			}
		default:
			goto tr137
		}
		goto st0
tr137:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line lex.go:15009
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr138
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr138
			}
		default:
			goto tr138
		}
		goto st0
tr138:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line lex.go:15032
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr139
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr139
			}
		default:
			goto tr139
		}
		goto st0
tr139:
//line lex.rl:203
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line lex.go:15055
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr140
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr140
			}
		default:
			goto tr140
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 39:
			goto tr142
		case 92:
			goto tr143
		}
		goto tr141
tr15:
//line lex.rl:250
 ch = 0 
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st65
tr146:
//line lex.rl:252
 buf.WriteByte(ch) 
//line lex.rl:250
 ch = 0 
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lex.go:15100
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
//line lex.rl:251
 ch = (ch << 4) | unhex(data[p]) 
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lex.go:15123
		if data[p] == 39 {
			goto tr145
		}
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
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 32:
			goto tr739
		case 33:
			goto tr740
		case 34:
			goto tr741
		case 35:
			goto tr742
		case 36:
			goto tr743
		case 38:
			goto tr745
		case 39:
			goto tr746
		case 43:
			goto tr747
		case 45:
			goto tr748
		case 46:
			goto tr749
		case 47:
			goto tr750
		case 48:
			goto tr751
		case 58:
			goto tr753
		case 60:
			goto tr754
		case 62:
			goto tr755
		case 63:
			goto tr756
		case 64:
			goto tr757
		case 88:
			goto tr759
		case 95:
			goto tr760
		case 98:
			goto tr762
		case 101:
			goto tr763
		case 120:
			goto tr764
		case 124:
			goto tr765
		case 126:
			goto tr766
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr744
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr739
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr758
					}
				case data[p] >= 59:
					goto tr744
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr744
					}
				case data[p] >= 97:
					goto tr761
				}
			default:
				goto tr744
			}
		default:
			goto tr752
		}
		goto tr767
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 32:
			goto tr768
		case 33:
			goto tr769
		case 34:
			goto tr770
		case 35:
			goto tr771
		case 36:
			goto tr772
		case 38:
			goto tr774
		case 39:
			goto tr775
		case 43:
			goto tr776
		case 45:
			goto tr777
		case 46:
			goto tr778
		case 47:
			goto tr779
		case 48:
			goto tr780
		case 58:
			goto tr782
		case 60:
			goto tr783
		case 62:
			goto tr784
		case 63:
			goto tr785
		case 64:
			goto tr786
		case 88:
			goto tr788
		case 95:
			goto tr789
		case 98:
			goto tr791
		case 101:
			goto tr792
		case 120:
			goto tr793
		case 124:
			goto tr794
		case 126:
			goto tr795
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr773
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr768
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr787
					}
				case data[p] >= 59:
					goto tr773
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr773
					}
				case data[p] >= 97:
					goto tr790
				}
			default:
				goto tr773
			}
		default:
			goto tr781
		}
		goto tr796
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto tr797
		case 33:
			goto tr798
		case 34:
			goto tr799
		case 35:
			goto tr800
		case 36:
			goto tr801
		case 38:
			goto tr803
		case 39:
			goto tr804
		case 43:
			goto tr805
		case 45:
			goto tr806
		case 46:
			goto tr807
		case 47:
			goto tr808
		case 48:
			goto tr809
		case 58:
			goto tr811
		case 60:
			goto tr812
		case 62:
			goto tr813
		case 63:
			goto tr814
		case 64:
			goto tr815
		case 88:
			goto tr817
		case 95:
			goto tr818
		case 98:
			goto tr820
		case 101:
			goto tr821
		case 120:
			goto tr822
		case 124:
			goto tr823
		case 126:
			goto tr824
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr802
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr797
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr816
					}
				case data[p] >= 59:
					goto tr802
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr802
					}
				case data[p] >= 97:
					goto tr819
				}
			default:
				goto tr802
			}
		default:
			goto tr810
		}
		goto tr825
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto tr826
		case 33:
			goto tr827
		case 34:
			goto tr828
		case 35:
			goto tr829
		case 36:
			goto tr830
		case 38:
			goto tr832
		case 39:
			goto tr833
		case 43:
			goto tr834
		case 45:
			goto tr835
		case 46:
			goto tr836
		case 47:
			goto tr837
		case 48:
			goto tr838
		case 58:
			goto tr840
		case 59:
			goto tr831
		case 60:
			goto tr841
		case 61:
			goto st118
		case 62:
			goto tr843
		case 63:
			goto tr844
		case 64:
			goto tr845
		case 88:
			goto tr847
		case 95:
			goto tr848
		case 98:
			goto tr850
		case 101:
			goto tr851
		case 120:
			goto tr852
		case 124:
			goto tr853
		case 126:
			goto tr854
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr831
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr826
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr846
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr831
					}
				case data[p] >= 97:
					goto tr849
				}
			default:
				goto tr831
			}
		default:
			goto tr839
		}
		goto tr855
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 32:
			goto tr856
		case 33:
			goto tr857
		case 34:
			goto tr858
		case 35:
			goto tr859
		case 36:
			goto tr860
		case 38:
			goto tr862
		case 39:
			goto tr863
		case 43:
			goto tr864
		case 45:
			goto tr865
		case 46:
			goto tr866
		case 47:
			goto tr867
		case 48:
			goto tr868
		case 58:
			goto tr870
		case 60:
			goto tr871
		case 62:
			goto tr872
		case 63:
			goto tr873
		case 64:
			goto tr874
		case 88:
			goto tr876
		case 95:
			goto tr877
		case 98:
			goto tr879
		case 101:
			goto tr880
		case 120:
			goto tr881
		case 124:
			goto tr882
		case 126:
			goto tr883
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr861
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr856
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr875
					}
				case data[p] >= 59:
					goto tr861
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr861
					}
				case data[p] >= 97:
					goto tr878
				}
			default:
				goto tr861
			}
		default:
			goto tr869
		}
		goto tr884
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto tr885
		case 33:
			goto tr886
		case 34:
			goto tr887
		case 35:
			goto tr888
		case 36:
			goto tr889
		case 38:
			goto tr891
		case 39:
			goto tr892
		case 43:
			goto tr893
		case 45:
			goto tr894
		case 46:
			goto tr895
		case 47:
			goto tr896
		case 48:
			goto tr897
		case 58:
			goto tr899
		case 60:
			goto tr900
		case 62:
			goto tr901
		case 63:
			goto tr902
		case 64:
			goto tr903
		case 88:
			goto tr905
		case 95:
			goto tr906
		case 98:
			goto tr908
		case 101:
			goto tr909
		case 120:
			goto tr910
		case 124:
			goto tr911
		case 126:
			goto tr912
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr890
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr885
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr904
					}
				case data[p] >= 59:
					goto tr890
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr890
					}
				case data[p] >= 97:
					goto tr907
				}
			default:
				goto tr890
			}
		default:
			goto tr898
		}
		goto tr913
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto tr914
		case 33:
			goto tr915
		case 34:
			goto tr916
		case 35:
			goto tr917
		case 36:
			goto tr918
		case 38:
			goto tr920
		case 39:
			goto tr921
		case 43:
			goto tr922
		case 45:
			goto tr923
		case 46:
			goto tr924
		case 47:
			goto tr925
		case 48:
			goto tr926
		case 58:
			goto tr928
		case 60:
			goto tr929
		case 62:
			goto tr930
		case 63:
			goto tr931
		case 64:
			goto tr932
		case 88:
			goto tr934
		case 95:
			goto tr935
		case 98:
			goto tr937
		case 101:
			goto tr938
		case 120:
			goto tr939
		case 124:
			goto tr940
		case 126:
			goto tr941
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr919
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr914
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr933
					}
				case data[p] >= 59:
					goto tr919
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr919
					}
				case data[p] >= 97:
					goto tr936
				}
			default:
				goto tr919
			}
		default:
			goto tr927
		}
		goto tr942
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto tr943
		case 33:
			goto tr944
		case 34:
			goto tr945
		case 35:
			goto tr946
		case 36:
			goto tr947
		case 38:
			goto tr949
		case 39:
			goto tr950
		case 43:
			goto tr951
		case 45:
			goto tr952
		case 46:
			goto tr953
		case 47:
			goto tr954
		case 48:
			goto tr955
		case 58:
			goto tr957
		case 60:
			goto tr958
		case 62:
			goto tr959
		case 63:
			goto tr960
		case 64:
			goto tr961
		case 88:
			goto tr963
		case 95:
			goto tr964
		case 98:
			goto tr966
		case 101:
			goto tr967
		case 120:
			goto tr968
		case 124:
			goto tr969
		case 126:
			goto tr970
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr948
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr943
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr962
					}
				case data[p] >= 59:
					goto tr948
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr948
					}
				case data[p] >= 97:
					goto tr965
				}
			default:
				goto tr948
			}
		default:
			goto tr956
		}
		goto tr971
tr323:
//line lex.rl:70

			isFconst = true
		
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lex.go:15974
		switch data[p] {
		case 43:
			goto st68
		case 45:
			goto st68
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr305
		case 33:
			goto tr306
		case 34:
			goto tr307
		case 35:
			goto tr308
		case 36:
			goto tr309
		case 38:
			goto tr311
		case 39:
			goto tr312
		case 43:
			goto tr313
		case 45:
			goto tr314
		case 46:
			goto tr315
		case 47:
			goto tr316
		case 58:
			goto tr317
		case 60:
			goto tr318
		case 62:
			goto tr319
		case 63:
			goto tr320
		case 64:
			goto tr321
		case 88:
			goto tr324
		case 95:
			goto tr325
		case 98:
			goto tr327
		case 101:
			goto tr972
		case 120:
			goto tr328
		case 124:
			goto tr329
		case 126:
			goto tr330
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr310
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr305
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr322
					}
				case data[p] >= 59:
					goto tr310
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr310
					}
				case data[p] >= 97:
					goto tr326
				}
			default:
				goto tr310
			}
		default:
			goto st122
		}
		goto tr331
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st123
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st123
			}
		default:
			goto st123
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto tr973
		case 33:
			goto tr974
		case 34:
			goto tr975
		case 35:
			goto tr976
		case 36:
			goto tr977
		case 38:
			goto tr979
		case 39:
			goto tr980
		case 43:
			goto tr981
		case 45:
			goto tr982
		case 46:
			goto tr983
		case 47:
			goto tr984
		case 58:
			goto tr985
		case 59:
			goto tr978
		case 60:
			goto tr986
		case 61:
			goto tr978
		case 62:
			goto tr987
		case 63:
			goto tr988
		case 64:
			goto tr989
		case 88:
			goto tr991
		case 95:
			goto tr992
		case 120:
			goto tr994
		case 124:
			goto tr995
		case 126:
			goto tr996
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr978
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr973
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto tr978
					}
				case data[p] >= 71:
					goto tr990
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr978
					}
				case data[p] >= 103:
					goto tr993
				}
			default:
				goto st123
			}
		default:
			goto st123
		}
		goto tr997
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto tr998
		case 33:
			goto tr999
		case 34:
			goto tr1000
		case 35:
			goto tr1001
		case 36:
			goto tr1002
		case 38:
			goto tr1004
		case 39:
			goto tr1005
		case 43:
			goto tr1006
		case 45:
			goto tr1007
		case 46:
			goto tr1008
		case 47:
			goto tr1009
		case 48:
			goto tr1010
		case 58:
			goto tr1012
		case 60:
			goto tr1013
		case 62:
			goto st125
		case 63:
			goto tr1015
		case 64:
			goto tr1016
		case 88:
			goto tr1018
		case 95:
			goto tr1019
		case 98:
			goto tr1021
		case 101:
			goto tr1022
		case 120:
			goto tr1023
		case 124:
			goto tr1024
		case 126:
			goto tr1025
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1003
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr998
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1017
					}
				case data[p] >= 59:
					goto tr1003
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1003
					}
				case data[p] >= 97:
					goto tr1020
				}
			default:
				goto tr1003
			}
		default:
			goto tr1011
		}
		goto tr1026
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto tr1027
		case 33:
			goto tr1028
		case 34:
			goto tr1029
		case 35:
			goto tr1030
		case 36:
			goto tr1031
		case 38:
			goto tr1033
		case 39:
			goto tr1034
		case 43:
			goto tr1035
		case 45:
			goto tr1036
		case 46:
			goto tr1037
		case 47:
			goto tr1038
		case 48:
			goto tr1039
		case 58:
			goto tr1041
		case 60:
			goto tr1042
		case 62:
			goto tr1043
		case 63:
			goto tr1044
		case 64:
			goto tr1045
		case 88:
			goto tr1047
		case 95:
			goto tr1048
		case 98:
			goto tr1050
		case 101:
			goto tr1051
		case 120:
			goto tr1052
		case 124:
			goto tr1053
		case 126:
			goto tr1054
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1032
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1027
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1046
					}
				case data[p] >= 59:
					goto tr1032
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1032
					}
				case data[p] >= 97:
					goto tr1049
				}
			default:
				goto tr1032
			}
		default:
			goto tr1040
		}
		goto tr1055
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 32:
			goto tr1056
		case 33:
			goto tr1057
		case 34:
			goto tr1058
		case 35:
			goto tr1059
		case 36:
			goto tr1060
		case 38:
			goto tr1062
		case 39:
			goto tr1063
		case 43:
			goto tr1064
		case 45:
			goto tr1065
		case 46:
			goto tr1066
		case 47:
			goto tr1067
		case 58:
			goto tr1068
		case 60:
			goto tr1069
		case 62:
			goto tr1070
		case 63:
			goto tr1071
		case 64:
			goto tr1072
		case 88:
			goto tr1074
		case 95:
			goto tr1075
		case 98:
			goto tr1077
		case 101:
			goto tr1078
		case 120:
			goto tr1079
		case 124:
			goto tr1080
		case 126:
			goto tr1081
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1061
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1056
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1073
					}
				case data[p] >= 59:
					goto tr1061
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1061
					}
				case data[p] >= 97:
					goto tr1076
				}
			default:
				goto tr1061
			}
		default:
			goto st126
		}
		goto tr1082
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 32:
			goto tr1083
		case 33:
			goto tr1084
		case 34:
			goto tr1085
		case 35:
			goto tr1086
		case 36:
			goto tr1087
		case 38:
			goto tr1089
		case 39:
			goto tr1090
		case 43:
			goto tr1091
		case 45:
			goto tr1092
		case 46:
			goto tr1093
		case 47:
			goto tr1094
		case 48:
			goto tr1095
		case 58:
			goto tr1097
		case 60:
			goto tr1098
		case 62:
			goto tr1099
		case 63:
			goto tr1100
		case 64:
			goto tr1101
		case 88:
			goto tr1103
		case 95:
			goto tr1104
		case 98:
			goto tr1106
		case 101:
			goto tr1107
		case 120:
			goto tr1108
		case 124:
			goto tr1109
		case 126:
			goto tr1110
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1088
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1083
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1102
					}
				case data[p] >= 59:
					goto tr1088
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1088
					}
				case data[p] >= 97:
					goto tr1105
				}
			default:
				goto tr1088
			}
		default:
			goto tr1096
		}
		goto tr1111
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto tr1112
		case 33:
			goto tr1113
		case 34:
			goto tr1114
		case 35:
			goto tr1115
		case 36:
			goto tr1116
		case 38:
			goto tr1118
		case 39:
			goto tr1119
		case 43:
			goto tr1120
		case 45:
			goto tr1121
		case 46:
			goto tr1122
		case 47:
			goto tr1123
		case 48:
			goto tr1124
		case 58:
			goto tr1126
		case 60:
			goto tr1127
		case 62:
			goto st129
		case 63:
			goto tr1129
		case 64:
			goto tr1130
		case 88:
			goto tr1132
		case 95:
			goto tr1133
		case 98:
			goto tr1135
		case 101:
			goto tr1136
		case 120:
			goto tr1137
		case 124:
			goto tr1138
		case 126:
			goto tr1139
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1117
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1112
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1131
					}
				case data[p] >= 59:
					goto tr1117
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1117
					}
				case data[p] >= 97:
					goto tr1134
				}
			default:
				goto tr1117
			}
		default:
			goto tr1125
		}
		goto tr1140
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 32:
			goto tr1141
		case 33:
			goto tr1142
		case 34:
			goto tr1143
		case 35:
			goto tr1144
		case 36:
			goto tr1145
		case 38:
			goto tr1147
		case 39:
			goto tr1148
		case 43:
			goto tr1149
		case 45:
			goto tr1150
		case 46:
			goto tr1151
		case 47:
			goto tr1152
		case 48:
			goto tr1153
		case 58:
			goto tr1155
		case 60:
			goto tr1156
		case 62:
			goto tr1157
		case 63:
			goto tr1158
		case 64:
			goto tr1159
		case 88:
			goto tr1161
		case 95:
			goto tr1162
		case 98:
			goto tr1164
		case 101:
			goto tr1165
		case 120:
			goto tr1166
		case 124:
			goto tr1167
		case 126:
			goto tr1168
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1146
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1141
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1160
					}
				case data[p] >= 59:
					goto tr1146
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1146
					}
				case data[p] >= 97:
					goto tr1163
				}
			default:
				goto tr1146
			}
		default:
			goto tr1154
		}
		goto tr1169
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto tr1170
		case 33:
			goto tr1171
		case 34:
			goto tr1172
		case 35:
			goto tr1173
		case 36:
			goto tr1174
		case 38:
			goto tr1176
		case 39:
			goto tr1177
		case 43:
			goto tr1178
		case 45:
			goto tr1179
		case 46:
			goto tr1180
		case 47:
			goto tr1181
		case 48:
			goto tr1182
		case 58:
			goto tr1184
		case 60:
			goto tr1185
		case 62:
			goto tr1186
		case 63:
			goto tr1187
		case 64:
			goto tr1188
		case 88:
			goto tr1190
		case 95:
			goto tr1191
		case 98:
			goto tr1193
		case 101:
			goto tr1194
		case 120:
			goto tr1195
		case 124:
			goto tr1196
		case 126:
			goto tr1197
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1175
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1170
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1189
					}
				case data[p] >= 59:
					goto tr1175
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1175
					}
				case data[p] >= 97:
					goto tr1192
				}
			default:
				goto tr1175
			}
		default:
			goto tr1183
		}
		goto tr1198
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr1199
		case 33:
			goto tr1200
		case 34:
			goto tr1201
		case 35:
			goto tr1202
		case 36:
			goto tr1203
		case 38:
			goto tr1205
		case 39:
			goto tr1206
		case 42:
			goto st132
		case 43:
			goto tr1208
		case 45:
			goto tr1209
		case 46:
			goto tr1210
		case 47:
			goto tr1211
		case 48:
			goto tr1212
		case 58:
			goto tr1214
		case 60:
			goto tr1215
		case 62:
			goto tr1216
		case 63:
			goto tr1217
		case 64:
			goto tr1218
		case 88:
			goto tr1220
		case 95:
			goto tr1221
		case 98:
			goto tr1223
		case 101:
			goto tr1224
		case 120:
			goto tr1225
		case 124:
			goto tr1226
		case 126:
			goto tr1227
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1204
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1199
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1219
					}
				case data[p] >= 59:
					goto tr1204
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1204
					}
				case data[p] >= 97:
					goto tr1222
				}
			default:
				goto tr1204
			}
		default:
			goto tr1213
		}
		goto tr1228
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 32:
			goto tr1229
		case 33:
			goto tr1230
		case 34:
			goto tr1231
		case 35:
			goto tr1232
		case 36:
			goto tr1233
		case 38:
			goto tr1235
		case 39:
			goto tr1236
		case 43:
			goto tr1237
		case 45:
			goto tr1238
		case 46:
			goto tr1239
		case 47:
			goto tr1240
		case 48:
			goto tr1241
		case 58:
			goto tr1243
		case 60:
			goto tr1244
		case 62:
			goto tr1245
		case 63:
			goto tr1246
		case 64:
			goto tr1247
		case 88:
			goto tr1249
		case 95:
			goto tr1250
		case 98:
			goto tr1252
		case 101:
			goto tr1253
		case 120:
			goto tr1254
		case 124:
			goto tr1255
		case 126:
			goto tr1256
		case 127:
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
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1234
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1229
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1248
					}
				case data[p] >= 59:
					goto tr1234
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1234
					}
				case data[p] >= 97:
					goto tr1251
				}
			default:
				goto tr1234
			}
		default:
			goto tr1242
		}
		goto tr1257
	st_out:
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
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
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 126:
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
			emit(lex.PLACEHOLDER, s)
		
		case 84, 89, 90, 122:
//line lex.rl:56

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
		case 123:
//line lex.rl:73

			str()
			emit(lex.ICONST, s)
		
		case 103, 105, 110, 112:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
		case 73:
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
			emit(lex.IDENT, string(b))
		
		case 80:
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
			emit(lex.SCONST, string(b))
		
		case 106, 111:
//line lex.rl:235

			emit(lex.BCONST, buf.String())
		
		case 113:
//line lex.rl:239

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
		case 71, 72, 74, 75, 76, 77, 79, 81, 82, 85, 87, 91, 94, 97, 99, 101, 104, 108:
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
		case 86:
//line lex.rl:271
 emitToken(lex.DOT_DOT) 
		case 130:
//line lex.rl:273
 emitToken(lex.NOT_EQUALS) 
		case 132:
//line lex.rl:274
 emitToken(lex.NOT_REGIMATCH) 
		case 131:
//line lex.rl:275
 emitToken(lex.NOT_REGMATCH) 
		case 115:
//line lex.rl:277
 emitToken(lex.HELPTOKEN) 
		case 116:
//line lex.rl:278
 emitToken(lex.JSON_SOME_EXISTS) 
		case 100:
//line lex.rl:279
 emitToken(lex.JSON_ALL_EXISTS) 
		case 96:
//line lex.rl:281
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
		case 95:
//line lex.rl:282
 emitToken(lex.LSHIFT) 
		case 120:
//line lex.rl:283
 emitToken(lex.NOT_EQUALS) 
		case 119:
//line lex.rl:284
 emitToken(lex.LESS_EQUALS) 
		case 121:
//line lex.rl:285
 emitToken(lex.CONTAINED_BY) 
		case 118:
//line lex.rl:287
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
		case 117:
//line lex.rl:288
 emitToken(lex.RSHIFT) 
		case 98:
//line lex.rl:289
 emitToken(lex.GREATER_EQUALS) 
		case 93:
//line lex.rl:291
 emitToken(lex.TYPEANNOTATE) 
		case 92:
//line lex.rl:292
 emitToken(lex.TYPECAST) 
		case 114:
//line lex.rl:294
 emitToken(lex.CONCAT) 
		case 88:
//line lex.rl:296
 emitToken(lex.FLOORDIV) 
		case 109:
//line lex.rl:298
 emitToken(lex.REGIMATCH) 
		case 102:
//line lex.rl:300
 emitToken(lex.CONTAINS) 
		case 78:
//line lex.rl:302
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
		case 125:
//line lex.rl:304
 emitToken(lex.FETCHTEXT) 
		case 124:
//line lex.rl:305
 emitToken(lex.FETCHVAL) 
		case 129:
//line lex.rl:307
 emitToken(lex.FETCHTEXT_PATH) 
		case 128:
//line lex.rl:308
 emitToken(lex.FETCHVAL_PATH) 
		case 127:
//line lex.rl:309
 emitToken(lex.REMOVE_PATH) 
		case 107:
//line lex.rl:77

			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
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
			emit(lex.IDENT, s)
		
//line lex.rl:269
 emitToken(Tok(data[p-1])) 
//line lex.go:17506
		}
	}

	_out: {}
	}

//line lex.rl:318


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
