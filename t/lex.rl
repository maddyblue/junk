package main

import (
	"fmt"
	"bytes"
	"github.com/pkg/errors"
)

var _ = fmt.Println

func lexSQL(data []byte) (ret []string, err error) {
	%% machine scanner;
	%% write data;

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

	%%{
		action mark { mark = p }
		action start_comment {
			commentDepth++
		}
		action end_comment {
			commentDepth--
			if commentDepth == 0 {
				fret;
			}
		}
		ws = ' ' | '\t' | '\r' | '\f' | '\n';

		comment := (
			start: (
				'/' -> maybeStart |
				'*' -> maybeEnd |
				^[/*] -> start
			),
			maybeStart: (
				'*' @start_comment -> start |
				^[*] -> start
			),
			maybeEnd: (
				'/' @end_comment -> start |
				^[/] -> start
			)
			) @err{ return nil, errors.New("unterminated comment") }
			;
			
		action write_buf {
			buf.WriteString(string(data[mark+1:p]))
		}
		action write_single_quote {
			buf.WriteByte('\'')
		}
		action done_single_quote {
			ret = append(ret, buf.String())
			fhold;
			fret;
		}
		action string_err {
			return nil, errors.New("unterminated string")
		}
		not_newline = ws - '\n';
		sq = "'";
		action read {println("read", p, string(data[p]))}
		singleQuote := (
			start: (
				sq @write_buf @mark -> quote |
				^sq -> start
			),
			quote: (
				sq -> start |
				not_newline -> waitNewline |
				'\n' -> eatWs |
				^(ws | sq) @done_single_quote -> final |
				zlen %done_single_quote -> final
			),
			waitNewline: (
				'\n' -> eatWs |
				^[\n] @done_single_quote -> final |
				zlen %done_single_quote -> final
			),
			eatWs: (
				ws -> eatWs |
				sq @mark -> start |
				^(ws | sq) @done_single_quote -> final |
				zlen %done_single_quote -> final
			)
			) @err(string_err)
			;

		top =
			ws |
			'/*' @{ commentDepth = 1; fcall comment; } |
			digit+ >mark %{ ret = append(ret, string(data[mark:p])) } |
			sq @mark @{ buf.Reset(); fcall singleQuote; }
			;
		main :=
			top** %err{ return nil, errors.Errorf("ERROR at %d of %d", p, len(data)) }
			;

		write init;
		write exec;
	}%%

	return ret, nil
}
