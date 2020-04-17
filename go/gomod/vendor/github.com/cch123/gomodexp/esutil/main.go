package esutil

import "fmt"
import "github.com/cch123/elasticsql"

func init() {
	fmt.Println("init function in etcdcli")
}

func OhES() {
	fmt.Println("OhETCD")
	elasticsql.Convert("select * from a")
}
