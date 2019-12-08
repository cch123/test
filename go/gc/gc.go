package main

import (
	"fmt"
	"runtime"
	"time"
)

type P struct {
	Age int
}

func getPartOfSlice() []*P {
	var s = make([]*P, 0, 10000)
	for i := 0; i < 10000; i++ {
		var p = &P{i}
		runtime.SetFinalizer(p, func(x *P) { println("gc happen on p", x) })
		s = append(s, p)
	}
	return s[100:101]
}

func main() {
	var k = getPartOfSlice()
	// type 1
	// print then gc
	//fmt.Println(k[0])
	//runtime.GC()
	// type 2
	// gc then print
	runtime.GC()
	fmt.Println(k[0])
	time.Sleep(time.Hour)
}
