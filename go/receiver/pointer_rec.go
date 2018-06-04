package main

import "fmt"

type person struct{ age int }

func (p *person) prin() {
	fmt.Printf("%p\n", p)
	p.age = 2
}

func main() {
	var p person
	// 会自动取地址
	p.prin()
	println(&p)
	fmt.Println(p)
}
