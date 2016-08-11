package main

import (
	"bytes"
	"fmt"
)

type data struct{}

func (this *data) Error() string { return "" }

func test() interface{} {
	var p *data = nil
	return p
}

func test2() interface{} {
	var p *interface{} = nil
	//p = nil
	return p
}

func test3() interface{} {
	var p interface{} = nil
	return p
}

func main() {
	//var a *bytes.Buffer = new(bytes.Buffer)
	var a *bytes.Buffer
	var b interface{} = a
	fmt.Println(a, b)
	fmt.Println(a == nil, b == nil)
	fmt.Printf("ptr a %p\n", a)
	fmt.Printf("ptr b %p\n", b)
	fmt.Println("test() return nil ?", test() == nil)
	fmt.Println("test2() return nil ?", test2() == nil)
	fmt.Println("test3() return nil ?", test3() == nil)
}
