package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	jobChan := make(chan struct{})
	ctx, cancelFn := context.WithCancel(context.TODO())
	worker := func() {
		jobLoop:
		for {
			select {
			case <-jobChan:
				// do my job
				fmt.Println("do my job")
			case <-ctx.Done():
				// parent want me to quit
				fmt.Println("parent call me to quit")
				break jobLoop
			}
		}
	}

	go worker()

	cancelFn()
	time.Sleep(time.Second)
}
