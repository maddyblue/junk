
//line lex.rl:1
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var _ = fmt.Print

func lexSQL(data []rune) (ret []string, err error) {
	
//line lex.rl:15
	
//line lex.rl:16
	
//line lex.go:22
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 18, 1, 19, 1, 20, 1, 23, 
	1, 28, 1, 29, 1, 30, 2, 0, 
	26, 2, 0, 27, 2, 4, 0, 2, 
	8, 0, 2, 9, 0, 2, 10, 0, 
	2, 11, 0, 2, 12, 0, 2, 13, 
	0, 2, 14, 0, 2, 16, 17, 2, 
	22, 23, 2, 25, 0, 3, 17, 15, 
	0, 3, 18, 21, 0, 3, 19, 21, 
	0, 3, 23, 24, 0, 3, 29, 0, 
	26, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 2, 4, 5, 6, 9, 
	13, 26, 28, 30, 36, 42, 48, 54, 
	60, 66, 72, 78, 84, 90, 96, 102, 
	108, 114, 124, 134, 140, 140, 141, 147, 
	153, 153, 154, 
}

var _scanner_trans_keys []int32 = []int32{
	42, 39, 42, 47, 47, 42, 39, 0, 
	127, 39, 92, 0, 127, 85, 88, 97, 
	98, 102, 110, 114, 116, 117, 118, 120, 
	48, 55, 48, 55, 48, 55, 48, 57, 
	65, 70, 97, 102, 48, 57, 65, 70, 
	97, 102, 48, 57, 65, 70, 97, 102, 
	48, 57, 65, 70, 97, 102, 48, 57, 
	65, 70, 97, 102, 48, 57, 65, 70, 
	97, 102, 48, 57, 65, 70, 97, 102, 
	48, 57, 65, 70, 97, 102, 48, 57, 
	65, 70, 97, 102, 48, 57, 65, 70, 
	97, 102, 48, 57, 65, 70, 97, 102, 
	48, 57, 65, 70, 97, 102, 48, 57, 
	65, 70, 97, 102, 48, 57, 65, 70, 
	97, 102, 32, 39, 47, 101, 9, 10, 
	12, 13, 48, 57, 32, 39, 47, 101, 
	9, 10, 12, 13, 48, 57, 9, 10, 
	32, 39, 12, 13, 10, 32, 39, 9, 
	10, 12, 13, 9, 10, 32, 39, 12, 
	13, 10, 32, 39, 9, 10, 12, 13, 
	
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 1, 2, 1, 1, 1, 2, 
	11, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 4, 4, 4, 0, 1, 2, 4, 
	0, 1, 2, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 1, 1, 
	1, 1, 1, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 1, 0, 0, 2, 1, 
	0, 0, 2, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 4, 7, 9, 11, 14, 
	18, 31, 33, 35, 39, 43, 47, 51, 
	55, 59, 63, 67, 71, 75, 79, 83, 
	87, 91, 99, 107, 113, 114, 116, 121, 
	127, 128, 130, 
}

var _scanner_indicies []byte = []byte{
	0, 1, 2, 1, 4, 5, 3, 6, 
	3, 7, 3, 10, 9, 8, 13, 14, 
	12, 11, 17, 18, 19, 20, 21, 22, 
	23, 24, 25, 26, 18, 16, 15, 28, 
	27, 29, 27, 30, 30, 30, 27, 31, 
	31, 31, 27, 32, 32, 32, 27, 33, 
	33, 33, 27, 34, 34, 34, 27, 35, 
	35, 35, 27, 36, 36, 36, 27, 37, 
	37, 37, 27, 38, 38, 38, 27, 39, 
	39, 39, 27, 40, 40, 40, 27, 41, 
	41, 41, 27, 42, 42, 42, 27, 43, 
	43, 43, 27, 45, 46, 47, 49, 45, 
	45, 48, 44, 50, 51, 52, 54, 50, 
	50, 53, 44, 56, 57, 56, 9, 56, 
	55, 1, 57, 55, 57, 58, 57, 57, 
	55, 60, 61, 60, 12, 60, 59, 1, 
	61, 59, 61, 62, 61, 61, 59, 
}

var _scanner_trans_targs []byte = []byte{
	25, 0, 25, 3, 4, 5, 3, 3, 
	6, 6, 27, 7, 7, 31, 8, 7, 
	9, 11, 19, 7, 7, 7, 7, 7, 
	7, 21, 7, 0, 10, 7, 12, 13, 
	14, 15, 16, 17, 18, 7, 20, 7, 
	22, 23, 24, 7, 0, 25, 25, 1, 
	26, 2, 25, 25, 1, 26, 2, 28, 
	29, 30, 6, 32, 33, 34, 7, 
}

var _scanner_trans_actions []byte = []byte{
	25, 0, 34, 0, 0, 0, 5, 3, 
	15, 0, 37, 15, 0, 37, 9, 67, 
	64, 21, 0, 40, 43, 46, 49, 52, 
	55, 21, 58, 13, 23, 82, 19, 19, 
	19, 19, 19, 19, 19, 78, 61, 70, 
	17, 17, 17, 74, 29, 0, 31, 0, 
	1, 0, 27, 86, 27, 0, 27, 11, 
	0, 0, 1, 11, 0, 0, 1, 
}

