package main

import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

var (
	v interface{}
	r io.Reader
	f *os.File
)

func main() {
	fmt.Println(v == nil)
	fmt.Println(r == nil)
	fmt.Println(f == nil)
	v = r
	fmt.Println(v == nil)
	v = f
	fmt.Println(v == nil)
	r = f
	fmt.Println(r == nil)

	type a struct {
		E struct{}
		V int32
	}

	var x a
	fmt.Println(unsafe.Sizeof(x))
}
