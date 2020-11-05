package main

import (
	"fmt"

	atomicex "go.uber.org/atomic"
)

func main() {
	var a atomicex.Uint32
	fmt.Println(a.Load())
	a.Store(23)
	fmt.Println(a.Load())
}
