package main

import "fmt"

type person struct {
	age int
}

func (p person) printAge() {
	fmt.Printf("%p\n", &p)
}

func main() {
	p := person{age: 10}
	fmt.Printf("%p\n", &p)
	p.printAge()
}
