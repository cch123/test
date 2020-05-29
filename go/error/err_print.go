package main

import "fmt"

type P struct {
	Age  int
	Name int
}

func (p P) Error() string {
	return "fuck"
}

func main() {
	var err P
	fmt.Printf("%#v\n", err) // main.P{Age:0, Name:0}
	fmt.Printf("%v\n", err) // fuck
	fmt.Printf("%s\n", err) // fuck
}

