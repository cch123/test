package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	ID int
}

var a = map[string]reflect.Type{
	"person": reflect.TypeOf(Person{}),
}

func main() {
	p := reflect.New(a["person"])
	fmt.Println(p)
	fmt.Println(reflect.ValueOf(p).CanSet())
	// will panic fmt.Println(reflect.ValueOf(p).Elem().CanSet())
	fmt.Println(reflect.ValueOf(&p).CanSet())
	fmt.Println(reflect.ValueOf(&p).Elem().CanSet())
	d := p.Elem()
	fmt.Println(d.CanSet())
	fmt.Println(reflect.TypeOf(d).Kind())
	fmt.Println(d.NumField())

	d.Field(0).SetInt(100)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", d)
	// 反射出来的东西也只能通过反射的方式set给别的变量么。。
	var q *Person
	fmt.Println(a["person"].AssignableTo(reflect.TypeOf(q)))
	qElem := reflect.ValueOf(&q).Elem()
	qElem.Set(p)
	fmt.Printf("%#v", q)
}
