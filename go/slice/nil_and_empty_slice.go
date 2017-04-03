package main

import "fmt"

func main() {
	var a []int
	var b = make([]int, 0)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
