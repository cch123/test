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
		for url := range tasks {
			fmt.Println(url)
		}
	}()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	time.Sleep(time.Second * 3)
	fmt.Println(runtime.NumGoroutine())
}
