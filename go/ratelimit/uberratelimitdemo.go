package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

/**
理论上漏桶内的水滴增加速率是随意的，
只是对出水(take)速率有控制，
用于流量整形，即使有突发流量，也会平滑(让你等待)地进行处理
不过感觉uber的实现不太好啊，Take直接阻塞，
没办法控制take成功与否之后的逻辑
*/
func main() {
	rl := ratelimit.New(10) // per second

	prev := time.Now()
	time.Sleep(time.Second * 4)
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		//每次调用前take一下就好了，有时间看看实现~
		go func() {
			now := rl.Take()
			fmt.Println(i, now.Sub(prev))
			prev = now
		}()
	}
	<-ch
}
