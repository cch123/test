package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int
	var b []int
	reflect.ValueOf(&a).Elem().Set(reflect.MakeSlice(reflect.TypeOf(a), 0, 10))
	reflect.ValueOf(&b).Elem().Set(reflect.ValueOf([]int{1, 2}))
	reflect.ValueOf(&a).Elem().Set(reflect.Append(reflect.ValueOf(a), reflect.ValueOf(1)))
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
}
