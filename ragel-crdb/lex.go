
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
const scanner_start int = 71
const scanner_first_final int = 71
const scanner_error int = 0

const scanner_en_main int = 71


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

	
//line lex.go:53
	{
	cs = scanner_start
	}

//line lex.go:58
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 71:
		goto st_case_71
	case 0:
		goto st_case_0
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 1:
		goto st_case_1
	case 74:
		goto st_case_74
	case 2:
		goto st_case_2
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
	case 80:
		goto st_case_80
	case 3:
		goto st_case_3
	case 81:
		goto st_case_81
	case 4:
		goto st_case_4
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 5:
		goto st_case_5
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
	case 106:
		goto st_case_106
	case 6:
		goto st_case_6
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 7:
		goto st_case_7
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 116:
		goto st_case_116
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
	case 67:
		goto st_case_67
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
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 125:
		goto st_case_125
	case 70:
		goto st_case_70
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
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	}
	goto st_out
tr182:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st71
tr213:
//line lex.rl:130

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
tr248:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st71
tr278:
//line lex.rl:164

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
tr312:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st71
tr341:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st71
tr372:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st71
tr405:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st71
tr435:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st71
tr469:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st71
tr500:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st71
tr532:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st71
tr565:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st71
tr596:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st71
tr626:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st71
tr648:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st71
tr679:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st71
tr709:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st71
tr728:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st71
tr760:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st71
tr789:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st71
tr819:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st71
tr849:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st71
tr879:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st71
tr910:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st71
tr940:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st71
tr970:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st71
tr1000:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st71
tr1031:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st71
tr1056:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st71
tr1086:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st71
tr1116:
//line lex.rl:43

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
tr1144:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st71
tr1174:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st71
tr1204:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st71
tr1234:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st71
tr1264:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st71
tr1295:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lex.go:624
		switch data[p] {
		case 32:
			goto st71
		case 33:
			goto st72
		case 34:
			goto tr154
		case 35:
			goto st75
		case 36:
			goto tr156
		case 38:
			goto st78
		case 39:
			goto tr159
		case 43:
			goto tr160
		case 45:
			goto tr161
		case 46:
			goto tr162
		case 47:
			goto st88
		case 48:
			goto tr164
		case 58:
			goto st92
		case 60:
			goto st95
		case 62:
			goto st98
		case 63:
			goto st100
		case 64:
			goto st102
		case 66:
			goto tr172
		case 88:
			goto tr173
		case 95:
			goto tr174
		case 98:
			goto tr176
		case 101:
			goto tr177
		case 120:
			goto tr178
		case 124:
			goto st105
		case 126:
			goto st111
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
						goto st77
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto st71
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr171
					}
				case data[p] >= 59:
					goto st77
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto st77
					}
				case data[p] >= 97:
					goto tr175
				}
			default:
				goto st77
			}
		default:
			goto tr165
		}
		goto tr181
st_case_0:
	st0:
		cs = 0
		goto _out
tr183:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st72
tr214:
//line lex.rl:130

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
		
	goto st72
tr249:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st72
tr279:
//line lex.rl:164

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
		
	goto st72
tr313:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st72
tr342:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st72
tr373:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st72
tr406:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st72
tr436:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st72
tr470:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st72
tr501:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st72
tr533:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st72
tr566:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st72
tr597:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st72
tr627:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st72
tr649:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st72
tr680:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st72
tr710:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st72
tr729:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st72
tr761:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st72
tr790:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st72
tr820:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st72
tr850:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st72
tr880:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st72
tr911:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st72
tr941:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st72
tr971:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st72
tr1001:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st72
tr1032:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st72
tr1057:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st72
tr1087:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st72
tr1117:
//line lex.rl:43

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
		
	goto st72
tr1145:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st72
tr1175:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st72
tr1205:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st72
tr1235:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st72
tr1265:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st72
tr1296:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lex.go:1015
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 59:
			goto tr187
		case 60:
			goto tr197
		case 61:
			goto st133
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto st134
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr202
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
tr154:
//line lex.rl:40
 mark = p 
	goto st73
tr184:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st73
tr250:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st73
tr280:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st73
tr314:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st73
tr343:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st73
tr374:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st73
tr407:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st73
tr437:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st73
tr471:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st73
tr502:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr534:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr567:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st73
tr598:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st73
tr628:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:40
 mark = p 
	goto st73
tr650:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st73
tr681:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st73
tr711:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st73
tr730:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st73
tr762:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st73
tr791:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st73
tr821:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st73
tr851:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st73
tr881:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st73
tr912:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr942:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr972:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr1002:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st73
tr1033:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st73
tr1058:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st73
tr1088:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st73
tr1118:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st73
tr1146:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st73
tr1176:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st73
tr1206:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st73
tr1236:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st73
tr1266:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st73
tr1297:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lex.go:1451
		if data[p] == 34 {
			goto st74
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr2:
//line lex.rl:110
 isNotASCII = true 
	goto st1
tr3:
//line lex.rl:124
 numQuote++ 
	goto st1
tr5:
//line lex.rl:124
 numQuote++ 
//line lex.rl:110
 isNotASCII = true 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:1478
		if data[p] == 34 {
			goto st74
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:124
 numQuote++ 
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lex.go:1495
		switch data[p] {
		case 32:
			goto tr213
		case 33:
			goto tr214
		case 34:
			goto st2
		case 35:
			goto tr216
		case 36:
			goto tr217
		case 38:
			goto tr219
		case 39:
			goto tr220
		case 43:
			goto tr221
		case 45:
			goto tr222
		case 46:
			goto tr223
		case 47:
			goto tr224
		case 48:
			goto tr225
		case 58:
			goto tr227
		case 60:
			goto tr228
		case 62:
			goto tr229
		case 63:
			goto tr230
		case 64:
			goto tr231
		case 66:
			goto tr233
		case 88:
			goto tr234
		case 95:
			goto tr235
		case 98:
			goto tr237
		case 101:
			goto tr238
		case 120:
			goto tr239
		case 124:
			goto tr240
		case 126:
			goto tr241
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
						goto tr218
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr213
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr232
					}
				case data[p] >= 59:
					goto tr218
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr218
					}
				case data[p] >= 97:
					goto tr236
				}
			default:
				goto tr218
			}
		default:
			goto tr226
		}
		goto tr242
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
tr185:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st75
tr216:
//line lex.rl:130

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
		
	goto st75
tr251:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st75
tr281:
//line lex.rl:164

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
		
	goto st75
tr315:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st75
tr344:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st75
tr375:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st75
tr408:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st75
tr438:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st75
tr472:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st75
tr503:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st75
tr535:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st75
tr568:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st75
tr599:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st75
tr629:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st75
tr651:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st75
tr682:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st75
tr712:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st75
tr731:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st75
tr763:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st75
tr792:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st75
tr822:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st75
tr852:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st75
tr882:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st75
tr913:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st75
tr943:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st75
tr973:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st75
tr1003:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st75
tr1034:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st75
tr1059:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st75
tr1089:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st75
tr1119:
//line lex.rl:43

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
		
	goto st75
tr1147:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st75
tr1177:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st75
tr1207:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st75
tr1237:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st75
tr1267:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st75
tr1298:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lex.go:1894
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto st130
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto st131
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
tr156:
//line lex.rl:40
 mark = p 
	goto st76
tr186:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st76
tr217:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st76
tr252:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st76
tr282:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st76
tr316:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st76
tr345:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st76
tr376:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st76
tr409:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st76
tr439:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st76
tr473:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st76
tr504:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr536:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr569:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st76
tr600:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st76
tr652:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st76
tr683:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st76
tr732:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st76
tr764:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st76
tr793:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st76
tr823:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st76
tr853:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st76
tr883:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st76
tr914:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr944:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr974:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr1004:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st76
tr1035:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st76
tr1060:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st76
tr1090:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st76
tr1120:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st76
tr1148:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st76
tr1178:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st76
tr1208:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st76
tr1238:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st76
tr1268:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st76
tr1299:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lex.go:2299
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto st129
		}
		goto tr212
tr187:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st77
tr218:
//line lex.rl:130

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
tr253:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st77
tr283:
//line lex.rl:164

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
tr317:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st77
tr346:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st77
tr377:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st77
tr410:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st77
tr440:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st77
tr474:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st77
tr505:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st77
tr537:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st77
tr570:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st77
tr601:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st77
tr631:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st77
tr653:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st77
tr684:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st77
tr713:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st77
tr733:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st77
tr765:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st77
tr794:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st77
tr824:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st77
tr854:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st77
tr884:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st77
tr915:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st77
tr945:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st77
tr975:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st77
tr1005:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st77
tr1036:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st77
tr1061:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st77
tr1091:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st77
tr1121:
//line lex.rl:43

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
tr1149:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st77
tr1179:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st77
tr1209:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st77
tr1239:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st77
tr1269:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st77
tr1300:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lex.go:2684
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
tr188:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st78
tr219:
//line lex.rl:130

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
		
	goto st78
tr254:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st78
tr284:
//line lex.rl:164

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
		
	goto st78
tr318:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st78
tr347:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st78
tr378:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st78
tr411:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st78
tr441:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st78
tr475:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st78
tr506:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st78
tr538:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st78
tr571:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st78
tr602:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st78
tr632:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st78
tr654:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st78
tr685:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st78
tr714:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st78
tr734:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st78
tr766:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st78
tr795:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st78
tr825:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st78
tr855:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st78
tr885:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st78
tr916:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st78
tr946:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st78
tr976:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st78
tr1006:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st78
tr1037:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st78
tr1062:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st78
tr1092:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st78
tr1122:
//line lex.rl:43

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
		
	goto st78
tr1150:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st78
tr1180:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st78
tr1210:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st78
tr1240:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st78
tr1270:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st78
tr1301:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line lex.go:3071
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto st79
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto tr248
		case 33:
			goto tr249
		case 34:
			goto tr250
		case 35:
			goto tr251
		case 36:
			goto tr252
		case 38:
			goto tr254
		case 39:
			goto tr255
		case 43:
			goto tr256
		case 45:
			goto tr257
		case 46:
			goto tr258
		case 47:
			goto tr259
		case 48:
			goto tr260
		case 58:
			goto tr262
		case 60:
			goto tr263
		case 62:
			goto tr264
		case 63:
			goto tr265
		case 64:
			goto tr266
		case 66:
			goto tr268
		case 88:
			goto tr269
		case 95:
			goto tr270
		case 98:
			goto tr272
		case 101:
			goto tr273
		case 120:
			goto tr274
		case 124:
			goto tr275
		case 126:
			goto tr276
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
						goto tr253
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr248
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr267
					}
				case data[p] >= 59:
					goto tr253
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr253
					}
				case data[p] >= 97:
					goto tr271
				}
			default:
				goto tr253
			}
		default:
			goto tr261
		}
		goto tr277
tr159:
//line lex.rl:40
 mark = p 
	goto st80
tr189:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st80
tr220:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st80
tr255:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st80
tr319:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st80
tr348:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st80
tr379:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st80
tr412:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st80
tr442:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st80
tr476:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st80
tr507:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr539:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr572:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st80
tr603:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st80
tr633:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:40
 mark = p 
	goto st80
