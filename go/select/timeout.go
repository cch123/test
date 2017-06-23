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
	timeoutChan := time.After(3 * time.Second)
	dataChan := make(chan int, 100)
	needReadSignalChannel := make(chan struct{})
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
			if len(dataChan) > 50 {
				needReadSignalChannel <- struct{}{}
			}
		}
	}()

	// should be in the consumer goroutine
	for {
		select {
		case <-timeoutChan:
			readChannel(dataChan)
			timeoutChan = time.After(3 * time.Second)
			println("timeout")
		case <-needReadSignalChannel:
			readChannel(dataChan)
		}
	}

}
