package main

import "runtime"

func say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		println(str)
	}
}
func main() {
	go say("hello")
	say("yess")
}
