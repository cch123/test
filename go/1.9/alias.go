package main

import "fmt"
import "reflect"

// 上面这个是 type alias，两者实际上完全一致，在类型 switch 的时候只能有一个分支
type myint = int

// 这个是定义了一个新类型，赋值需要做强制转换
type mmint int

func main() {
	var a myint
	var b = 1

	// will cause error below 1.9
	a = b
	fmt.Println(a)

	var c mmint
	// ./alias.go:21:4: cannot use b (type int) as type mmint in assignment
	// c = b
	c = 1
	fmt.Println(c)
	// main.mmint
	fmt.Println(reflect.TypeOf(c))
	// int
	fmt.Println(reflect.TypeOf(a))
}
