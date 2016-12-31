// go build -gcflags "-m"
// 可以看到编译器会做哪些内联优化
package main

func test() {
}
func main() {
	test()
}
