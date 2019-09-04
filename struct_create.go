package main

import "fmt"

type Foo struct {
	B string
}
type Bar struct {
	X int
}
type Some struct {
	Foo
	Bar
}

func main() {
	var x = Some{
		Foo{B: "abc"},
		Bar{X: 1},
	}
	fmt.Println(x)
}