tr655:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st80
tr686:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st80
tr715:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st80
tr735:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st80
tr796:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st80
tr826:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st80
tr856:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st80
tr886:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st80
tr917:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr947:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr977:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr1007:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st80
tr1038:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st80
tr1063:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st80
tr1093:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st80
tr1123:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st80
tr1151:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st80
tr1181:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st80
tr1211:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st80
tr1241:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st80
tr1271:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st80
tr1302:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lex.go:3602
		if data[p] == 39 {
			goto st81
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr8:
//line lex.rl:110
 isNotASCII = true 
	goto st3
tr9:
//line lex.rl:158
 numQuote++ 
	goto st3
tr11:
//line lex.rl:158
 numQuote++ 
//line lex.rl:110
 isNotASCII = true 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:3629
		if data[p] == 39 {
			goto st81
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr10:
//line lex.rl:158
 numQuote++ 
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lex.go:3646
		switch data[p] {
		case 32:
			goto tr278
		case 33:
			goto tr279
		case 34:
			goto tr280
		case 35:
			goto tr281
		case 36:
			goto tr282
		case 38:
			goto tr284
		case 39:
			goto st4
		case 43:
			goto tr286
		case 45:
			goto tr287
		case 46:
			goto tr288
		case 47:
			goto tr289
		case 48:
			goto tr290
		case 58:
			goto tr292
		case 60:
			goto tr293
		case 62:
			goto tr294
		case 63:
			goto tr295
		case 64:
			goto tr296
		case 66:
			goto tr298
		case 88:
			goto tr299
		case 95:
			goto tr300
		case 98:
			goto tr302
		case 101:
			goto tr303
		case 120:
			goto tr304
		case 124:
			goto tr305
		case 126:
			goto tr306
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
						goto tr283
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr278
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr297
					}
				case data[p] >= 59:
					goto tr283
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr283
					}
				case data[p] >= 97:
					goto tr301
				}
			default:
				goto tr283
			}
		default:
			goto tr291
		}
		goto tr307
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
tr160:
//line lex.rl:40
 mark = p 
	goto st82
tr190:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st82
tr221:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st82
tr256:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st82
tr286:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st82
tr320:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st82
tr349:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st82
tr380:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st82
tr413:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st82
tr443:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st82
tr477:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st82
tr508:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr540:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr573:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st82
tr604:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st82
tr634:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:40
 mark = p 
	goto st82
tr656:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st82
tr687:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st82
tr716:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st82
tr736:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st82
tr767:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st82
tr797:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st82
tr827:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st82
tr857:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st82
tr887:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st82
tr918:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr948:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr978:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr1008:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st82
tr1039:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st82
tr1064:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st82
tr1094:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st82
tr1124:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st82
tr1152:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st82
tr1182:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st82
tr1212:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st82
tr1242:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st82
tr1273:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st82
tr1303:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lex.go:4125
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr308
		case 47:
			goto tr193
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr309
		}
		goto tr212
tr161:
//line lex.rl:40
 mark = p 
	goto st83
tr191:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st83
tr222:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st83
tr257:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st83
tr287:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st83
tr321:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st83
tr350:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st83
tr381:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st83
tr414:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st83
tr444:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st83
tr478:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st83
tr509:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr541:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr574:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st83
tr605:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st83
tr635:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:40
 mark = p 
	goto st83
tr657:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st83
tr688:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st83
tr717:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st83
tr737:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st83
tr768:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st83
tr798:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st83
tr828:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st83
tr858:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st83
tr888:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st83
tr919:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr949:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr979:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr1009:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st83
tr1040:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st83
tr1065:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st83
tr1095:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st83
tr1125:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st83
tr1153:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st83
tr1183:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st83
tr1213:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st83
tr1243:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st83
tr1274:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st83
tr1304:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lex.go:4590
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto st84
		case 46:
			goto tr308
		case 47:
			goto tr193
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto st127
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr309
		}
		goto tr212
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 10 {
			goto st71
		}
		goto st84
tr308:
//line lex.rl:69

			isFconst = true
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line lex.go:4709
		if 48 <= data[p] && data[p] <= 57 {
			goto st85
		}
		goto st0
tr402:
//line lex.rl:69

			isFconst = true
		
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lex.go:4725
		switch data[p] {
		case 32:
			goto tr312
		case 33:
			goto tr313
		case 34:
			goto tr314
		case 35:
			goto tr315
		case 36:
			goto tr316
		case 38:
			goto tr318
		case 39:
			goto tr319
		case 43:
			goto tr320
		case 45:
			goto tr321
		case 46:
			goto tr322
		case 47:
			goto tr323
		case 58:
			goto tr324
		case 60:
			goto tr325
		case 62:
			goto tr326
		case 63:
			goto tr327
		case 64:
			goto tr328
		case 66:
			goto tr330
		case 69:
			goto tr331
		case 88:
			goto tr332
		case 95:
			goto tr333
		case 98:
			goto tr335
		case 101:
			goto tr331
		case 120:
			goto tr336
		case 124:
			goto tr337
		case 126:
			goto tr338
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
						goto tr317
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr312
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr329
					}
				case data[p] >= 59:
					goto tr317
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr317
					}
				case data[p] >= 97:
					goto tr334
				}
			default:
				goto tr317
			}
		default:
			goto st85
		}
		goto tr339
tr162:
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr192:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr223:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr258:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr288:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr322:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr351:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr382:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr415:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr445:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr479:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr510:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr542:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr575:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr606:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr636:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr658:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr689:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr718:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr738:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr769:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr799:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr829:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr859:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr889:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr920:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr950:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr980:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1010:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1041:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1066:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1096:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1126:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1154:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1184:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1214:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1244:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1275:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
tr1305:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:69

			isFconst = true
		
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line lex.go:5348
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto st87
		case 47:
			goto tr193
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto st85
		}
		goto tr212
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 32:
			goto tr341
		case 33:
			goto tr342
		case 34:
			goto tr343
		case 35:
			goto tr344
		case 36:
			goto tr345
		case 38:
			goto tr347
		case 39:
			goto tr348
		case 43:
			goto tr349
		case 45:
			goto tr350
		case 46:
			goto tr351
		case 47:
			goto tr352
		case 48:
			goto tr353
		case 58:
			goto tr355
		case 60:
			goto tr356
		case 62:
			goto tr357
		case 63:
			goto tr358
		case 64:
			goto tr359
		case 66:
			goto tr361
		case 88:
			goto tr362
		case 95:
			goto tr363
		case 98:
			goto tr365
		case 101:
			goto tr366
		case 120:
			goto tr367
		case 124:
			goto tr368
		case 126:
			goto tr369
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
						goto tr346
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr341
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr360
					}
				case data[p] >= 59:
					goto tr346
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr346
					}
				case data[p] >= 97:
					goto tr364
				}
			default:
				goto tr346
			}
		default:
			goto tr354
		}
		goto tr370
tr193:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st88
tr224:
//line lex.rl:130

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
		
	goto st88
tr259:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st88
tr289:
//line lex.rl:164

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
		
	goto st88
tr323:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st88
tr352:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st88
tr383:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st88
tr416:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st88
tr446:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st88
tr480:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st88
tr511:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st88
tr543:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st88
tr576:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st88
tr607:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st88
tr637:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st88
tr659:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st88
tr690:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st88
tr719:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st88
tr739:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st88
tr770:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st88
tr800:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st88
tr830:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st88
tr860:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st88
tr890:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st88
tr921:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st88
tr951:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st88
tr981:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st88
tr1011:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st88
tr1042:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st88
tr1067:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st88
tr1097:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st88
tr1127:
//line lex.rl:43

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
		
	goto st88
tr1155:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st88
tr1185:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st88
tr1215:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st88
tr1245:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st88
tr1276:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st88
tr1306:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line lex.go:5838
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto st89
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto tr372
		case 33:
			goto tr373
		case 34:
			goto tr374
		case 35:
			goto tr375
		case 36:
			goto tr376
		case 38:
			goto tr378
		case 39:
			goto tr379
		case 43:
			goto tr380
		case 45:
			goto tr381
		case 46:
			goto tr382
		case 47:
			goto tr383
		case 48:
			goto tr384
		case 58:
			goto tr386
		case 60:
			goto tr387
		case 62:
			goto tr388
		case 63:
			goto tr389
		case 64:
			goto tr390
		case 66:
			goto tr392
		case 88:
			goto tr393
		case 95:
			goto tr394
		case 98:
			goto tr396
		case 101:
			goto tr397
		case 120:
			goto tr398
		case 124:
			goto tr399
		case 126:
			goto tr400
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
						goto tr377
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr372
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr391
					}
				case data[p] >= 59:
					goto tr377
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr377
					}
				case data[p] >= 97:
					goto tr395
				}
			default:
				goto tr377
			}
		default:
			goto tr385
		}
		goto tr401
tr164:
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr194:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr225:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr260:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr290:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr353:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr384:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr417:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr447:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr481:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr512:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr544:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr577:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr608:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr660:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr691:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr740:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr771:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr801:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr831:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr861:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr891:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr922:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr952:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr982:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1012:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1068:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1098:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1156:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1186:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1216:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1246:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1277:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1307:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line lex.go:6512
		switch data[p] {
		case 32:
			goto tr312
		case 33:
			goto tr313
		case 34:
			goto tr314
		case 35:
			goto tr315
		case 36:
			goto tr316
		case 38:
			goto tr318
		case 39:
			goto tr319
		case 43:
			goto tr320
		case 45:
			goto tr321
		case 46:
			goto tr402
		case 47:
			goto tr323
		case 58:
			goto tr324
		case 60:
			goto tr325
		case 62:
			goto tr326
		case 63:
			goto tr327
		case 64:
			goto tr328
		case 66:
			goto tr330
		case 69:
			goto tr331
		case 88:
			goto st70
		case 95:
			goto tr333
		case 98:
			goto tr335
		case 101:
			goto tr331
		case 120:
			goto st70
		case 124:
			goto tr337
		case 126:
			goto tr338
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
						goto tr317
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr312
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr329
					}
				case data[p] >= 59:
					goto tr317
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr317
					}
				case data[p] >= 97:
					goto tr334
				}
			default:
				goto tr317
			}
		default:
			goto tr309
		}
		goto tr339
tr165:
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr195:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr226:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr261:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr291:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr309:
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr354:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr385:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr418:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr448:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr482:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr513:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr545:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr578:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr609:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr661:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr692:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr741:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr772:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr802:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr832:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr862:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr892:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr923:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr953:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr983:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1013:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1069:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1099:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1157:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1187:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1217:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1247:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1278:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1308:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:64

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line lex.go:7089
		switch data[p] {
		case 32:
			goto tr312
		case 33:
			goto tr313
		case 34:
			goto tr314
		case 35:
			goto tr315
		case 36:
			goto tr316
		case 38:
			goto tr318
		case 39:
			goto tr319
		case 43:
			goto tr320
		case 45:
			goto tr321
		case 46:
			goto tr402
		case 47:
			goto tr323
		case 58:
			goto tr324
		case 60:
			goto tr325
		case 62:
			goto tr326
		case 63:
			goto tr327
		case 64:
			goto tr328
		case 66:
			goto tr330
		case 69:
			goto tr331
		case 88:
			goto tr332
		case 95:
			goto tr333
		case 98:
			goto tr335
		case 101:
			goto tr331
		case 120:
			goto tr336
		case 124:
			goto tr337
		case 126:
			goto tr338
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
						goto tr317
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr312
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr329
					}
				case data[p] >= 59:
					goto tr317
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr317
					}
				case data[p] >= 97:
					goto tr334
				}
			default:
				goto tr317
			}
		default:
			goto tr309
		}
		goto tr339
