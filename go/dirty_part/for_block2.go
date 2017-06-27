package main

import "time"

func main() {
	var ch = make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- 1
			time.Sleep(time.Second)
		}
	}()

	for {
		// the wrong part
		if len(ch) == 100 {
			sum := 0
			itemNum := len(ch)
			for i := 0; i < itemNum; i++ {
				sum += <-ch
			}
			if sum == itemNum {
				return
			}
		}
	}

}
