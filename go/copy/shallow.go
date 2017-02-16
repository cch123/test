package main

import "fmt"

type P struct {
	age *int
}

func main() {
	var a *P = new(P)
	a.age = new(int)
	*a.age = 1
	fmt.Printf("a is %#v\n", a)
	fmt.Printf("a ptr is %p\n", a)
	fmt.Printf("*a.age is %#v\n", *a.age)
	fmt.Printf("a.age is %#p\n", a.age)

	var b *P = new(P)
	// like C and C 艹, = is shallow copy
	*b = *a
	fmt.Printf("b is %#v\n", b)
	fmt.Printf("b ptr is %p\n", b)
	fmt.Printf("*b.age is %#v\n", *b.age)
	fmt.Printf("a.age is %#p\n", b.age)

}
