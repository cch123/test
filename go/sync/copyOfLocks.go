package main

import (
	"fmt"
	"sync"
)

func main() {
	//var muA []sync.Mutex = make([]sync.Mutex, 5)
	var muA [5]sync.Mutex
	muB := muA
	//copy(muB, muA[:])
	fmt.Printf("%p\n", &muB)
	fmt.Printf("%p\n", &muA)

	for k := range muA {
		fmt.Printf("%p %p\n", &muA[k], &muB[k])
	}

	var muC [5]sync.Mutex
	muD := muC[:]
	fmt.Printf("%p\n", &muC)
	fmt.Printf("%p\n", &muD)

	for k := range muA {
		fmt.Printf("%p %p\n", &muC[k], &muD[k])
	}

	var mmuA [5][5]sync.Mutex
	mmuB := mmuA
	mmuA = mmuB
	mmuSlice := mmuA[:]
	fmt.Printf("%#v\n", mmuB)
	fmt.Printf("%#v\n", mmuA)
	fmt.Printf("%#v\n", mmuSlice)

	var fmuA [5][][5]sync.Mutex
	fmuB := fmuA
	fmuA = fmuB
	fmuSlice := fmuA[:]
	fmt.Printf("%#v", fmuA)
	fmt.Printf("%#v", fmuB)
	fmt.Printf("%#v", fmuSlice)
}
