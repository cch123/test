package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.ValueOf(1).Kind())
	var a = 0.0
	fmt.Println(reflect.ValueOf(a).Kind())
	var b = 0
	fmt.Println(reflect.ValueOf(b).Kind())
}
