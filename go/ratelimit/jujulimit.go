package main

import (
	"fmt"
	"time"

	"github.com/juju/ratelimit"
)

//var totalLimiter *ratelimit.Bucket
//var apiLimiters map[string]*ratelimit.Bucket

func main() {
	var limiter *ratelimit.Bucket
	limiter = ratelimit.NewBucket(time.Second, 10)
	for i := 0; i < 100; i++ {
		fmt.Println(limiter.TakeAvailable(1))
		if i%10 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
