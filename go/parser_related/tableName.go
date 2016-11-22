package main

import (
	"fmt"

	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/parser"
)

var sql = `select * from AA left join A_b on (AA.id = A_b.id) where A=1 and B>1 order by C desc, D asc limit 10,10`

func main() {
	p := parser.New()
	stmt, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n\n", stmt)
	selectStmt := stmt.(*ast.SelectStmt)

	fmt.Printf("selectStmt.From.TableRefs.Left is %#v\n", selectStmt.From.TableRefs.Left)
	fmt.Printf("selectStmt.From.TableRefs.Right is %#v\n\n", selectStmt.From.TableRefs.Right)

	tableSource := selectStmt.From.TableRefs.Left.(*ast.TableSource)
	fmt.Printf("left tableSource.Source is %#v\n", tableSource.Source)
	tableName := tableSource.Source.(*ast.TableName)
	fmt.Printf("left tableName.Name.O is %#v\n", tableName.Name.O)
	tableSource = selectStmt.From.TableRefs.Right.(*ast.TableSource)
	fmt.Printf("right tableSource.Source is %#v\n", tableSource.Source)
	tableName = tableSource.Source.(*ast.TableName)
	fmt.Printf("right tableName.Name.O is %#v\n", tableName.Name.O)
}
