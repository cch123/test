package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	cnt := 0
	go func() {
		for {
			time.Sleep(time.Second)
			ch <- 1
			cnt++
			if cnt == 3 {
				close(ch)
				break
			}
		}
	}()

	for {
		select {
		case msg, ok := <-ch:
			fmt.Println(msg, ok)
		}
	}
}
