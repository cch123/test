package main

import "reflect"
import "fmt"

func main() {
	var i = 1
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(i).Type())
	fmt.Println(reflect.ValueOf(i).Kind())
	fmt.Println(reflect.ValueOf(i).Int())
}
