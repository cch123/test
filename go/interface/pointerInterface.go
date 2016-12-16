package main

import (
	"fmt"
)

type Human interface {
	Name() string
	Speak()
}

type Person struct {
	name string
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Speak() {
	fmt.Println("fuc")
}

func printHuman(h Human) {
	fmt.Println(h.Name())
	h.Speak()
}

func main() {
	var p Person = Person{"alex"}
	printHuman(&p)
	//printHuman(p) error, it is person pointer who implement the interface
}
