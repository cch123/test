package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	// if the src parameter is nil, then will auto read the second filepath file
	f, _ := parser.ParseFile(fset, "./example.go", nil, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}

	fmt.Printf("%#v",f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Tag.Value)
	//.(*ast.TypeSpec.Type.(*ast.StructType).Fields[0].(*ast.FieldList).List[0].Tag))
}
