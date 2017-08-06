package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	stopCh := make(chan bool)
	cnt := 0
	go func() {
		for {
			time.Sleep(time.Second)
			ch <- "a"
			cnt++
			if cnt == 3 {
				close(ch)
				stopCh <- true
				break
			}
		}
	}()

	//loop:
	for {
		select {
		case msg, ok := <-ch:
			fmt.Println(msg, ok)
		case <-stopCh:
			//try add the following line
			//break loop
		}
	}
}
