package main

import "fmt"

// 注意观察这个函数的 args 和 locals
// 以及生成的汇编的 $x-y 之间的关系
func test(x []int) int {
	var a = 1
	var b = struct{ X int }{X: 93}
	fmt.Println(b)
	return a + b.X
}

func main() {
}
