package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func main() {
	var x = a[1]
	b := (*person)(unsafe.Pointer((&x)))
	fmt.Println(b)
}

type person struct {
	age    byte
	age1   byte
	age2   byte
	age3   byte
	height int32
	width  int32
}

var a = map[int][12]byte{1: [12]byte{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x = a[1]
		b := (*person)(unsafe.Pointer((&x)))
		_ = b.age
	}
}
