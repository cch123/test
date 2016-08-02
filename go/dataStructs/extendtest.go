package main

import "fmt"

type Human struct {
	age    int
	weight float32
	height float32
	name   string
}

type Man struct {
	Human
	description string
}

func main() {
	base := Human{11, 12, 11, "name"}
	a := Man{Human: Human{11, 12.1, 12, "Xargin"}, description: "this is a sb"}
	//b := Man{11, 12, 12, "sdfsd", "sdfsdf"} error!
	fmt.Printf("%+v\n", base)
	fmt.Printf("%+v\n", a)
	b := Man{}
	c := new(Man)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
}
