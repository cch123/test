package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	idx := suffixarray.New([]byte("abcdbcbcbcefg"))
	res := idx.Lookup([]byte("bc"), -1)
	fmt.Println(res)
}