var _scanner_eof_actions []byte = []byte{
	0, 0, 0, 7, 7, 7, 13, 13, 
	13, 13, 13, 13, 13, 13, 13, 13, 
	13, 13, 13, 13, 13, 13, 13, 13, 
	13, 0, 27, 11, 0, 11, 11, 11, 
	0, 11, 11, 
}

const scanner_start int = 25
const scanner_first_final int = 25
const scanner_error int = 0

const scanner_en_comment int = 3
const scanner_en_singleQuoteNoEscape int = 6
const scanner_en_singleQuoteEscape int = 7
const scanner_en_main int = 25


//line lex.rl:17

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var (
		mark int
		stack [1]int
		top int
		commentDepth int
		buf bytes.Buffer
		isNotASCII, checkUTF8 bool
		ch byte
		rn rune
		_ = rn
	)

	
//line lex.go:170
	{
	cs = scanner_start
	top = 0
	}

//line lex.go:176
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_scanner_key_offsets[cs])
	_trans = int(_scanner_index_offsets[cs])

	_klen = int(_scanner_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _scanner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_scanner_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_scanner_indicies[_trans])
	cs = int(_scanner_trans_targs[_trans])

	if _scanner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_scanner_trans_actions[_trans])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _scanner_actions[_acts-1] {
		case 0:
//line lex.rl:32

			//fmt.Println("mark", p)
			mark = p
		
		case 1:
//line lex.rl:36

			commentDepth++
		
		case 2:
//line lex.rl:39

			commentDepth--
			if commentDepth == 0 {
				top--; cs = stack[top]
goto _again

			}
		
		case 4:
//line lex.rl:67

			//fmt.Println("write buf", string(data[mark+1:p]))
			buf.WriteString(string(data[mark+1:p]))
		
		case 5:
//line lex.rl:74

			if checkUTF8 && isNotASCII && !utf8.Valid(buf.Bytes()) {
				return nil, errors.New("invalid utf8 encoding")
			}
			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
		case 6:
//line lex.rl:82

			return nil, errors.New("unterminated string")
		
		case 7:
//line lex.rl:89
 isNotASCII = true 
		case 8:
//line lex.rl:118
 buf.WriteByte('\a') 
		case 9:
//line lex.rl:119
 buf.WriteByte('\b') 
		case 10:
//line lex.rl:120
 buf.WriteByte('\f') 
		case 11:
//line lex.rl:121
 buf.WriteByte('\n') 
		case 12:
//line lex.rl:122
 buf.WriteByte('\r') 
		case 13:
//line lex.rl:123
 buf.WriteByte('\t') 
		case 14:
//line lex.rl:124
 buf.WriteByte('\v') 
		case 15:
//line lex.rl:126
 buf.WriteByte(ch) 
		case 16:
//line lex.rl:130
 ch = 0 
		case 17:
//line lex.rl:131
 ch = (ch << 4) | unhex(data[p]) 
		case 18:
//line lex.rl:137
 rn = (rn << 4) | rune(unhex(data[p])) 
		case 19:
//line lex.rl:140
 rn = (rn << 4) | rune(unhex(data[p])) 
		case 20:
//line lex.rl:142
 rn = 0 
		case 21:
//line lex.rl:143
 buf.WriteRune(rn) 
		case 22:
//line lex.rl:147
 ch = 0 
		case 23:
//line lex.rl:148
 ch = (ch << 3) | byte(data[p]) - '0' 
		case 24:
//line lex.rl:149
 buf.WriteByte(ch) 
		case 25:
//line lex.rl:151

			//fmt.Println("write rune", string(data[p]))
			buf.WriteRune(data[p])
		
		case 26:
//line lex.rl:190

			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			stack[top] = cs; top++; cs = 6; goto _again

		
		case 27:
//line lex.rl:196

			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			stack[top] = cs; top++; cs = 7; goto _again

		
		case 28:
//line lex.rl:202

			commentDepth = 1
			stack[top] = cs; top++; cs = 3; goto _again

		
		case 29:
//line lex.rl:206

			ret = append(ret, string(data[mark:p]))
		
		case 30:
//line lex.rl:217
 return nil, errors.Errorf("ERROR at %d of %d", p, len(data)) 
//line lex.go:393
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _scanner_eof_actions[cs]
		__nacts := uint(_scanner_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _scanner_actions[__acts-1] {
			case 3:
//line lex.rl:45

			return nil, errors.New("unterminated comment")
		
			case 5:
//line lex.rl:74

			if checkUTF8 && isNotASCII && !utf8.Valid(buf.Bytes()) {
				return nil, errors.New("invalid utf8 encoding")
			}
			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
			case 6:
//line lex.rl:82

			return nil, errors.New("unterminated string")
		
			case 29:
//line lex.rl:206

			ret = append(ret, string(data[mark:p]))
		
//line lex.go:440
			}
		}
	}

	_out: {}
	}

//line lex.rl:222


	return ret, nil
}

func unhex(r rune) byte {
	c := byte(r)
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
