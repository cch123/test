package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	age int
}

func main() {
	var a Person
	fmt.Println(reflect.TypeOf(a))
}
