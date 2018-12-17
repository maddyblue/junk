package main

import "fmt"

func main() {
	lexSQL([]byte(`''`))
	lexSQL([]byte(`''''`))
	lexSQL([]byte(`'testing'`))
	lexSQL([]byte(`'test'
		'ing'`))
	lexSQL([]byte(`'test'
		'ing' 'onlyafter'`))
	lexSQL([]byte(`'onlybefore' 'test'
		'ing'`))
	lexSQL([]byte(`'before' 'test'
		'ing' 'after'`))

	fmt.Println("\nEXPECT UNTERMINATED")

	lexSQL([]byte(`'`))
	lexSQL([]byte(`'''`))

	return

	lexSQL([]byte(`/* */`))
	lexSQL([]byte(`/* */ /*   */`))
	lexSQL([]byte(`/* /* */   */`))
	lexSQL([]byte(`/* /* */  /* /* */ */ */`))
	lexSQL([]byte(`  `))

	fmt.Println("\nEXPECT UNTERMINATED")

	lexSQL([]byte(`/*`))
	lexSQL([]byte(`/*/`))
	lexSQL([]byte(`/* /* */   * /`))

	fmt.Println("\nEXPECT ERROR")

	lexSQL([]byte(`/* */* */`))
}
