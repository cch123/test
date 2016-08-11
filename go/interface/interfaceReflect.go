package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{} = int64(58)
	fmt.Println(reflect.TypeOf(val))
	val = 50
	fmt.Println(reflect.TypeOf(val))
}