tr196:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st92
tr227:
//line lex.rl:130

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
		
	goto st92
tr262:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st92
tr292:
//line lex.rl:164

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
		
	goto st92
tr324:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st92
tr355:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st92
tr386:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st92
tr449:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st92
tr483:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st92
tr514:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st92
tr546:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st92
tr579:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st92
tr610:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st92
tr638:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st92
tr662:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st92
tr693:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st92
tr720:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st92
tr742:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st92
tr773:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st92
tr803:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st92
tr833:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st92
tr863:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st92
tr893:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st92
tr924:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st92
tr954:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st92
tr984:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st92
tr1014:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st92
tr1043:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st92
tr1070:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st92
tr1100:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st92
tr1128:
//line lex.rl:43

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
		
	goto st92
tr1158:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st92
tr1188:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st92
tr1218:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st92
tr1248:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st92
tr1279:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st92
tr1309:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line lex.go:7472
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto st93
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto tr405
		case 33:
			goto tr406
		case 34:
			goto tr407
		case 35:
			goto tr408
		case 36:
			goto tr409
		case 38:
			goto tr411
		case 39:
			goto tr412
		case 43:
			goto tr413
		case 45:
			goto tr414
		case 46:
			goto tr415
		case 47:
			goto tr416
		case 48:
			goto tr417
		case 58:
			goto st94
		case 60:
			goto tr420
		case 62:
			goto tr421
		case 63:
			goto tr422
		case 64:
			goto tr423
		case 66:
			goto tr425
		case 88:
			goto tr426
		case 95:
			goto tr427
		case 98:
			goto tr429
		case 101:
			goto tr430
		case 120:
			goto tr431
		case 124:
			goto tr432
		case 126:
			goto tr433
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
						goto tr410
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr405
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr424
					}
				case data[p] >= 59:
					goto tr410
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr410
					}
				case data[p] >= 97:
					goto tr428
				}
			default:
				goto tr410
			}
		default:
			goto tr418
		}
		goto tr434
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 32:
			goto tr435
		case 33:
			goto tr436
		case 34:
			goto tr437
		case 35:
			goto tr438
		case 36:
			goto tr439
		case 38:
			goto tr441
		case 39:
			goto tr442
		case 43:
			goto tr443
		case 45:
			goto tr444
		case 46:
			goto tr445
		case 47:
			goto tr446
		case 48:
			goto tr447
		case 58:
			goto tr449
		case 60:
			goto tr450
		case 62:
			goto tr451
		case 63:
			goto tr452
		case 64:
			goto tr453
		case 66:
			goto tr455
		case 88:
			goto tr456
		case 95:
			goto tr457
		case 98:
			goto tr459
		case 101:
			goto tr460
		case 120:
			goto tr461
		case 124:
			goto tr462
		case 126:
			goto tr463
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
						goto tr440
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr435
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr454
					}
				case data[p] >= 59:
					goto tr440
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr440
					}
				case data[p] >= 97:
					goto tr458
				}
			default:
				goto tr440
			}
		default:
			goto tr448
		}
		goto tr464
tr197:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st95
tr228:
//line lex.rl:130

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
		
	goto st95
tr263:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st95
tr293:
//line lex.rl:164

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
		
	goto st95
tr325:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st95
tr356:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st95
tr387:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st95
tr420:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st95
tr450:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st95
tr484:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st95
tr515:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st95
tr547:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st95
tr580:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st95
tr611:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st95
tr639:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st95
tr663:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st95
tr694:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st95
tr721:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st95
tr743:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st95
tr774:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st95
tr804:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st95
tr834:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st95
tr864:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st95
tr894:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st95
tr925:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st95
tr955:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st95
tr985:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st95
tr1015:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st95
tr1044:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st95
tr1071:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st95
tr1101:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st95
tr1129:
//line lex.rl:43

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
		
	goto st95
tr1159:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st95
tr1189:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st95
tr1219:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st95
tr1249:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st95
tr1280:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st95
tr1310:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line lex.go:8069
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 59:
			goto tr187
		case 60:
			goto st96
		case 61:
			goto st122
		case 62:
			goto st123
		case 63:
			goto tr200
		case 64:
			goto st124
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr202
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto tr469
		case 33:
			goto tr470
		case 34:
			goto tr471
		case 35:
			goto tr472
		case 36:
			goto tr473
		case 38:
			goto tr475
		case 39:
			goto tr476
		case 43:
			goto tr477
		case 45:
			goto tr478
		case 46:
			goto tr479
		case 47:
			goto tr480
		case 48:
			goto tr481
		case 58:
			goto tr483
		case 59:
			goto tr474
		case 60:
			goto tr484
		case 61:
			goto st97
		case 62:
			goto tr486
		case 63:
			goto tr487
		case 64:
			goto tr488
		case 66:
			goto tr490
		case 88:
			goto tr491
		case 95:
			goto tr492
		case 98:
			goto tr494
		case 101:
			goto tr495
		case 120:
			goto tr496
		case 124:
			goto tr497
		case 126:
			goto tr498
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
						goto tr474
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr469
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr489
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr474
					}
				case data[p] >= 97:
					goto tr493
				}
			default:
				goto tr474
			}
		default:
			goto tr482
		}
		goto tr499
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto tr500
		case 33:
			goto tr501
		case 34:
			goto tr502
		case 35:
			goto tr503
		case 36:
			goto tr504
		case 38:
			goto tr506
		case 39:
			goto tr507
		case 43:
			goto tr508
		case 45:
			goto tr509
		case 46:
			goto tr510
		case 47:
			goto tr511
		case 48:
			goto tr512
		case 58:
			goto tr514
		case 60:
			goto tr515
		case 62:
			goto tr516
		case 63:
			goto tr517
		case 64:
			goto tr518
		case 66:
			goto tr520
		case 88:
			goto tr521
		case 95:
			goto tr522
		case 98:
			goto tr524
		case 101:
			goto tr525
		case 120:
			goto tr526
		case 124:
			goto tr527
		case 126:
			goto tr528
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
						goto tr505
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr500
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr519
					}
				case data[p] >= 59:
					goto tr505
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr505
					}
				case data[p] >= 97:
					goto tr523
				}
			default:
				goto tr505
			}
		default:
			goto tr513
		}
		goto tr529
tr199:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st98
tr229:
//line lex.rl:130

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
		
	goto st98
tr264:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st98
tr294:
//line lex.rl:164

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
		
	goto st98
tr326:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st98
tr357:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st98
tr388:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st98
tr421:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st98
tr451:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st98
tr486:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st98
tr516:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st98
tr548:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st98
tr581:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st98
tr612:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st98
tr640:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st98
tr664:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st98
tr695:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st98
tr722:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st98
tr744:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st98
tr775:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st98
tr805:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st98
tr835:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st98
tr865:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st98
tr896:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st98
tr926:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st98
tr956:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st98
tr986:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st98
tr1016:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st98
tr1045:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st98
tr1102:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st98
tr1130:
//line lex.rl:43

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
		
	goto st98
tr1160:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st98
tr1220:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st98
tr1250:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st98
tr1281:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st98
tr1311:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line lex.go:8656
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 59:
			goto tr187
		case 60:
			goto tr197
		case 61:
			goto st99
		case 62:
			goto st120
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr202
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto tr532
		case 33:
			goto tr533
		case 34:
			goto tr534
		case 35:
			goto tr535
		case 36:
			goto tr536
		case 38:
			goto tr538
		case 39:
			goto tr539
		case 43:
			goto tr540
		case 45:
			goto tr541
		case 46:
			goto tr542
		case 47:
			goto tr543
		case 48:
			goto tr544
		case 58:
			goto tr546
		case 60:
			goto tr547
		case 62:
			goto tr548
		case 63:
			goto tr549
		case 64:
			goto tr550
		case 66:
			goto tr552
		case 88:
			goto tr553
		case 95:
			goto tr554
		case 98:
			goto tr556
		case 101:
			goto tr557
		case 120:
			goto tr558
		case 124:
			goto tr559
		case 126:
			goto tr560
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
						goto tr537
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr532
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr551
					}
				case data[p] >= 59:
					goto tr537
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr537
					}
				case data[p] >= 97:
					goto tr555
				}
			default:
				goto tr537
			}
		default:
			goto tr545
		}
		goto tr561
tr200:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st100
tr230:
//line lex.rl:130

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
		
	goto st100
tr265:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st100
tr295:
//line lex.rl:164

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
		
	goto st100
tr327:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st100
tr358:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st100
tr389:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st100
tr422:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st100
tr452:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st100
tr487:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st100
tr517:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st100
tr549:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st100
tr582:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st100
tr613:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st100
tr641:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st100
tr665:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st100
tr696:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st100
tr723:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st100
tr745:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st100
tr776:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st100
tr806:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st100
tr836:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st100
tr866:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st100
tr897:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st100
tr927:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st100
tr957:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st100
tr987:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st100
tr1017:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st100
tr1046:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st100
tr1073:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st100
tr1103:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st100
tr1131:
//line lex.rl:43

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
		
	goto st100
tr1161:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st100
tr1191:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st100
tr1221:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st100
tr1251:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st100
tr1282:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st100
tr1312:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line lex.go:9147
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto st101
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto st118
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto st119
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto tr565
		case 33:
			goto tr566
		case 34:
			goto tr567
		case 35:
			goto tr568
		case 36:
			goto tr569
		case 38:
			goto tr571
		case 39:
			goto tr572
		case 43:
			goto tr573
		case 45:
			goto tr574
		case 46:
			goto tr575
		case 47:
			goto tr576
		case 48:
			goto tr577
		case 58:
			goto tr579
		case 60:
			goto tr580
		case 62:
			goto tr581
		case 63:
			goto tr582
		case 64:
			goto tr583
		case 66:
			goto tr585
		case 88:
			goto tr586
		case 95:
			goto tr587
		case 98:
			goto tr589
		case 101:
			goto tr590
		case 120:
			goto tr591
		case 124:
			goto tr592
		case 126:
			goto tr593
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
						goto tr570
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr565
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr584
					}
				case data[p] >= 59:
					goto tr570
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr570
					}
				case data[p] >= 97:
					goto tr588
				}
			default:
				goto tr570
			}
		default:
			goto tr578
		}
		goto tr594
tr201:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st102
tr231:
//line lex.rl:130

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
		
	goto st102
tr266:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st102
tr296:
//line lex.rl:164

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
		
	goto st102
tr328:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st102
tr359:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st102
tr390:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st102
tr423:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st102
tr453:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st102
tr488:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st102
tr518:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st102
tr550:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st102
tr583:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st102
tr614:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st102
tr642:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st102
tr666:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st102
tr697:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st102
tr724:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st102
tr746:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st102
tr777:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st102
tr807:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st102
tr837:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st102
tr867:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st102
tr898:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st102
tr928:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st102
tr958:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st102
tr988:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st102
tr1018:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st102
tr1047:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st102
tr1074:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st102
tr1104:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st102
tr1132:
//line lex.rl:43

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
		
	goto st102
