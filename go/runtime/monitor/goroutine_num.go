package main

import (
	"fmt"
	"runtime"
	"time"
)

func makeGoRoutine() {
	for {
		time.Sleep(time.Millisecond)
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
}

func main() {
	tick := time.NewTicker(time.Second)
	go makeGoRoutine()

	for {
		select {
		case <-tick.C:
			fmt.Println(int64(runtime.NumGoroutine()))
		}
	}
}
