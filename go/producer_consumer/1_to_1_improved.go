package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	tasks := make(chan string)

	// producer
	go func() {
		for i := 0; i < 100; i++ {
			tasks <- "http://eshop.com/user/" + strconv.Itoa(i)
		}
		close(tasks)
	}()

	// worker
	go func() {
		for {
			if url, ok := <-tasks; ok {
				fmt.Println(url)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	fmt.Println(runtime.NumGoroutine())
}
