package main

import "fmt"

func testSlice(sl []int) {
	for i := 0; i < 1000; i++ {
		sl = append(sl, i)
	}
	fmt.Println(sl)
}

func main() {
	var a = []int{}
	testSlice(a)
	fmt.Println(a)
}
