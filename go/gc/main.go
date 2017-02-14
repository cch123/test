package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	go func() {
		for i := 0; i < 1000000; i++ {
			time.Sleep(time.Second)
			c <- i
		}
	}()
	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case <-time.After(10 * time.Second):
			fmt.Println("timeout!")
		}
	}
}
