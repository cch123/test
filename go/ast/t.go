package main

import (
	"fmt"
	"go/parser"
)

func main() {
	expr, _ := parser.ParseExpr("a * -1")
	fmt.Printf("%#v", expr)

}
