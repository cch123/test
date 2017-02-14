package main

import "runtime"
import "time"

func main() {

	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i)
	}
	time.Sleep(2 * time.Second)
}
