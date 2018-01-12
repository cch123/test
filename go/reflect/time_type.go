package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var a time.Time
	var struc struct{}
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(time.Time{}))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(struc))
}