tr1162:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st102
tr1192:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st102
tr1222:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st102
tr1252:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st102
tr1283:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st102
tr1313:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line lex.go:9639
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto st103
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr596
		case 33:
			goto tr597
		case 34:
			goto tr598
		case 35:
			goto tr599
		case 36:
			goto tr600
		case 38:
			goto tr602
		case 39:
			goto tr603
		case 43:
			goto tr604
		case 45:
			goto tr605
		case 46:
			goto tr606
		case 47:
			goto tr607
		case 48:
			goto tr608
		case 58:
			goto tr610
		case 60:
			goto tr611
		case 62:
			goto tr612
		case 63:
			goto tr613
		case 64:
			goto tr614
		case 66:
			goto tr616
		case 88:
			goto tr617
		case 95:
			goto tr618
		case 98:
			goto tr620
		case 101:
			goto tr621
		case 120:
			goto tr622
		case 124:
			goto tr623
		case 126:
			goto tr624
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
						goto tr601
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr596
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr615
					}
				case data[p] >= 59:
					goto tr601
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr601
					}
				case data[p] >= 97:
					goto tr619
				}
			default:
				goto tr601
			}
		default:
			goto tr609
		}
		goto tr625
tr646:
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr175:
//line lex.rl:40
 mark = p 
	goto st104
tr171:
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr181:
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr206:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st104
tr202:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr212:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr236:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st104
tr232:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr242:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr271:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st104
tr267:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr277:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr301:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st104
tr297:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr307:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr334:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st104
tr329:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr339:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr364:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st104
tr360:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr370:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr395:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st104
tr391:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr401:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr428:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st104
tr424:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr434:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr458:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st104
tr454:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr464:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr493:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st104
tr489:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr499:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr523:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr519:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr529:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr555:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr551:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr561:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr588:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st104
tr584:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr594:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr619:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st104
tr615:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr625:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr643:
//line lex.rl:113
 isUpper = true 
	goto st104
tr671:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st104
tr667:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr677:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr702:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st104
tr698:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr708:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr751:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st104
tr747:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr757:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr782:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st104
tr778:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr788:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr812:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st104
tr808:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr818:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr842:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st104
tr838:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr848:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr872:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st104
tr868:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr878:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr903:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st104
tr899:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr909:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr933:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr929:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr939:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr963:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr959:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr969:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr993:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr989:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr999:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1023:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st104
tr1019:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1029:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1051:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st104
tr1048:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1055:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1079:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st104
tr1075:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1085:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1109:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st104
tr1105:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1115:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1137:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st104
tr1133:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1143:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1167:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st104
tr1163:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1173:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1197:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st104
tr1193:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1203:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1227:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st104
tr1223:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1233:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1257:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st104
tr1253:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1263:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1288:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st104
tr1284:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1294:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
tr1318:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st104
tr1314:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st104
tr1324:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:110
 isNotASCII = true 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line lex.go:10903
		switch data[p] {
		case 32:
			goto tr626
		case 33:
			goto tr627
		case 34:
			goto tr628
		case 35:
			goto tr629
		case 36:
			goto st104
		case 38:
			goto tr632
		case 39:
			goto tr633
		case 43:
			goto tr634
		case 45:
			goto tr635
		case 46:
			goto tr636
		case 47:
			goto tr637
		case 58:
			goto tr638
		case 60:
			goto tr639
		case 62:
			goto tr640
		case 63:
			goto tr641
		case 64:
			goto tr642
		case 95:
			goto st104
		case 124:
			goto tr644
		case 126:
			goto tr645
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
						goto tr631
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr626
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr631
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr631
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr631
			}
		default:
			goto st104
		}
		goto tr646
tr210:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st105
tr240:
//line lex.rl:130

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
		
	goto st105
tr275:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st105
tr305:
//line lex.rl:164

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
		
	goto st105
tr337:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st105
tr368:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st105
tr399:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st105
tr432:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st105
tr462:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st105
tr497:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st105
tr527:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st105
tr559:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st105
tr592:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st105
tr623:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st105
tr644:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st105
tr675:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st105
tr706:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st105
tr725:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st105
tr755:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st105
tr786:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st105
tr816:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st105
tr846:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st105
tr876:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st105
tr907:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st105
tr937:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st105
tr967:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st105
tr997:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st105
tr1027:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st105
tr1053:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st105
tr1083:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st105
tr1113:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st105
tr1141:
//line lex.rl:43

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
		
	goto st105
tr1171:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st105
tr1201:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st105
tr1231:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st105
tr1261:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st105
tr1292:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st105
tr1322:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line lex.go:11278
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto st117
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
tr172:
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr203:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr233:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr268:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr298:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr330:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr361:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr392:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr425:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr455:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr490:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr520:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr552:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr585:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr616:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr668:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr699:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr748:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr779:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr809:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr839:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr869:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr900:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr930:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr960:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr990:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1020:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1076:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1106:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1134:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1164:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1194:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1224:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1254:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1285:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
tr1315:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line lex.go:11746
		switch data[p] {
		case 32:
			goto tr626
		case 33:
			goto tr627
		case 34:
			goto tr628
		case 35:
			goto tr629
		case 36:
			goto st104
		case 38:
			goto tr632
		case 39:
			goto st6
		case 43:
			goto tr634
		case 45:
			goto tr635
		case 46:
			goto tr636
		case 47:
			goto tr637
		case 58:
			goto tr638
		case 60:
			goto tr639
		case 62:
			goto tr640
		case 63:
			goto tr641
		case 64:
			goto tr642
		case 95:
			goto st104
		case 124:
			goto tr644
		case 126:
			goto tr645
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
						goto tr631
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr626
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr631
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr631
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr631
			}
		default:
			goto st104
		}
		goto tr646
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 39 {
			goto st107
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st6
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto tr648
		case 33:
			goto tr649
		case 34:
			goto tr650
		case 35:
			goto tr651
		case 36:
			goto tr652
		case 38:
			goto tr654
		case 39:
			goto tr655
		case 43:
			goto tr656
		case 45:
			goto tr657
		case 46:
			goto tr658
		case 47:
			goto tr659
		case 48:
			goto tr660
		case 58:
			goto tr662
		case 60:
			goto tr663
		case 62:
			goto tr664
		case 63:
			goto tr665
		case 64:
			goto tr666
		case 66:
			goto tr668
		case 88:
			goto tr669
		case 95:
			goto tr670
		case 98:
			goto tr672
		case 101:
			goto tr673
		case 120:
			goto tr674
		case 124:
			goto tr675
		case 126:
			goto tr676
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
						goto tr653
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr648
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr667
					}
				case data[p] >= 59:
					goto tr653
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr653
					}
				case data[p] >= 97:
					goto tr671
				}
			default:
				goto tr653
			}
		default:
			goto tr661
		}
		goto tr677
tr178:
//line lex.rl:40
 mark = p 
	goto st108
tr173:
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr209:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st108
tr204:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr239:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st108
tr234:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr274:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st108
tr269:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr304:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st108
tr299:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr336:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st108
tr332:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr367:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st108
tr362:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr398:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st108
tr393:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr431:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st108
tr426:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr461:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st108
tr456:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr496:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st108
tr491:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr526:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr521:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr558:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr553:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr591:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st108
tr586:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr622:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st108
tr617:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr674:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st108
tr669:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr705:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st108
tr700:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr754:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st108
tr749:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr785:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st108
tr780:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr815:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st108
tr810:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr845:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st108
tr840:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr875:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st108
tr870:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr906:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st108
tr901:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr936:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr931:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr966:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr961:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr996:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr991:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1026:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st108
tr1021:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1052:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st108
tr1049:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1082:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st108
tr1077:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1112:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st108
tr1107:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1140:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st108
tr1135:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1170:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st108
tr1165:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1200:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st108
tr1195:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1230:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st108
tr1225:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1260:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st108
tr1255:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1291:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st108
tr1286:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
tr1321:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st108
tr1316:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
//line lex.rl:113
 isUpper = true 
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line lex.go:12629
		switch data[p] {
		case 32:
			goto tr626
		case 33:
			goto tr627
		case 34:
			goto tr628
		case 35:
			goto tr629
		case 36:
			goto st104
		case 38:
			goto tr632
		case 39:
			goto tr678
		case 43:
			goto tr634
		case 45:
			goto tr635
		case 46:
			goto tr636
		case 47:
			goto tr637
		case 58:
			goto tr638
		case 60:
			goto tr639
		case 62:
			goto tr640
		case 63:
			goto tr641
		case 64:
			goto tr642
		case 95:
			goto st104
		case 124:
			goto tr644
		case 126:
			goto tr645
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
						goto tr631
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr626
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr631
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr631
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr631
			}
		default:
			goto st104
		}
		goto tr646
tr678:
//line lex.rl:250
 buf = new(bytes.Buffer) 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line lex.go:12727
		if data[p] == 39 {
			goto st109
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
tr147:
//line lex.rl:256
 buf.WriteByte(ch) 
	goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line lex.go:12753
		switch data[p] {
		case 32:
			goto tr679
		case 33:
			goto tr680
		case 34:
			goto tr681
		case 35:
			goto tr682
		case 36:
			goto tr683
		case 38:
			goto tr685
		case 39:
			goto tr686
		case 43:
			goto tr687
		case 45:
			goto tr688
		case 46:
			goto tr689
		case 47:
			goto tr690
		case 48:
			goto tr691
		case 58:
			goto tr693
		case 60:
			goto tr694
		case 62:
			goto tr695
		case 63:
			goto tr696
		case 64:
			goto tr697
		case 66:
			goto tr699
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
						goto tr684
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr679
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr698
					}
				case data[p] >= 59:
					goto tr684
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr684
					}
				case data[p] >= 97:
					goto tr702
				}
			default:
				goto tr684
			}
		default:
			goto tr692
		}
		goto tr708
tr174:
//line lex.rl:40
 mark = p 
	goto st110
tr205:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st110
tr235:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st110
tr270:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st110
tr300:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st110
tr333:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st110
tr363:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st110
tr394:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st110
tr427:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st110
tr457:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st110
tr492:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st110
tr522:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr554:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr587:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st110
tr618:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st110
tr670:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st110
tr701:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st110
tr750:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st110
tr781:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st110
tr811:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st110
tr841:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st110
tr871:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st110
tr902:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st110
tr932:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr962:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr992:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr1022:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st110
tr1050:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
//line lex.rl:40
 mark = p 
	goto st110
tr1078:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st110
tr1108:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st110
tr1136:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st110
tr1166:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st110
tr1196:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st110
tr1226:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st110
tr1256:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st110
tr1287:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st110
tr1317:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line lex.go:13158
		switch data[p] {
		case 32:
			goto tr709
		case 33:
			goto tr710
		case 34:
			goto tr711
		case 35:
			goto tr712
		case 36:
			goto st104
		case 38:
			goto tr714
		case 39:
			goto tr715
		case 43:
			goto tr716
		case 45:
			goto tr717
		case 46:
			goto tr718
		case 47:
			goto tr719
		case 58:
			goto tr720
		case 60:
			goto tr721
		case 62:
			goto tr722
		case 63:
			goto tr723
		case 64:
			goto tr724
		case 95:
			goto st104
		case 124:
			goto tr725
		case 126:
			goto tr726
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
						goto tr713
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr709
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr713
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr713
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr713
			}
		default:
			goto st104
		}
		goto tr646
tr245:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st111
tr241:
//line lex.rl:130

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
		
	goto st111
tr276:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
	goto st111
tr306:
//line lex.rl:164

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
		
	goto st111
tr338:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
	goto st111
