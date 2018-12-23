
//line lex.rl:1
package main

import (
	"bytes"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func lexSQL(data []rune) (ret []string, err error) {
	
//line lex.rl:12
	
//line lex.rl:13
	
//line lex.go:19
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 5, 1, 6, 1, 7, 1, 13, 
	1, 17, 1, 18, 1, 19, 2, 0, 
	15, 2, 0, 16, 2, 4, 0, 2, 
	10, 0, 2, 12, 0, 2, 13, 5, 
	2, 13, 7, 2, 13, 8, 3, 13, 
	4, 0, 3, 13, 7, 5, 3, 13, 
	7, 8, 3, 18, 0, 15, 4, 12, 
	11, 14, 0, 4, 13, 4, 0, 5, 
	4, 13, 4, 0, 8, 4, 13, 9, 
	10, 0, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 2, 4, 5, 6, 9, 
	13, 15, 27, 43, 51, 55, 59, 69, 
	79, 85, 85, 86, 92, 99, 103, 108, 
	110, 117, 127, 138, 142, 147, 157, 169, 
}

var _scanner_trans_keys []int32 = []int32{
	42, 39, 42, 47, 47, 42, 39, 0, 
	127, 39, 92, 0, 127, 39, 92, 85, 
	88, 102, 110, 114, 120, 48, 55, 97, 
	98, 116, 118, 39, 92, 0, 47, 48, 
	57, 58, 64, 65, 70, 71, 96, 97, 
	102, 103, 127, 39, 92, 48, 57, 65, 
	70, 97, 102, 39, 92, 0, 127, 39, 
	92, 0, 127, 32, 39, 47, 101, 9, 
	10, 12, 13, 48, 57, 32, 39, 47, 
	101, 9, 10, 12, 13, 48, 57, 9, 
	10, 32, 39, 12, 13, 10, 32, 39, 
	9, 10, 12, 13, 9, 10, 32, 39, 
	92, 12, 13, 39, 92, 0, 127, 10, 
	39, 92, 0, 127, 39, 92, 32, 39, 
	92, 9, 10, 12, 13, 11, 32, 39, 
	92, 0, 8, 9, 13, 14, 127, 9, 
	10, 32, 39, 92, 0, 11, 12, 13, 
	14, 127, 39, 92, 0, 127, 10, 39, 
	92, 0, 127, 11, 32, 39, 92, 0, 
	8, 9, 13, 14, 127, 85, 88, 102, 
	110, 114, 120, 48, 55, 97, 98, 116, 
	118, 
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 1, 2, 1, 1, 1, 2, 
	2, 6, 2, 2, 2, 2, 4, 4, 
	4, 0, 1, 2, 5, 2, 3, 2, 
	3, 4, 5, 2, 3, 4, 6, 0, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 1, 1, 
	0, 3, 7, 3, 1, 1, 3, 3, 
	1, 0, 0, 2, 1, 1, 1, 0, 
	2, 3, 3, 1, 1, 3, 3, 0, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 4, 7, 9, 11, 14, 
	18, 21, 31, 41, 47, 51, 55, 63, 
	71, 77, 78, 80, 85, 92, 96, 101, 
	104, 110, 118, 127, 131, 136, 144, 154, 
}

var _scanner_indicies []byte = []byte{
	0, 1, 2, 1, 4, 5, 3, 6, 
	3, 7, 3, 10, 9, 8, 13, 14, 
	12, 11, 16, 16, 15, 15, 18, 15, 
	15, 15, 18, 15, 15, 15, 17, 13, 
	14, 12, 19, 12, 19, 12, 19, 12, 
	11, 16, 16, 20, 20, 20, 15, 23, 
	24, 22, 21, 13, 14, 25, 11, 27, 
	28, 29, 31, 27, 27, 30, 26, 32, 
	33, 34, 36, 32, 32, 35, 26, 38, 
	39, 38, 9, 38, 37, 1, 39, 37, 
	39, 40, 39, 39, 37, 42, 43, 42, 
	15, 44, 42, 41, 13, 14, 12, 11, 
	47, 48, 49, 46, 45, 1, 1, 15, 
	43, 50, 44, 43, 43, 41, 46, 47, 
	51, 49, 46, 47, 46, 45, 53, 54, 
	53, 51, 49, 52, 53, 52, 45, 13, 
	14, 25, 11, 54, 48, 49, 52, 45, 
	52, 54, 51, 49, 52, 54, 52, 45, 
	15, 18, 15, 15, 15, 18, 15, 15, 
	15, 17, 1, 
}

var _scanner_trans_targs []byte = []byte{
	14, 0, 14, 3, 4, 5, 3, 3, 
	6, 6, 16, 7, 8, 20, 9, 7, 
	0, 7, 10, 11, 12, 7, 8, 20, 
	9, 13, 0, 14, 14, 1, 15, 2, 
	14, 14, 1, 15, 2, 17, 18, 19, 
	6, 21, 22, 25, 31, 21, 23, 24, 
	20, 30, 7, 26, 27, 28, 29, 
}

var _scanner_trans_actions []byte = []byte{
	17, 0, 26, 0, 0, 0, 5, 3, 
	13, 0, 29, 41, 15, 47, 15, 0, 
	11, 63, 35, 78, 32, 55, 44, 73, 
	44, 15, 21, 0, 23, 0, 1, 0, 
	19, 59, 19, 0, 19, 9, 0, 0, 
	1, 9, 0, 0, 9, 51, 38, 15, 
	68, 38, 1, 47, 38, 15, 15, 
}

var _scanner_eof_actions []byte = []byte{
	0, 0, 0, 7, 7, 7, 11, 11, 
	11, 11, 11, 11, 11, 11, 0, 19, 
	9, 0, 9, 9, 9, 0, 9, 0, 
	9, 9, 9, 0, 9, 9, 0, 0, 
}

const scanner_start int = 14
const scanner_first_final int = 14
const scanner_error int = 0

const scanner_en_comment int = 3
const scanner_en_singleQuoteNoEscape int = 6
const scanner_en_singleQuoteEscape int = 7
const scanner_en_main int = 14


//line lex.rl:14

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

	
//line lex.go:163
	{
	cs = scanner_start
	top = 0
	}

//line lex.go:169
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
//line lex.rl:29
 mark = p 
		case 1:
//line lex.rl:30

			commentDepth++
		
		case 2:
//line lex.rl:33

			commentDepth--
			if commentDepth == 0 {
				top--; cs = stack[top]
goto _again

			}
		
		case 4:
//line lex.rl:61

			buf.WriteString(string(data[mark+1:p]))
		
		case 5:
//line lex.rl:67

			if checkUTF8 && isNotASCII && !utf8.Valid(buf.Bytes()) {
				return nil, errors.New("invalid utf8 encoding")
			}
			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
		case 6:
//line lex.rl:75

			return nil, errors.New("unterminated string")
		
		case 7:
//line lex.rl:82
 isNotASCII = true 
		case 8:
//line lex.rl:119
 buf.WriteByte(ch) 
		case 9:
//line lex.rl:123
 ch = 0 
		case 10:
//line lex.rl:124
 ch = (ch << 4) | unhex(data[p]) 
		case 11:
//line lex.rl:144
 buf.WriteRune(data[p]) 
		case 12:
//line lex.rl:145
 println("slash", string(data), p, string(data[p])) 
		case 13:
//line lex.rl:152
println("readstart", string(data[p]))
		case 14:
//line lex.rl:158
println("had slash")
		case 15:
//line lex.rl:181

			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			stack[top] = cs; top++; cs = 6; goto _again

		
		case 16:
//line lex.rl:187

			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			stack[top] = cs; top++; cs = 7; goto _again

		
		case 17:
//line lex.rl:193

			commentDepth = 1
			stack[top] = cs; top++; cs = 3; goto _again

		
		case 18:
//line lex.rl:197

			ret = append(ret, string(data[mark:p]))
		
		case 19:
//line lex.rl:208
 return nil, errors.Errorf("ERROR at %d of %d", p, len(data)) 
//line lex.go:346
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
//line lex.rl:39

			return nil, errors.New("unterminated comment")
		
			case 5:
//line lex.rl:67

			if checkUTF8 && isNotASCII && !utf8.Valid(buf.Bytes()) {
				return nil, errors.New("invalid utf8 encoding")
			}
			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
			case 6:
//line lex.rl:75

			return nil, errors.New("unterminated string")
		
			case 18:
//line lex.rl:197

			ret = append(ret, string(data[mark:p]))
		
//line lex.go:393
			}
		}
	}

	_out: {}
	}

//line lex.rl:213


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
