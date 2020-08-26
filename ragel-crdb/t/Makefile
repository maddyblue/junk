test:
	ragel -Z lex.rl
	go test

dot:
	ragel -Z -V -p -o lex.dot lex.rl -M singleQuoteEscape
	dot -Tpdf -o lex.pdf lex.dot
