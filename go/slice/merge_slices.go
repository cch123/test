package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	merged := append([]int{}, a[:5]...)
	merged = append(merged, a[6:]...)
	fmt.Println(merged)
}
