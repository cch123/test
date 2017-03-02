// go build -gcflags "-m"
// 可以看到编译器会做哪些内联优化
package main

import "fmt"

func test() {
}

func main() {
	test()
	var a []string
	fmt.Println(len(a))
}
