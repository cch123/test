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

	// interface pointer to a
	var iter interface{}
	iter = &a
	// slice 的指针需要先用 Elem
	// ValueOf(slicePointer).Elem().Type
	rval := reflect.MakeSlice(reflect.ValueOf(iter).Elem().Type(), 10, 10)
	reflect.ValueOf(iter).Elem().Set(rval)
	fmt.Println(iter)
	fmt.Println(a)

	// slice 的元素类型可以直接这么返回
	fmt.Println(reflect.TypeOf(a).Elem())

	// slice 指针就比较麻烦了
	fmt.Println(reflect.ValueOf(&a).Elem().Type().Elem())
}