tr369:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
	goto st111
tr400:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
	goto st111
tr433:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
	goto st111
tr463:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
	goto st111
tr498:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
	goto st111
tr528:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
	goto st111
tr560:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
	goto st111
tr593:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
	goto st111
tr624:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
	goto st111
tr645:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
	goto st111
tr676:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
	goto st111
tr707:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
	goto st111
tr726:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
	goto st111
tr756:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
	goto st111
tr787:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
	goto st111
tr817:
//line lex.rl:307
 emitToken(lex.CONCAT) 
	goto st111
tr847:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
	goto st111
tr877:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
	goto st111
tr908:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
	goto st111
tr938:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
	goto st111
tr968:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
	goto st111
tr998:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
	goto st111
tr1028:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
	goto st111
tr1054:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
	goto st111
tr1084:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
	goto st111
tr1114:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
	goto st111
tr1142:
//line lex.rl:43

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
		
	goto st111
tr1172:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
	goto st111
tr1202:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
	goto st111
tr1232:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
	goto st111
tr1262:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
	goto st111
tr1293:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
	goto st111
tr1323:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line lex.go:13533
		switch data[p] {
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 34:
			goto tr184
		case 35:
			goto tr185
		case 36:
			goto tr186
		case 38:
			goto tr188
		case 39:
			goto tr189
		case 42:
			goto st112
		case 43:
			goto tr190
		case 45:
			goto tr191
		case 46:
			goto tr192
		case 47:
			goto tr193
		case 48:
			goto tr194
		case 58:
			goto tr196
		case 60:
			goto tr197
		case 62:
			goto tr199
		case 63:
			goto tr200
		case 64:
			goto tr201
		case 66:
			goto tr203
		case 88:
			goto tr204
		case 95:
			goto tr205
		case 98:
			goto tr207
		case 101:
			goto tr208
		case 120:
			goto tr209
		case 124:
			goto tr210
		case 126:
			goto tr245
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
						goto tr187
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr182
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr202
					}
				case data[p] >= 59:
					goto tr187
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr187
					}
				case data[p] >= 97:
					goto tr206
				}
			default:
				goto tr187
			}
		default:
			goto tr195
		}
		goto tr212
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 32:
			goto tr728
		case 33:
			goto tr729
		case 34:
			goto tr730
		case 35:
			goto tr731
		case 36:
			goto tr732
		case 38:
			goto tr734
		case 39:
			goto tr735
		case 43:
			goto tr736
		case 45:
			goto tr737
		case 46:
			goto tr738
		case 47:
			goto tr739
		case 48:
			goto tr740
		case 58:
			goto tr742
		case 60:
			goto tr743
		case 62:
			goto tr744
		case 63:
			goto tr745
		case 64:
			goto tr746
		case 66:
			goto tr748
		case 88:
			goto tr749
		case 95:
			goto tr750
		case 98:
			goto tr752
		case 101:
			goto tr753
		case 120:
			goto tr754
		case 124:
			goto tr755
		case 126:
			goto tr756
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
						goto tr733
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr728
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr747
					}
				case data[p] >= 59:
					goto tr733
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr733
					}
				case data[p] >= 97:
					goto tr751
				}
			default:
				goto tr733
			}
		default:
			goto tr741
		}
		goto tr757
tr176:
//line lex.rl:40
 mark = p 
	goto st113
tr207:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st113
tr237:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st113
tr272:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st113
tr302:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st113
tr335:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st113
tr365:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st113
tr396:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st113
tr429:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st113
tr459:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st113
tr494:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st113
tr524:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr556:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr589:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st113
tr620:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st113
tr672:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st113
tr703:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st113
tr752:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st113
tr783:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st113
tr813:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st113
tr843:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st113
tr873:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st113
tr904:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st113
tr934:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr964:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr994:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr1024:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st113
tr1080:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st113
tr1110:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st113
tr1138:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st113
tr1168:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st113
tr1198:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st113
tr1228:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st113
tr1258:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st113
tr1289:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st113
tr1319:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line lex.go:14036
		switch data[p] {
		case 32:
			goto tr626
		case 33:
			goto tr627
		case 34:
			goto tr628
		case 35:
			goto tr629
		case 36:
			goto st104
		case 38:
			goto tr632
		case 39:
			goto st8
		case 43:
			goto tr634
		case 45:
			goto tr635
		case 46:
			goto tr636
		case 47:
			goto tr637
		case 58:
			goto tr638
		case 60:
			goto tr639
		case 62:
			goto tr640
		case 63:
			goto tr641
		case 64:
			goto tr642
		case 95:
			goto st104
		case 124:
			goto tr644
		case 126:
			goto tr645
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
						goto tr631
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr626
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr631
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr631
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr631
			}
		default:
			goto st104
		}
		goto tr646
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 39:
			goto tr19
		case 92:
			goto tr20
		}
		goto tr18
tr18:
//line lex.rl:222
 buf = new(bytes.Buffer) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr21:
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr88:
//line lex.rl:233
 buf.WriteByte(data[p]) 
	goto st9
tr102:
//line lex.rl:219
 buf.WriteByte(ch) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr113:
//line lex.rl:213
 buf.WriteRune(rn) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr118:
//line lex.rl:202
 buf.WriteByte(ch) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr121:
//line lex.rl:190
 buf.WriteByte('\a') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr124:
//line lex.rl:191
 buf.WriteByte('\b') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr127:
//line lex.rl:192
 buf.WriteByte('\f') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr130:
//line lex.rl:193
 buf.WriteByte('\n') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr133:
//line lex.rl:194
 buf.WriteByte('\r') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr136:
//line lex.rl:195
 buf.WriteByte('\t') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
tr143:
//line lex.rl:196
 buf.WriteByte('\v') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:14216
		switch data[p] {
		case 39:
			goto st114
		case 92:
			goto st39
		}
		goto tr21
tr19:
//line lex.rl:222
 buf = new(bytes.Buffer) 
	goto st114
tr103:
//line lex.rl:219
 buf.WriteByte(ch) 
	goto st114
tr114:
//line lex.rl:213
 buf.WriteRune(rn) 
	goto st114
tr119:
//line lex.rl:202
 buf.WriteByte(ch) 
	goto st114
tr122:
//line lex.rl:190
 buf.WriteByte('\a') 
	goto st114
tr125:
//line lex.rl:191
 buf.WriteByte('\b') 
	goto st114
tr128:
//line lex.rl:192
 buf.WriteByte('\f') 
	goto st114
tr131:
//line lex.rl:193
 buf.WriteByte('\n') 
	goto st114
tr134:
//line lex.rl:194
 buf.WriteByte('\r') 
	goto st114
tr137:
//line lex.rl:195
 buf.WriteByte('\t') 
	goto st114
tr144:
//line lex.rl:196
 buf.WriteByte('\v') 
	goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line lex.go:14273
		switch data[p] {
		case 32:
			goto tr679
		case 33:
			goto tr680
		case 34:
			goto tr681
		case 35:
			goto tr682
		case 36:
			goto tr683
		case 38:
			goto tr685
		case 39:
			goto tr21
		case 43:
			goto tr687
		case 45:
			goto tr688
		case 46:
			goto tr689
		case 47:
			goto tr690
		case 48:
			goto tr691
		case 58:
			goto tr693
		case 60:
			goto tr694
		case 62:
			goto tr695
		case 63:
			goto tr696
		case 64:
			goto tr697
		case 66:
			goto tr699
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
						goto tr684
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr679
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr698
					}
				case data[p] >= 59:
					goto tr684
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr684
					}
				case data[p] >= 97:
					goto tr702
				}
			default:
				goto tr684
			}
		default:
			goto tr692
		}
		goto tr708
tr177:
//line lex.rl:40
 mark = p 
	goto st115
tr208:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.rl:40
 mark = p 
	goto st115
tr238:
//line lex.rl:130

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
		
//line lex.rl:40
 mark = p 
	goto st115
tr273:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st115
tr303:
//line lex.rl:164

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
		
//line lex.rl:40
 mark = p 
	goto st115
tr1030:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
//line lex.rl:40
 mark = p 
	goto st115
tr366:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
//line lex.rl:40
 mark = p 
	goto st115
tr397:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
//line lex.rl:40
 mark = p 
	goto st115
tr430:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
//line lex.rl:40
 mark = p 
	goto st115
tr460:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
//line lex.rl:40
 mark = p 
	goto st115
tr495:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
//line lex.rl:40
 mark = p 
	goto st115
tr525:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr557:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr590:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st115
tr621:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
//line lex.rl:40
 mark = p 
	goto st115
tr673:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
//line lex.rl:40
 mark = p 
	goto st115
tr704:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st115
tr753:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st115
tr784:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
//line lex.rl:40
 mark = p 
	goto st115
tr814:
//line lex.rl:307
 emitToken(lex.CONCAT) 
//line lex.rl:40
 mark = p 
	goto st115
tr844:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
//line lex.rl:40
 mark = p 
	goto st115
tr874:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
//line lex.rl:40
 mark = p 
	goto st115
tr905:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
//line lex.rl:40
 mark = p 
	goto st115
tr935:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr965:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr995:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr1025:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
//line lex.rl:40
 mark = p 
	goto st115
tr1081:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
//line lex.rl:40
 mark = p 
	goto st115
tr1111:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
//line lex.rl:40
 mark = p 
	goto st115
tr1139:
//line lex.rl:43

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
		
//line lex.rl:40
 mark = p 
	goto st115
tr1169:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
//line lex.rl:40
 mark = p 
	goto st115
tr1199:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
//line lex.rl:40
 mark = p 
	goto st115
tr1229:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
//line lex.rl:40
 mark = p 
	goto st115
tr1259:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
//line lex.rl:40
 mark = p 
	goto st115
tr1290:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
//line lex.rl:40
 mark = p 
	goto st115
tr1320:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
//line lex.rl:40
 mark = p 
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line lex.go:14669
		switch data[p] {
		case 32:
			goto tr626
		case 33:
			goto tr627
		case 34:
			goto tr628
		case 35:
			goto tr629
		case 36:
			goto st104
		case 38:
			goto tr632
		case 39:
			goto st10
		case 43:
			goto tr634
		case 45:
			goto tr635
		case 46:
			goto tr636
		case 47:
			goto tr637
		case 58:
			goto tr638
		case 60:
			goto tr639
		case 62:
			goto tr640
		case 63:
			goto tr641
		case 64:
			goto tr642
		case 95:
			goto st104
		case 124:
			goto tr644
		case 126:
			goto tr645
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
						goto tr631
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr626
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr643
					}
				case data[p] >= 59:
					goto tr631
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr631
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr631
			}
		default:
			goto st104
		}
		goto tr646
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 39:
			goto tr25
		case 92:
			goto tr26
		}
		goto tr24
tr24:
//line lex.rl:222
 buf = new(bytes.Buffer) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr27:
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr30:
//line lex.rl:233
 buf.WriteByte(data[p]) 
	goto st11
tr44:
//line lex.rl:219
 buf.WriteByte(ch) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr55:
//line lex.rl:213
 buf.WriteRune(rn) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr60:
//line lex.rl:202
 buf.WriteByte(ch) 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr63:
//line lex.rl:190
 buf.WriteByte('\a') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr66:
//line lex.rl:191
 buf.WriteByte('\b') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr69:
//line lex.rl:192
 buf.WriteByte('\f') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr72:
//line lex.rl:193
 buf.WriteByte('\n') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr75:
//line lex.rl:194
 buf.WriteByte('\r') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr78:
//line lex.rl:195
 buf.WriteByte('\t') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
tr85:
//line lex.rl:196
 buf.WriteByte('\v') 
//line lex.rl:227
 buf.WriteByte(data[p]) 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line lex.go:14849
		switch data[p] {
		case 39:
			goto st116
		case 92:
			goto st12
		}
		goto tr27
tr25:
//line lex.rl:222
 buf = new(bytes.Buffer) 
	goto st116
tr45:
//line lex.rl:219
 buf.WriteByte(ch) 
	goto st116
tr56:
//line lex.rl:213
 buf.WriteRune(rn) 
	goto st116
tr61:
//line lex.rl:202
 buf.WriteByte(ch) 
	goto st116
tr64:
//line lex.rl:190
 buf.WriteByte('\a') 
	goto st116
tr67:
//line lex.rl:191
 buf.WriteByte('\b') 
	goto st116
tr70:
//line lex.rl:192
 buf.WriteByte('\f') 
	goto st116
tr73:
//line lex.rl:193
 buf.WriteByte('\n') 
	goto st116
tr76:
//line lex.rl:194
 buf.WriteByte('\r') 
	goto st116
tr79:
//line lex.rl:195
 buf.WriteByte('\t') 
	goto st116
tr86:
//line lex.rl:196
 buf.WriteByte('\v') 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line lex.go:14906
		switch data[p] {
		case 32:
			goto tr760
		case 33:
			goto tr761
		case 34:
			goto tr762
		case 35:
			goto tr763
		case 36:
			goto tr764
		case 38:
			goto tr766
		case 39:
			goto tr27
		case 43:
			goto tr767
		case 45:
			goto tr768
		case 46:
			goto tr769
		case 47:
			goto tr770
		case 48:
			goto tr771
		case 58:
			goto tr773
		case 60:
			goto tr774
		case 62:
			goto tr775
		case 63:
			goto tr776
		case 64:
			goto tr777
		case 66:
			goto tr779
		case 88:
			goto tr780
		case 95:
			goto tr781
		case 98:
			goto tr783
		case 101:
			goto tr784
		case 120:
			goto tr785
		case 124:
			goto tr786
		case 126:
			goto tr787
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
						goto tr765
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr760
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr778
					}
				case data[p] >= 59:
					goto tr765
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr765
					}
				case data[p] >= 97:
					goto tr782
				}
			default:
				goto tr765
			}
		default:
			goto tr772
		}
		goto tr788
tr26:
//line lex.rl:222
 buf = new(bytes.Buffer) 
	goto st12
tr46:
//line lex.rl:219
 buf.WriteByte(ch) 
	goto st12
tr57:
//line lex.rl:213
 buf.WriteRune(rn) 
	goto st12
tr62:
//line lex.rl:202
 buf.WriteByte(ch) 
	goto st12
tr65:
//line lex.rl:190
 buf.WriteByte('\a') 
	goto st12
tr68:
//line lex.rl:191
 buf.WriteByte('\b') 
	goto st12
tr71:
//line lex.rl:192
 buf.WriteByte('\f') 
	goto st12
tr74:
//line lex.rl:193
 buf.WriteByte('\n') 
	goto st12
tr77:
//line lex.rl:194
 buf.WriteByte('\r') 
	goto st12
tr80:
//line lex.rl:195
 buf.WriteByte('\t') 
	goto st12
tr87:
//line lex.rl:196
 buf.WriteByte('\v') 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:15056
		switch data[p] {
		case 85:
			goto tr32
		case 88:
			goto st25
		case 97:
			goto st28
		case 98:
			goto st29
		case 102:
			goto st30
		case 110:
			goto st31
		case 114:
			goto st32
		case 116:
			goto st33
		case 117:
			goto tr40
		case 118:
			goto st38
		case 120:
			goto st25
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr31
		}
		goto tr30
tr31:
//line lex.rl:217
 ch = 0 
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:15096
		if 48 <= data[p] && data[p] <= 55 {
			goto tr42
		}
		goto st0
tr42:
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line lex.go:15110
		if 48 <= data[p] && data[p] <= 55 {
			goto tr43
		}
		goto st0
tr43:
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:15124
		switch data[p] {
		case 39:
			goto tr45
		case 92:
			goto tr46
		}
		goto tr44
tr32:
//line lex.rl:212
 rn = 0 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:15141
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line lex.go:15164
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:15187
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:15210
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:15233
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:15256
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line lex.go:15279
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:15302
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st24
tr84:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line lex.go:15329
		switch data[p] {
		case 39:
			goto tr56
		case 92:
			goto tr57
		}
		goto tr55
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
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
//line lex.rl:200
 ch = 0 
//line lex.rl:201
 ch = (ch << 4) | unhex(data[p]) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:15366
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
//line lex.rl:201
 ch = (ch << 4) | unhex(data[p]) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line lex.go:15389
		switch data[p] {
		case 39:
			goto tr61
		case 92:
			goto tr62
		}
		goto tr60
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 39:
			goto tr64
		case 92:
			goto tr65
		}
		goto tr63
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 39:
			goto tr67
		case 92:
			goto tr68
		}
		goto tr66
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 39:
			goto tr70
		case 92:
			goto tr71
		}
		goto tr69
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 39:
			goto tr73
		case 92:
			goto tr74
		}
		goto tr72
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr76
		case 92:
			goto tr77
		}
		goto tr75
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 39:
			goto tr79
		case 92:
			goto tr80
		}
		goto tr78
tr40:
//line lex.rl:212
 rn = 0 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line lex.go:15478
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
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line lex.go:15501
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
tr82:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line lex.go:15524
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr83
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr83
			}
		default:
			goto tr83
		}
		goto st0
tr83:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line lex.go:15547
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
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 39:
			goto tr86
		case 92:
			goto tr87
		}
		goto tr85
tr20:
//line lex.rl:222
 buf = new(bytes.Buffer) 
	goto st39
tr104:
//line lex.rl:219
 buf.WriteByte(ch) 
	goto st39
tr115:
//line lex.rl:213
 buf.WriteRune(rn) 
	goto st39
tr120:
//line lex.rl:202
 buf.WriteByte(ch) 
	goto st39
tr123:
//line lex.rl:190
 buf.WriteByte('\a') 
	goto st39
tr126:
//line lex.rl:191
 buf.WriteByte('\b') 
	goto st39
tr129:
//line lex.rl:192
 buf.WriteByte('\f') 
	goto st39
tr132:
//line lex.rl:193
 buf.WriteByte('\n') 
	goto st39
tr135:
//line lex.rl:194
 buf.WriteByte('\r') 
	goto st39
tr138:
//line lex.rl:195
 buf.WriteByte('\t') 
	goto st39
tr145:
//line lex.rl:196
 buf.WriteByte('\v') 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:15622
		switch data[p] {
		case 85:
			goto tr90
		case 88:
			goto st52
		case 97:
			goto st55
		case 98:
			goto st56
		case 102:
			goto st57
		case 110:
			goto st58
		case 114:
			goto st59
		case 116:
			goto st60
		case 117:
			goto tr98
		case 118:
			goto st65
		case 120:
			goto st52
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr89
		}
		goto tr88
tr89:
//line lex.rl:217
 ch = 0 
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lex.go:15662
		if 48 <= data[p] && data[p] <= 55 {
			goto tr100
		}
		goto st0
tr100:
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:15676
		if 48 <= data[p] && data[p] <= 55 {
			goto tr101
		}
		goto st0
tr101:
//line lex.rl:218
 ch = (ch << 3) | data[p] - '0' 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line lex.go:15690
		switch data[p] {
		case 39:
			goto tr103
		case 92:
			goto tr104
		}
		goto tr102
tr90:
//line lex.rl:212
 rn = 0 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lex.go:15707
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lex.go:15730
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:15753
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lex.go:15776
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:15799
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lex.go:15822
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line lex.go:15845
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:15868
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
//line lex.rl:210
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st51
tr142:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line lex.go:15895
		switch data[p] {
		case 39:
			goto tr114
		case 92:
			goto tr115
		}
		goto tr113
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
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
//line lex.rl:200
 ch = 0 
//line lex.rl:201
 ch = (ch << 4) | unhex(data[p]) 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lex.go:15932
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr117
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr117
			}
		default:
			goto tr117
		}
		goto st0
tr117:
//line lex.rl:201
 ch = (ch << 4) | unhex(data[p]) 
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line lex.go:15955
		switch data[p] {
		case 39:
			goto tr119
		case 92:
			goto tr120
		}
		goto tr118
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto tr123
		}
		goto tr121
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 39:
			goto tr125
		case 92:
			goto tr126
		}
		goto tr124
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 39:
			goto tr128
		case 92:
			goto tr129
		}
		goto tr127
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 39:
			goto tr131
		case 92:
			goto tr132
		}
		goto tr130
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 39:
			goto tr134
		case 92:
			goto tr135
		}
		goto tr133
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 39:
			goto tr137
		case 92:
			goto tr138
		}
		goto tr136
tr98:
//line lex.rl:212
 rn = 0 
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line lex.go:16044
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
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line lex.go:16067
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
tr140:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line lex.go:16090
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr141
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr141
			}
		default:
			goto tr141
		}
		goto st0
tr141:
//line lex.rl:207
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line lex.go:16113
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
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 39:
			goto tr144
		case 92:
			goto tr145
		}
		goto tr143
tr17:
//line lex.rl:254
 ch = 0 
//line lex.rl:255
 ch = (ch << 4) | unhex(data[p]) 
	goto st66
tr148:
//line lex.rl:256
 buf.WriteByte(ch) 
//line lex.rl:254
 ch = 0 
