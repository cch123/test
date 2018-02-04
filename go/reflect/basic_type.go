package main

import (
	"fmt"
	"reflect"
)

func main() {
	var vList = []interface{}{1, 1.1, true, uint(1), uintptr(1), complex(1, 1)}
	for _, v := range vList {
		to(v)
	}

	var x = 1
	fmt.Println("basic type can set? ", reflect.ValueOf(x).CanSet())

	// error!! SetInt not addressable
	// reflect.ValueOf(x).SetInt(2)
	// we can use pointer to set basic value
	reflect.ValueOf(&x).Elem().SetInt(2)
	fmt.Println(x)
}

func to(v interface{}) {
	fmt.Printf("val: %v, reflect.TypeOf: %v, reflect.TypeOf.Kind : %v\n", v, reflect.TypeOf(v), reflect.TypeOf(v).Kind())
}
