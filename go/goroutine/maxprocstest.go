package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	for {
		fmt.Println("f1")
		time.Sleep(time.Second)
	}
}

func f2() {
	for {
		fmt.Println("f2")
		time.Sleep(time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go f1()
	go f2()
	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}
