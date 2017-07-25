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

func main() {
	var t T
	t.age = 99
	t.name = "seth"
	fmt.Println(t.getName())
	fmt.Println(t.getAge())
}
