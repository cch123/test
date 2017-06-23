package main

import (
	"fmt"
	"time"
)

func readChannel(dataChan chan int) {
	for i := 0; i < len(dataChan); i++ {
		fmt.Printf("%d ", <-dataChan)
	}
	fmt.Println()
}

func main() {
	dataChan := make(chan int, 100)
	minNeedReadAmountNotifyChan := make(chan struct{})
	tickChan := time.Tick(3 * time.Second)
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
			if len(dataChan) > 50 {
				minNeedReadAmountNotifyChan <- struct{}{}
			}
		}
	}()

	// should be in the consumer goroutine
	for {
		select {
		case <-tickChan:
			readChannel(dataChan)
			// 如果直接在 case 上写 time.After 的话
			// 实际上每次 select 都会生成新的 timer
			println("timeout")
		case <-minNeedReadAmountNotifyChan:
			readChannel(dataChan)
		}
	}

}
