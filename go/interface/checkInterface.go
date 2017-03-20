package main

import "io"

type Person struct {
	age int
}

func (p Person) Write(input []byte) (int, error) {
	println(1)
	return 1, nil
}

func main() {
	// check whether person implements the io.Writer interface
	// but can only be done while compile time
	// so will often appear on the beginning of the struct definition
	// to ensure you implement the interface
	// similar to implements

	var _ io.Writer = Person{}
}
