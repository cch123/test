package main

// interface 的type和value同时为nil，== nil才成立
// golang的经典神坑

import "fmt"
import "reflect"

type Obj struct {
	One *Obj
}

func main() {
	var o = &Obj{}
	var arr = []interface{}{o.One}
	fmt.Println(arr[0])
	fmt.Println(arr[0] == nil)
	fmt.Printf("%T: %#v\n", arr[0], arr[0])
	fmt.Println(reflect.ValueOf(arr[0]).IsNil())
}
