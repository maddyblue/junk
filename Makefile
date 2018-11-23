test:
	#go generate
	ragel -Z -G2 lex.rl
	go run *.go

dot:
	ragel -Z -V -p lex.rl -o lex.dot
	dot -Tpdf -o lex.pdf lex.dot
