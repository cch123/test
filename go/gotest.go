package main

import "fmt"

import "runtime"

//import "time"

func say(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
		//runtime.Gosched()
		//time.Sleep(time.Millisecond)
	}
}

func main() {
	a := ""
	go func(str string) {
		for i := 0; i < 5; i++ {
			println(str)
		}
	}(a)
	runtime.Gosched()

	//say("x")
}
