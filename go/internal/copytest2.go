package main

import "fmt"

type integer int

func (i integer) printAge() {
	fmt.Printf("%p\n", &i)
}

func main() {
	i := integer(1)
	fmt.Printf("%p\n", &i)
	i.printAge()
}
