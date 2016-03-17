package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
)

type SpecialString interface{}

func main() {
	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
	//output interface{}
	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
	fmt.Println(reflect.TypeOf((*SpecialString)(nil)))
	//output main.SpecialString
	//如果不传指针类型的话，会报panic
	//下面这两个都不行
	//fmt.Println(inject.InterfaceOf(1))
	//fmt.Println(inject.InterfaceOf((*int)(nil)))
}
