package main

import (
	"reflect"
)

func callOther() float64

func main() {
	x := callOther()
	println(x)
	//println(math.Inf(4))
	x, err := reflect.TypeOf("runtime.g").FieldByName("goid")
	off := x.Offset
}
