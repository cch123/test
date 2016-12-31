package main

import (
	"fmt"
	"runtime"
)

func test() {
	if pc, _, _, ok := runtime.Caller(0); ok {
		f := runtime.FuncForPC(pc)
		fmt.Println(f.Name())
		if f.Name() != "main.test" {
			println("inline")
		}
	}
}

func main() {
	test()
}
