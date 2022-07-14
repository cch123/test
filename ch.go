package main

import "time"

func main() {
	var ch = make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch = nil
		time.Sleep(time.Second)
		ch = make(chan int)
	}()

	for {
		select {
		case v := <-ch:
			println(v)
		}
	}
}
