package main

import (
	"fmt"
	"reflect"
)

type aaa string

func main() {
	var a aaa
	a = aaa("oh no")
	fmt.Println(a)
	var b string
	b = string(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
