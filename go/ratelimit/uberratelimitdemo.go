package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(10) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		//每次调用前take一下就好了，有时间看看实现~
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