//line lex.rl:255
 ch = (ch << 4) | unhex(data[p]) 
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lex.go:16158
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
tr146:
//line lex.rl:255
 ch = (ch << 4) | unhex(data[p]) 
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lex.go:16181
		if data[p] == 39 {
			goto tr147
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr148
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto tr789
		case 33:
			goto tr790
		case 34:
			goto tr791
		case 35:
			goto tr792
		case 36:
			goto tr793
		case 38:
			goto tr795
		case 39:
			goto tr796
		case 43:
			goto tr797
		case 45:
			goto tr798
		case 46:
			goto tr799
		case 47:
			goto tr800
		case 48:
			goto tr801
		case 58:
			goto tr803
		case 60:
			goto tr804
		case 62:
			goto tr805
		case 63:
			goto tr806
		case 64:
			goto tr807
		case 66:
			goto tr809
		case 88:
			goto tr810
		case 95:
			goto tr811
		case 98:
			goto tr813
		case 101:
			goto tr814
		case 120:
			goto tr815
		case 124:
			goto tr816
		case 126:
			goto tr817
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
						goto tr794
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr789
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr808
					}
				case data[p] >= 59:
					goto tr794
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr794
					}
				case data[p] >= 97:
					goto tr812
				}
			default:
				goto tr794
			}
		default:
			goto tr802
		}
		goto tr818
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 32:
			goto tr819
		case 33:
			goto tr820
		case 34:
			goto tr821
		case 35:
			goto tr822
		case 36:
			goto tr823
		case 38:
			goto tr825
		case 39:
			goto tr826
		case 43:
			goto tr827
		case 45:
			goto tr828
		case 46:
			goto tr829
		case 47:
			goto tr830
		case 48:
			goto tr831
		case 58:
			goto tr833
		case 60:
			goto tr834
		case 62:
			goto tr835
		case 63:
			goto tr836
		case 64:
			goto tr837
		case 66:
			goto tr839
		case 88:
			goto tr840
		case 95:
			goto tr841
		case 98:
			goto tr843
		case 101:
			goto tr844
		case 120:
			goto tr845
		case 124:
			goto tr846
		case 126:
			goto tr847
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
						goto tr824
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr819
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr838
					}
				case data[p] >= 59:
					goto tr824
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr824
					}
				case data[p] >= 97:
					goto tr842
				}
			default:
				goto tr824
			}
		default:
			goto tr832
		}
		goto tr848
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto tr849
		case 33:
			goto tr850
		case 34:
			goto tr851
		case 35:
			goto tr852
		case 36:
			goto tr853
		case 38:
			goto tr855
		case 39:
			goto tr856
		case 43:
			goto tr857
		case 45:
			goto tr858
		case 46:
			goto tr859
		case 47:
			goto tr860
		case 48:
			goto tr861
		case 58:
			goto tr863
		case 60:
			goto tr864
		case 62:
			goto tr865
		case 63:
			goto tr866
		case 64:
			goto tr867
		case 66:
			goto tr869
		case 88:
			goto tr870
		case 95:
			goto tr871
		case 98:
			goto tr873
		case 101:
			goto tr874
		case 120:
			goto tr875
		case 124:
			goto tr876
		case 126:
			goto tr877
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
						goto tr854
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr849
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr868
					}
				case data[p] >= 59:
					goto tr854
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr854
					}
				case data[p] >= 97:
					goto tr872
				}
			default:
				goto tr854
			}
		default:
			goto tr862
		}
		goto tr878
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto tr879
		case 33:
			goto tr880
		case 34:
			goto tr881
		case 35:
			goto tr882
		case 36:
			goto tr883
		case 38:
			goto tr885
		case 39:
			goto tr886
		case 43:
			goto tr887
		case 45:
			goto tr888
		case 46:
			goto tr889
		case 47:
			goto tr890
		case 48:
			goto tr891
		case 58:
			goto tr893
		case 59:
			goto tr884
		case 60:
			goto tr894
		case 61:
			goto st121
		case 62:
			goto tr896
		case 63:
			goto tr897
		case 64:
			goto tr898
		case 66:
			goto tr900
		case 88:
			goto tr901
		case 95:
			goto tr902
		case 98:
			goto tr904
		case 101:
			goto tr905
		case 120:
			goto tr906
		case 124:
			goto tr907
		case 126:
			goto tr908
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
						goto tr884
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr879
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr899
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr884
					}
				case data[p] >= 97:
					goto tr903
				}
			default:
				goto tr884
			}
		default:
			goto tr892
		}
		goto tr909
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto tr910
		case 33:
			goto tr911
		case 34:
			goto tr912
		case 35:
			goto tr913
		case 36:
			goto tr914
		case 38:
			goto tr916
		case 39:
			goto tr917
		case 43:
			goto tr918
		case 45:
			goto tr919
		case 46:
			goto tr920
		case 47:
			goto tr921
		case 48:
			goto tr922
		case 58:
			goto tr924
		case 60:
			goto tr925
		case 62:
			goto tr926
		case 63:
			goto tr927
		case 64:
			goto tr928
		case 66:
			goto tr930
		case 88:
			goto tr931
		case 95:
			goto tr932
		case 98:
			goto tr934
		case 101:
			goto tr935
		case 120:
			goto tr936
		case 124:
			goto tr937
		case 126:
			goto tr938
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
						goto tr915
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr910
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr929
					}
				case data[p] >= 59:
					goto tr915
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr915
					}
				case data[p] >= 97:
					goto tr933
				}
			default:
				goto tr915
			}
		default:
			goto tr923
		}
		goto tr939
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr940
		case 33:
			goto tr941
		case 34:
			goto tr942
		case 35:
			goto tr943
		case 36:
			goto tr944
		case 38:
			goto tr946
		case 39:
			goto tr947
		case 43:
			goto tr948
		case 45:
			goto tr949
		case 46:
			goto tr950
		case 47:
			goto tr951
		case 48:
			goto tr952
		case 58:
			goto tr954
		case 60:
			goto tr955
		case 62:
			goto tr956
		case 63:
			goto tr957
		case 64:
			goto tr958
		case 66:
			goto tr960
		case 88:
			goto tr961
		case 95:
			goto tr962
		case 98:
			goto tr964
		case 101:
			goto tr965
		case 120:
			goto tr966
		case 124:
			goto tr967
		case 126:
			goto tr968
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
						goto tr945
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr940
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr959
					}
				case data[p] >= 59:
					goto tr945
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr945
					}
				case data[p] >= 97:
					goto tr963
				}
			default:
				goto tr945
			}
		default:
			goto tr953
		}
		goto tr969
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto tr970
		case 33:
			goto tr971
		case 34:
			goto tr972
		case 35:
			goto tr973
		case 36:
			goto tr974
		case 38:
			goto tr976
		case 39:
			goto tr977
		case 43:
			goto tr978
		case 45:
			goto tr979
		case 46:
			goto tr980
		case 47:
			goto tr981
		case 48:
			goto tr982
		case 58:
			goto tr984
		case 60:
			goto tr985
		case 62:
			goto tr986
		case 63:
			goto tr987
		case 64:
			goto tr988
		case 66:
			goto tr990
		case 88:
			goto tr991
		case 95:
			goto tr992
		case 98:
			goto tr994
		case 101:
			goto tr995
		case 120:
			goto tr996
		case 124:
			goto tr997
		case 126:
			goto tr998
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
						goto tr975
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr970
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr989
					}
				case data[p] >= 59:
					goto tr975
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr975
					}
				case data[p] >= 97:
					goto tr993
				}
			default:
				goto tr975
			}
		default:
			goto tr983
		}
		goto tr999
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto tr1000
		case 33:
			goto tr1001
		case 34:
			goto tr1002
		case 35:
			goto tr1003
		case 36:
			goto tr1004
		case 38:
			goto tr1006
		case 39:
			goto tr1007
		case 43:
			goto tr1008
		case 45:
			goto tr1009
		case 46:
			goto tr1010
		case 47:
			goto tr1011
		case 48:
			goto tr1012
		case 58:
			goto tr1014
		case 60:
			goto tr1015
		case 62:
			goto tr1016
		case 63:
			goto tr1017
		case 64:
			goto tr1018
		case 66:
			goto tr1020
		case 88:
			goto tr1021
		case 95:
			goto tr1022
		case 98:
			goto tr1024
		case 101:
			goto tr1025
		case 120:
			goto tr1026
		case 124:
			goto tr1027
		case 126:
			goto tr1028
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
						goto tr1005
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1000
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1019
					}
				case data[p] >= 59:
					goto tr1005
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1005
					}
				case data[p] >= 97:
					goto tr1023
				}
			default:
				goto tr1005
			}
		default:
			goto tr1013
		}
		goto tr1029
