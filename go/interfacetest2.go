package main

import (
	"fmt"
	"reflect"
)

//interface{}的slice
type powerful_list []interface{}

//定义element为空interface，类似void指针
type Element interface{}

//存储Element的slice
type more_powerful_list []Element
type Person struct {
	age    int
	weight float32
}

func (p Person) String() string {
	//str := fmt.Sprint("age:", p.age, ";weight:", p.weight)
	str := fmt.Sprintf("age: %d, weight: %.3f", p.age, p.weight)
	return str
}

func main() {
	li := make(powerful_list, 3)
	li[0] = 1
	li[1] = "fuck"
	li[2] = Person{12, 1222}
	//空interface数组里可以放各种各样的类型
	for idx, elem := range li {
		switch value := elem.(type) {
		case int:
			println(idx, value)
			println(idx, elem)
			fmt.Println(idx, reflect.TypeOf(elem))
			fmt.Println(idx, reflect.ValueOf(elem))
		case string:
			println(idx, value)
			println(idx, elem)
			fmt.Println(idx, reflect.TypeOf(elem))
			fmt.Println(idx, reflect.ValueOf(elem))
		case Person:
			//println(value) error重写String方法不影响内置的println，只能用fmt包里的print方法们
			fmt.Println(idx, value)
			fmt.Println(idx, elem)
		}
	}
	var x interface{}
	println(x)
	x = 2
	println(x)
}
