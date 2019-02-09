package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, "", src, parser.Mode(parser.ParseComments))
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Print(fset, fileAST)
}

var src = `
package main

import (
	"fmt"
	"github.com/cch123/elasticsql"
)

func help() {
	println("hello world")
}

func main() {
	var a = 1 // ast.DeclStmt
	fmt.Println(a) // ast.ExprStmt
	var b = 1
	c := a + b // ast.ExprStmt
	println(c) // ast.ExprStmt
}
`
