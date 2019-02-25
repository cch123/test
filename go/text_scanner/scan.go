package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

var src = `
where a = 1 and b != 2.1
`

func main() {
	var x scanner.Scanner
	x.Init(strings.NewReader(src))
	for tok := x.Scan(); tok != scanner.EOF; tok = x.Scan() {
		fmt.Println(tok, x.TokenText(), scanner.TokenString(tok))
		/*
		   -2 where Ident
		   -2 a Ident
		   61 = "="
		   -3 1 Int
		   -2 and Ident
		   -2 b Ident
		   33 ! "!"
		   61 = "="
		   -4 2.1 Float
		*/
	}
}
