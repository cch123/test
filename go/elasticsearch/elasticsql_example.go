package main

import (
	"fmt"
	"os"

	"github.com/cch123/elasticsql"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("参数数目错误")
		os.Exit(1)
	}
	var sql = os.Args[1]
	dsl, esType, err := elasticsql.Convert(sql)
	fmt.Println(dsl)
	fmt.Println(esType)
	fmt.Println(err)
}
