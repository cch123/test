package main

import (
	"fmt"
)

func t(a *int) {
}

func main() {
	var a [1]int
	c := a[:]
	fmt.Println(c)
	var b int
	t(&b)
}
