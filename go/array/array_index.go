package main

import (
	"fmt"
	"reflect"
)

func main() {
	var id int32 = 1
	var id2 int64 = 9
	var a = make([]int, 10)
	println(a[id])
	println(a[id2])
	fmt.Println(reflect.TypeOf(id))
	fmt.Println(reflect.TypeOf(id2))
	fmt.Printf("%#v", "a")
}
