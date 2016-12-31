// Package main provides tes
package main

import "fmt"

type A struct {
	a int
}

func main() {
	var m = []A{{1}, {2}}
	test(m...)
}

func test(s ...A) {
	//func test(s ...interface{}) {
	//cannot use m (type []A) as type []interface {} in argument to test
	for k, v := range s {
		fmt.Println(k, v)
	}
}
