package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var _ = fmt.Print

func lexSQL(data []rune) (ret []string, err error) {
	%% machine scanner;
	%% alphtype rune;
	%% write data;

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

	%%{
		action mark {
			//fmt.Println("mark", p)
			mark = p
		}
		action start_comment {
			commentDepth++
		}
		action end_comment {
			commentDepth--
			if commentDepth == 0 {
				fret;
			}
		}
		action comment_err {
			return nil, errors.New("unterminated comment")
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
			) @err(comment_err)
			;
			
		action write_buf {
			//fmt.Println("write buf", string(data[mark+1:p]))
			buf.WriteString(string(data[mark+1:p]))
		}
		action write_single_quote {
			buf.WriteByte('\'')
		}
		action done_quote {
			if checkUTF8 && isNotASCII && !utf8.Valid(buf.Bytes()) {
				return nil, errors.New("invalid utf8 encoding")
			}
			ret = append(ret, buf.String())
			fhold;
			fret;
		}
		action string_err {
			return nil, errors.New("unterminated string")
		}
		not_newline = ws - '\n';
		sq = "'";
		dq = '"';
		ASCII = 0..127;
		action notASCII { isNotASCII = true }
		notASCII = ^ASCII >notASCII;
		singleQuoteNoEscape := (
			start: (
				sq @write_buf @mark -> quote |
				(ASCII - sq) -> start |
				(notASCII - sq) -> start
			),
			quote: (
				sq -> start |
				not_newline -> waitNewline |
				'\n' -> eatWs |
				^(ws | sq) @done_quote -> final |
				zlen %done_quote -> final
			),
			waitNewline: (
				'\n' -> eatWs |
				^[\n] @done_quote -> final |
				zlen %done_quote -> final
			),
			eatWs: (
				ws -> eatWs |
				sq @mark -> start |
				^(ws | sq) @done_quote -> final |
				zlen %done_quote -> final
			)
			) @err(string_err)
			;
		escape =
			'a' @{ buf.WriteByte('\a') }
			| 'b' @{ buf.WriteByte('\b') }
			| 'f' @{ buf.WriteByte('\f') }
			| 'n' @{ buf.WriteByte('\n') }
			| 'r' @{ buf.WriteByte('\r') }
			| 't' @{ buf.WriteByte('\t') }
			| 'v' @{ buf.WriteByte('\v') }
			;
		action writeCh { buf.WriteByte(ch) }
		sl = "\\";
		slashHex =
			'x'i xdigit {2}
			>{ ch = 0 }
			${ ch = (ch << 4) | unhex(data[p]) }
			@writeCh
			;
		slashUnicode =
			((
				'u' xdigit {4}
				${ rn = (rn << 4) | rune(unhex(data[p])) }
			) | (
				'U' xdigit {8}
				${ rn = (rn << 4) | rune(unhex(data[p])) }
			))
			>{ rn = 0 }
			@{ buf.WriteRune(rn) }
			;
		slashOctal =
			('0'..'7') {3}
			>{ ch = 0 }
			${ ch = (ch << 3) | byte(data[p]) - '0' }
			@{ buf.WriteByte(ch) }
			;
		action writeRune {
			//fmt.Println("write rune", string(data[p]))
			buf.WriteRune(data[p])
		}
		singleQuoteEscape := (
			start: (
				sq @write_buf @mark -> quote |
				sl @write_buf -> slash |
				(ASCII - (sl | sq)) -> start |
				notASCII -> start
			),
			slash: (
				escape -> start |
				slashHex -> start |
				slashUnicode -> start |
				slashOctal -> start |
				^(escape | 'x'i | 'u'i | '0'..'7') $writeRune -> start
			) @mark,
			quote: (
				sq -> start |
				not_newline -> waitNewline |
				'\n' -> eatWs |
				^(ws | sq) @done_quote -> final |
				zlen %done_quote -> final
			),
			waitNewline: (
				'\n' -> eatWs |
				^[\n] @done_quote -> final |
				zlen %done_quote -> final
			),
			eatWs: (
				ws -> eatWs |
				sq @mark -> start |
				^(ws | sq) @done_quote -> final |
				zlen %done_quote -> final
			)
			) @err(string_err)
			;

		action singleQuoteNoEscape {
			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			fcall singleQuoteNoEscape;
		}
		action singleQuoteEscape {
			checkUTF8 = true
			isNotASCII = false
			buf.Reset()
			fcall singleQuoteEscape;
		}
		action comment {
			commentDepth = 1
			fcall comment;
		}
		action digit {
			ret = append(ret, string(data[mark:p]))
		}
		top =
			ws |
			'/*' @comment |
			digit+ >mark %digit |
			sq @mark @singleQuoteNoEscape |
			'e' sq @mark @singleQuoteEscape
			;
		main :=
			top** %err{ return nil, errors.Errorf("ERROR at %d of %d", p, len(data)) }
			;

		write init;
		write exec;
	}%%

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
