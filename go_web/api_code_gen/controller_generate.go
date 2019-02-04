package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

var (
	filePath = flag.String("f", "", "file path")
)

func getFields(node *ast.GenDecl) {
	if node.Specs == nil || len(node.Specs) == 0 {
		fmt.Println("Invalid spec, length must > 0")
		return
	}

	spec, ok := node.Specs[0].(*ast.TypeSpec)
	if !ok {
		fmt.Println("Invalid spec, expected ast.TypeSpec")
		return
	}

	structType, ok := spec.Type.(*ast.StructType)
	if !ok {
		fmt.Println("Invalid type, expect struct type")
		return
	}

	if structType.Fields == nil {
		fmt.Println("Empty field list , skip")
		return
	}

	fieldList := structType.Fields.List
	for _, field := range fieldList {
		// may panic
		fieldTypeStr := field.Type.(*ast.Ident).Name

		fieldName := field.Names[0].Name
		tagStr := field.Tag.Value
		fmt.Println(fieldTypeStr, fieldName, tagStr)
		getTagDefMap(tagStr)
	}

}

func getTagDefMap(tagStr string) map[string]string {
	var res = make(map[string]string)
	tagStr = strings.Replace(tagStr, "`", "", -1)
	tagStr = strings.Replace(tagStr, "\t", "", -1)

	splitedArr := strings.Split(tagStr, " ")
	filteredArr := []string{}

	for _, v := range splitedArr {
		v = strings.Replace(v, "\"", "", -1)

		// 过滤掉空字符串
		if len(v) == 0 || !strings.Contains(v, ":") {
			continue
		}

		splitedKeyAndVal := strings.Split(v, ":")
		if len(splitedKeyAndVal) < 2 {
			continue
		}

		// form : driver_id
		res[splitedKeyAndVal[0]] = splitedKeyAndVal[1]

		filteredArr = append(filteredArr, v)
	}

	//fmt.Println(filteredArr, len(filteredArr))
	fmt.Println(res)
	return res
}

func parse(filename string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(4))
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok != token.TYPE {
				fmt.Println("Not type declarition, skip...")
				return false
			}

			// is type declare
			if x.Doc != nil {
				commentStr := x.Doc.List[0].Text
				fmt.Println(commentStr)
				// extract method
				// extract request url
			}

			// get all fields and tags
			getFields(x)

			//ast.Print(fset, x)
			return false
		default:
			fmt.Println("Not type declarition, skip...")
		}
		return true
	})
	fmt.Println(err)
}

func main() {
	flag.Parse()
	if *filePath == "" {
		flag.Usage()
		return
	}

	contentBytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(contentBytes))
	fmt.Println(err)

	parse(*filePath)

}
