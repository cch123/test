package main

import (
	_ "runtime"
	_ "time"
	_ "unsafe"
)

//go:linkname throw runtime.throw
func throw(s string)

//go:nosplit 应该是没啥用

func main() {
	throw("abc")

	// cannot reach here
	println("ooo")
}
