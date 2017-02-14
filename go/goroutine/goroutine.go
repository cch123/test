package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
