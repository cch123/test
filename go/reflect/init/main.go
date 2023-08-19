// 接受任意的 interface
// 返回非空的 interface

package main

import (
	"fmt"
	"reflect"
)

func convertToNonEmpty(input any) any {
	fmt.Println("input == nil", input == nil, "input value is nil", reflect.ValueOf(input).IsNil())

	var result any
	if input == nil || reflect.ValueOf(input).IsNil() {
		// init the input and return
		result = reflect.New(reflect.TypeOf(input).Elem()).Interface()
	}

	return result
}

type Person struct {
	Age  int
	Name string
}

func main() {
	var p *Person
	convertToNonEmpty(p)
	fmt.Printf("before: %#v; after: %#v; deref: %#v", p, convertToNonEmpty(p), *convertToNonEmpty(p).(*Person))
}
