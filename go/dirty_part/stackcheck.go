// 我对这个demo给出的代码和结论表示怀疑
package main

import (
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	go goFor()
	time.Sleep(time.Second)
	runtime.GC()
	println("gc success passed")
	time.Sleep(time.Second * 3600)
}

func goFor() {
	for {
		callFuncStackCheckTest()
	}
}
func callFuncStackCheckTest() {
	var c1 int64 = 1

	var c2 int64 = 1

	var c3 int64 = 1

	var c4 int64 = 1

	var c5 int64 = 1

	var c6 int64 = 1

	var c7 int64 = 1

	var c8 int64 = 1

	var c9 int64 = 1

	var c10 int64 = 1

	var c11 int64 = 1

	var c12 int64 = 1

	var c13 int64 = 1

	var c14 int64 = 1

	var c15 int64 = 1

	var c16 int64 = 1 //移除该行及下一行注释，，则程序可正常GC，不会夯住
	c1 = c1 + c1 + c2 + c3 + c4 + c5 + c6 + c7 + c8 + c9 + c10 + c11 + c12 + c13 + c14 + c15 + c16
}
