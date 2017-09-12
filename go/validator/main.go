package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Person struct {
	Age int `json:"age" validate:"eq=10|eq=1"`
	// len can not be >, can only=
	// 值等于 abc
	Name string `json:"name" validate:"eq=abc"`
	// 不等于空字符串
	Proto string `json:"proto" validate:"ne="`
}

func main() {
	var p Person
	p.Age = 1
	p.Name = "abc"
	validator := validator.New()
	err := validator.Struct(p)
	fmt.Println(err)

}
