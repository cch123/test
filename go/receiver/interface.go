package main

import "fmt"

type T struct {
	name string
	age  int
}

func (t *T) getName() string {
	return t.name
}

func (t T) getAge() int {
	return t.age
}

type Name interface {
	getName() string
}

type Age interface {
	getAge() int
}

func main() {
	var ag Age = &T{age: 11, name: "alex"}
	fmt.Println(ag.getAge())

	var na Name = &T{age: 11, name: "xa"}
	fmt.Println(na.getName())

	var ag2 Age = T{age: 11, name: "alex"}
	fmt.Println(ag2.getAge())

	var na2 Name = T{age: 11, name: "xa"}
	fmt.Println(na2.getName())
}