tr331:
//line lex.rl:69

			isFconst = true
		
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lex.go:17048
		switch data[p] {
		case 43:
			goto st69
		case 45:
			goto st69
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto tr312
		case 33:
			goto tr313
		case 34:
			goto tr314
		case 35:
			goto tr315
		case 36:
			goto tr316
		case 38:
			goto tr318
		case 39:
			goto tr319
		case 43:
			goto tr320
		case 45:
			goto tr321
		case 46:
			goto tr322
		case 47:
			goto tr323
		case 58:
			goto tr324
		case 60:
			goto tr325
		case 62:
			goto tr326
		case 63:
			goto tr327
		case 64:
			goto tr328
		case 66:
			goto tr330
		case 88:
			goto tr332
		case 95:
			goto tr333
		case 98:
			goto tr335
		case 101:
			goto tr1030
		case 120:
			goto tr336
		case 124:
			goto tr337
		case 126:
			goto tr338
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
						goto tr317
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr312
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr329
					}
				case data[p] >= 59:
					goto tr317
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr317
					}
				case data[p] >= 97:
					goto tr334
				}
			default:
				goto tr317
			}
		default:
			goto st125
		}
		goto tr339
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 32:
			goto tr1031
		case 33:
			goto tr1032
		case 34:
			goto tr1033
		case 35:
			goto tr1034
		case 36:
			goto tr1035
		case 38:
			goto tr1037
		case 39:
			goto tr1038
		case 43:
			goto tr1039
		case 45:
			goto tr1040
		case 46:
			goto tr1041
		case 47:
			goto tr1042
		case 58:
			goto tr1043
		case 59:
			goto tr1036
		case 60:
			goto tr1044
		case 61:
			goto tr1036
		case 62:
			goto tr1045
		case 63:
			goto tr1046
		case 64:
			goto tr1047
		case 88:
			goto tr1049
		case 95:
			goto tr1050
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
						goto tr1036
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1031
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto tr1036
					}
				case data[p] >= 71:
					goto tr1048
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1036
					}
				case data[p] >= 103:
					goto tr1051
				}
			default:
				goto st126
			}
		default:
			goto st126
		}
		goto tr1055
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
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
		case 48:
			goto tr1068
		case 58:
			goto tr1070
		case 60:
			goto tr1071
		case 62:
			goto st128
		case 63:
			goto tr1073
		case 64:
			goto tr1074
		case 66:
			goto tr1076
		case 88:
			goto tr1077
		case 95:
			goto tr1078
		case 98:
			goto tr1080
		case 101:
			goto tr1081
		case 120:
			goto tr1082
		case 124:
			goto tr1083
		case 126:
			goto tr1084
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
						goto tr1075
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
					goto tr1079
				}
			default:
				goto tr1061
			}
		default:
			goto tr1069
		}
		goto tr1085
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto tr1086
		case 33:
			goto tr1087
		case 34:
			goto tr1088
		case 35:
			goto tr1089
		case 36:
			goto tr1090
		case 38:
			goto tr1092
		case 39:
			goto tr1093
		case 43:
			goto tr1094
		case 45:
			goto tr1095
		case 46:
			goto tr1096
		case 47:
			goto tr1097
		case 48:
			goto tr1098
		case 58:
			goto tr1100
		case 60:
			goto tr1101
		case 62:
			goto tr1102
		case 63:
			goto tr1103
		case 64:
			goto tr1104
		case 66:
			goto tr1106
		case 88:
			goto tr1107
		case 95:
			goto tr1108
		case 98:
			goto tr1110
		case 101:
			goto tr1111
		case 120:
			goto tr1112
		case 124:
			goto tr1113
		case 126:
			goto tr1114
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
						goto tr1091
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1086
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1105
					}
				case data[p] >= 59:
					goto tr1091
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1091
					}
				case data[p] >= 97:
					goto tr1109
				}
			default:
				goto tr1091
			}
		default:
			goto tr1099
		}
		goto tr1115
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 32:
			goto tr1116
		case 33:
			goto tr1117
		case 34:
			goto tr1118
		case 35:
			goto tr1119
		case 36:
			goto tr1120
		case 38:
			goto tr1122
		case 39:
			goto tr1123
		case 43:
			goto tr1124
		case 45:
			goto tr1125
		case 46:
			goto tr1126
		case 47:
			goto tr1127
		case 58:
			goto tr1128
		case 60:
			goto tr1129
		case 62:
			goto tr1130
		case 63:
			goto tr1131
		case 64:
			goto tr1132
		case 66:
			goto tr1134
		case 88:
			goto tr1135
		case 95:
			goto tr1136
		case 98:
			goto tr1138
		case 101:
			goto tr1139
		case 120:
			goto tr1140
		case 124:
			goto tr1141
		case 126:
			goto tr1142
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
						goto tr1121
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1116
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1133
					}
				case data[p] >= 59:
					goto tr1121
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1121
					}
				case data[p] >= 97:
					goto tr1137
				}
			default:
				goto tr1121
			}
		default:
			goto st129
		}
		goto tr1143
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto tr1144
		case 33:
			goto tr1145
		case 34:
			goto tr1146
		case 35:
			goto tr1147
		case 36:
			goto tr1148
		case 38:
			goto tr1150
		case 39:
			goto tr1151
		case 43:
			goto tr1152
		case 45:
			goto tr1153
		case 46:
			goto tr1154
		case 47:
			goto tr1155
		case 48:
			goto tr1156
		case 58:
			goto tr1158
		case 60:
			goto tr1159
		case 62:
			goto tr1160
		case 63:
			goto tr1161
		case 64:
			goto tr1162
		case 66:
			goto tr1164
		case 88:
			goto tr1165
		case 95:
			goto tr1166
		case 98:
			goto tr1168
		case 101:
			goto tr1169
		case 120:
			goto tr1170
		case 124:
			goto tr1171
		case 126:
			goto tr1172
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
						goto tr1149
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1144
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1163
					}
				case data[p] >= 59:
					goto tr1149
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1149
					}
				case data[p] >= 97:
					goto tr1167
				}
			default:
				goto tr1149
			}
		default:
			goto tr1157
		}
		goto tr1173
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr1174
		case 33:
			goto tr1175
		case 34:
			goto tr1176
		case 35:
			goto tr1177
		case 36:
			goto tr1178
		case 38:
			goto tr1180
		case 39:
			goto tr1181
		case 43:
			goto tr1182
		case 45:
			goto tr1183
		case 46:
			goto tr1184
		case 47:
			goto tr1185
		case 48:
			goto tr1186
		case 58:
			goto tr1188
		case 60:
			goto tr1189
		case 62:
			goto st132
		case 63:
			goto tr1191
		case 64:
			goto tr1192
		case 66:
			goto tr1194
		case 88:
			goto tr1195
		case 95:
			goto tr1196
		case 98:
			goto tr1198
		case 101:
			goto tr1199
		case 120:
			goto tr1200
		case 124:
			goto tr1201
		case 126:
			goto tr1202
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
						goto tr1179
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1174
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1193
					}
				case data[p] >= 59:
					goto tr1179
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1179
					}
				case data[p] >= 97:
					goto tr1197
				}
			default:
				goto tr1179
			}
		default:
			goto tr1187
		}
		goto tr1203
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 32:
			goto tr1204
		case 33:
			goto tr1205
		case 34:
			goto tr1206
		case 35:
			goto tr1207
		case 36:
			goto tr1208
		case 38:
			goto tr1210
		case 39:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 46:
			goto tr1214
		case 47:
			goto tr1215
		case 48:
			goto tr1216
		case 58:
			goto tr1218
		case 60:
			goto tr1219
		case 62:
			goto tr1220
		case 63:
			goto tr1221
		case 64:
			goto tr1222
		case 66:
			goto tr1224
		case 88:
			goto tr1225
		case 95:
			goto tr1226
		case 98:
			goto tr1228
		case 101:
			goto tr1229
		case 120:
			goto tr1230
		case 124:
			goto tr1231
		case 126:
			goto tr1232
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
						goto tr1209
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1204
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1223
					}
				case data[p] >= 59:
					goto tr1209
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1209
					}
				case data[p] >= 97:
					goto tr1227
				}
			default:
				goto tr1209
			}
		default:
			goto tr1217
		}
		goto tr1233
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 32:
			goto tr1234
		case 33:
			goto tr1235
		case 34:
			goto tr1236
		case 35:
			goto tr1237
		case 36:
			goto tr1238
		case 38:
			goto tr1240
		case 39:
			goto tr1241
		case 43:
			goto tr1242
		case 45:
			goto tr1243
		case 46:
			goto tr1244
		case 47:
			goto tr1245
		case 48:
			goto tr1246
		case 58:
			goto tr1248
		case 60:
			goto tr1249
		case 62:
			goto tr1250
		case 63:
			goto tr1251
		case 64:
			goto tr1252
		case 66:
			goto tr1254
		case 88:
			goto tr1255
		case 95:
			goto tr1256
		case 98:
			goto tr1258
		case 101:
			goto tr1259
		case 120:
			goto tr1260
		case 124:
			goto tr1261
		case 126:
			goto tr1262
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
						goto tr1239
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1234
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1253
					}
				case data[p] >= 59:
					goto tr1239
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1239
					}
				case data[p] >= 97:
					goto tr1257
				}
			default:
				goto tr1239
			}
		default:
			goto tr1247
		}
		goto tr1263
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 32:
			goto tr1264
		case 33:
			goto tr1265
		case 34:
			goto tr1266
		case 35:
			goto tr1267
		case 36:
			goto tr1268
		case 38:
			goto tr1270
		case 39:
			goto tr1271
		case 42:
			goto st135
		case 43:
			goto tr1273
		case 45:
			goto tr1274
		case 46:
			goto tr1275
		case 47:
			goto tr1276
		case 48:
			goto tr1277
		case 58:
			goto tr1279
		case 60:
			goto tr1280
		case 62:
			goto tr1281
		case 63:
			goto tr1282
		case 64:
			goto tr1283
		case 66:
			goto tr1285
		case 88:
			goto tr1286
		case 95:
			goto tr1287
		case 98:
			goto tr1289
		case 101:
			goto tr1290
		case 120:
			goto tr1291
		case 124:
			goto tr1292
		case 126:
			goto tr1293
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
						goto tr1269
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1264
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1284
					}
				case data[p] >= 59:
					goto tr1269
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1269
					}
				case data[p] >= 97:
					goto tr1288
				}
			default:
				goto tr1269
			}
		default:
			goto tr1278
		}
		goto tr1294
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 32:
			goto tr1295
		case 33:
			goto tr1296
		case 34:
			goto tr1297
		case 35:
			goto tr1298
		case 36:
			goto tr1299
		case 38:
			goto tr1301
		case 39:
			goto tr1302
		case 43:
			goto tr1303
		case 45:
			goto tr1304
		case 46:
			goto tr1305
		case 47:
			goto tr1306
		case 48:
			goto tr1307
		case 58:
			goto tr1309
		case 60:
			goto tr1310
		case 62:
			goto tr1311
		case 63:
			goto tr1312
		case 64:
			goto tr1313
		case 66:
			goto tr1315
		case 88:
			goto tr1316
		case 95:
			goto tr1317
		case 98:
			goto tr1319
		case 101:
			goto tr1320
		case 120:
			goto tr1321
		case 124:
			goto tr1322
		case 126:
			goto tr1323
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
						goto tr1300
					}
				case data[p] >= 14:
					goto st0
				}
			default:
				goto tr1295
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1314
					}
				case data[p] >= 59:
					goto tr1300
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1300
					}
				case data[p] >= 97:
					goto tr1318
				}
			default:
				goto tr1300
			}
		default:
			goto tr1308
		}
		goto tr1324
	st_out:
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
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
	_test_eof106: cs = 106; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
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
	_test_eof67: cs = 67; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 129:
//line lex.rl:43

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
		
		case 85, 90, 91, 125:
//line lex.rl:55

			str()
			if isFconst {
				emit(lex.FCONST, s)
			} else {
				emit(lex.ICONST, s)
			}
			isFconst = false
		
		case 126:
//line lex.rl:72

			str()
			emit(lex.ICONST, s)
		
		case 104, 106, 108, 113, 115:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
		case 74:
//line lex.rl:130

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
		
		case 81:
//line lex.rl:164

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
		
		case 109, 114:
//line lex.rl:239

			emit(lex.BCONST, buf.String())
		
		case 116:
//line lex.rl:243

			if !utf8.Valid(buf.Bytes()) {
				return fmt.Errorf("invalid UTF-8 string")
			}
			emit(lex.SCONST, buf.String())
		
		case 107:
//line lex.rl:265

			emit(lex.BITCONST, string(data[mark+2:p-1]))
		
		case 72, 73, 75, 76, 77, 78, 80, 82, 83, 86, 88, 92, 95, 98, 100, 102, 105, 111:
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
		case 87:
//line lex.rl:284
 emitToken(lex.DOT_DOT) 
		case 133:
//line lex.rl:286
 emitToken(lex.NOT_EQUALS) 
		case 135:
//line lex.rl:287
 emitToken(lex.NOT_REGIMATCH) 
		case 134:
//line lex.rl:288
 emitToken(lex.NOT_REGMATCH) 
		case 118:
//line lex.rl:290
 emitToken(lex.HELPTOKEN) 
		case 119:
//line lex.rl:291
 emitToken(lex.JSON_SOME_EXISTS) 
		case 101:
//line lex.rl:292
 emitToken(lex.JSON_ALL_EXISTS) 
		case 97:
//line lex.rl:294
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS) 
		case 96:
//line lex.rl:295
 emitToken(lex.LSHIFT) 
		case 123:
//line lex.rl:296
 emitToken(lex.NOT_EQUALS) 
		case 122:
//line lex.rl:297
 emitToken(lex.LESS_EQUALS) 
		case 124:
//line lex.rl:298
 emitToken(lex.CONTAINED_BY) 
		case 121:
//line lex.rl:300
 emitToken(lex.INET_CONTAINS_OR_EQUALS) 
		case 120:
//line lex.rl:301
 emitToken(lex.RSHIFT) 
		case 99:
//line lex.rl:302
 emitToken(lex.GREATER_EQUALS) 
		case 94:
//line lex.rl:304
 emitToken(lex.TYPEANNOTATE) 
		case 93:
//line lex.rl:305
 emitToken(lex.TYPECAST) 
		case 117:
//line lex.rl:307
 emitToken(lex.CONCAT) 
		case 89:
//line lex.rl:309
 emitToken(lex.FLOORDIV) 
		case 112:
//line lex.rl:311
 emitToken(lex.REGIMATCH) 
		case 103:
//line lex.rl:313
 emitToken(lex.CONTAINS) 
		case 79:
//line lex.rl:315
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY) 
		case 128:
//line lex.rl:317
 emitToken(lex.FETCHTEXT) 
		case 127:
//line lex.rl:318
 emitToken(lex.FETCHVAL) 
		case 132:
//line lex.rl:320
 emitToken(lex.FETCHTEXT_PATH) 
		case 131:
//line lex.rl:321
 emitToken(lex.FETCHVAL_PATH) 
		case 130:
//line lex.rl:322
 emitToken(lex.REMOVE_PATH) 
		case 110:
//line lex.rl:76

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

			if id, ok := lex.Keywords[s]; ok {
				emit(Tok(id.Tok), s)
			} else {
				emit(lex.IDENT, s)
			}
			isUpper = false
			isNotASCII = false
		
//line lex.rl:282
 emitToken(Tok(data[p-1])) 
//line lex.go:18618
		}
	}

	_out: {}
	}

//line lex.rl:331


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
