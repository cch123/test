package main

import (
	"fmt"
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
	}()

	// worker
	go func() {
		for {
			url := <-tasks
			fmt.Println(url)
		}
	}()
	time.Sleep(time.Second * 3)
}
