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
	for i := 0; i < 15; i++ {
		fmt.Println(limiter.TakeAvailable(1))
	}
}
