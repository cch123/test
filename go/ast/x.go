package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	// if the src parameter is nil, then will auto read the second filepath file
	f, _ := parser.ParseFile(fset, "./example.go", nil, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}

	fmt.Printf("%#v\n", f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Tag.Value)
	str := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Tag.Value
	str = strings.Trim(strings.Trim(str, "\""), "`")
	fmt.Println(str)
	arr := strings.Split(str, " ")
	fmt.Println(strings.Split(str, " "))
	for _, v := range arr {
		fmt.Println(v)
		fmt.Println(strings.Split(v, ":")[0])
		fmt.Println(strings.Split(v, ":")[1])
	}
	//.(*ast.TypeSpec.Type.(*ast.StructType).Fields[0].(*ast.FieldList).List[0].Tag))
}
