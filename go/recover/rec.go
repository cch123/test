// recover 要点：
// recover 的 panic 必须崩溃在当前函数的上下文内，或者当前函数调用对象的上下文内
// 也就是说
// a() {
//    b()
//    c()
// }
// 这种情况，在b中recover 不了c
// 在c中也recover不了b
package main

import "fmt"

func main() {
	t()
	p()
}

func p() {
	panic("abc")
}

func t() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
}
