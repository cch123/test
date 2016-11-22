package main

import (
	"fmt"

	"github.com/pingcap/tidb/parser"
)

func main() {
	p := parser.New()
	stmt, _ := p.ParseOneStmt("select * from AA where A=1 and B>1 order by C desc, D asc limit 10,10", "", "")
	fmt.Printf("%#v", stmt)
}
