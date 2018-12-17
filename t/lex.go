
//line lex.rl:1
package main

import (
	"fmt"
	"bytes"
	"github.com/pkg/errors"
)

var _ = fmt.Println

func lexSQL(data []byte) (ret []string, err error) {
	
//line lex.rl:13
	
//line lex.go:18
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 5, 1, 6, 1, 7, 1, 8, 
	1, 10, 2, 0, 9, 2, 4, 0, 
	3, 8, 0, 9, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 3, 4, 5, 6, 15, 
	24, 30, 30, 31, 
}

var _scanner_trans_keys []byte = []byte{
	42, 42, 47, 47, 42, 39, 32, 39, 
	47, 9, 10, 12, 13, 48, 57, 32, 
	39, 47, 9, 10, 12, 13, 48, 57, 
	9, 10, 32, 39, 12, 13, 10, 32, 
	39, 9, 10, 12, 13, 
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 2, 1, 1, 1, 3, 3, 
	4, 0, 1, 2, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 3, 3, 
	1, 0, 0, 2, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 5, 7, 9, 11, 18, 
	25, 31, 32, 34, 
}

var _scanner_trans_targs []byte = []byte{
	6, 0, 3, 4, 2, 2, 2, 2, 
	2, 8, 5, 6, 6, 1, 6, 6, 
	7, 0, 6, 6, 1, 6, 6, 7, 
	0, 10, 11, 10, 5, 10, 9, 0, 
	11, 9, 11, 5, 11, 11, 9, 
}

var _scanner_trans_actions []byte = []byte{
	13, 0, 0, 0, 0, 5, 0, 3, 
	0, 22, 0, 0, 19, 0, 0, 0, 
	1, 17, 15, 25, 15, 15, 15, 0, 
	17, 0, 0, 0, 0, 0, 9, 0, 
	0, 9, 0, 1, 0, 0, 9, 
}

var _scanner_eof_actions []byte = []byte{
	0, 0, 7, 7, 7, 11, 0, 15, 
	9, 0, 9, 9, 
}

const scanner_start int = 6
const scanner_first_final int = 6
const scanner_error int = 0

const scanner_en_comment int = 2
const scanner_en_singleQuote int = 5
const scanner_en_main int = 6


//line lex.rl:14

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof
	var (
		mark int
		_ = mark
		stack [1]int
		top int
		_, _ = stack, top
		commentDepth int
		buf bytes.Buffer
		_ = buf
	)

	
//line lex.go:100
	{
	cs = scanner_start
	top = 0
	}

//line lex.go:106
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
//line lex.rl:58

			buf.WriteString(string(data[mark+1:p]))
		
		case 5:
//line lex.rl:64

			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
		case 7:
//line lex.rl:103
 commentDepth = 1; stack[top] = cs; top++; cs = 2; goto _again
 
		case 8:
//line lex.rl:104
 ret = append(ret, string(data[mark:p])) 
		case 9:
//line lex.rl:105
 buf.Reset(); stack[top] = cs; top++; cs = 5; goto _again
 
		case 10:
//line lex.rl:108
 return nil, errors.Errorf("ERROR at %d of %d", p, len(data)) 
//line lex.go:231
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
//line lex.rl:55
 return nil, errors.New("unterminated comment") 
			case 5:
//line lex.rl:64

			ret = append(ret, buf.String())
			p--

			top--; cs = stack[top]
goto _again

		
			case 6:
//line lex.rl:69

			return nil, errors.New("unterminated string")
		
			case 8:
//line lex.rl:104
 ret = append(ret, string(data[mark:p])) 
//line lex.go:271
			}
		}
	}

	_out: {}
	}

//line lex.rl:113


	return ret, nil
}
