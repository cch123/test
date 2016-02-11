package main

import "fmt"

type Person struct {
	age    int
	height float32
	weight float32
}

func main() {
	p := new(Person)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", *p)
}
