package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"unsafe"
)

var a *map[int]int

func init() {
	b := make(map[int]int)
	a = &b
	for i := 0; i < 1000; i++ {
		(*a)[i] = rand.Intn(1000)
	}
}

func T() {
	go func() {
		for {
			_ = (*a)[10]
		}
	}()

	for i := 0; i < 1000; i++ {
		replaceGlobalMap()
	}

	fmt.Println(a)
}

func replaceGlobalMap() {
	tmpA := map[int]int{}
	for i := 0; i < 1000; i++ {
		tmpA[i] = rand.Intn(1000)
	}
	tmpAPtr := unsafe.Pointer(&tmpA)
	atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(a)), unsafe.Pointer(tmpAPtr))
	println(a, tmpAPtr)
}
