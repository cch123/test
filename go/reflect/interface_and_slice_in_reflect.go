package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = []int{1, 2, 3}
	var b interface{} = a
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(&a))
	fmt.Println(reflect.TypeOf(&b))

	//c := b.(reflect.TypeOf(a))
	// reflect.TypeOf is not a type
}
