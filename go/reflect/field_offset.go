package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	A struct {
		Age    int64
		Height int64
		a      *int64
	}
	low int64
	aa  int
	dd  int
}

func main() {
	var p Person
	t, err := reflect.TypeOf(p).FieldByName("aa")
	fmt.Println(t, err)
	println(t.Offset)
}
