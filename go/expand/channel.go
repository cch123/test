package main

import "fmt"

func testChan(ch chan int) {
	ch = make(chan int, 10)
	fmt.Println("ch == nil : ", ch == nil)
}
func main() {
	var ch chan int
	testChan(ch)
	fmt.Println("ch == nil : ", ch == nil)
}
