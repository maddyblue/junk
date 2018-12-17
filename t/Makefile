test:
	ragel -Z lex.rl
	go test *.go

dot:
	ragel -Z -V -p -o lex.dot lex.rl -M singleQuote
	dot -Tpdf -o lex.pdf lex.dot
