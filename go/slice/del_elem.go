package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a, "cap", cap(a), "len", len(a))
	a = append(a[:5], a[6:]...)
	fmt.Println(a, "cap", cap(a), "len", len(a))
}
