package main

import (
	"fmt"
	"reflect"

	"./a"
)

type A int

func main() {
	var d A
	var f a.A
	fmt.Println("reflect.TypeOf d", reflect.TypeOf(d))
	fmt.Println("reflect.TypeOf f", reflect.TypeOf(f))
	fmt.Println("reflect.TypeOf(d) == reflect.TypeOf(f) ?", reflect.TypeOf(d) == reflect.TypeOf(f))
}
