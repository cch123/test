package main

import "fmt"
import "reflect"

func main() {
	var a = 0.57
	var b = a * 100
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(b)
	fmt.Println(int(b))
}
