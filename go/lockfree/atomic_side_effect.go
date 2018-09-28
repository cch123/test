package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var doublebuffer = struct {
	confArr    [2]config
	currentIdx int64
}{
	confArr: [2]config{
		config{m: map[int]int{}},
		config{m: map[int]int{}},
	},
	currentIdx: 1,
}

type config struct {
	m map[int]int
}

func main() {
	go func() {
		for range time.Tick(time.Millisecond) {
			curIdx := atomic.LoadInt64(&doublebuffer.currentIdx)
			newIdx := 1 - curIdx
			// save new val to newIdx config
			doublebuffer.confArr[newIdx].m = map[int]int{43: rand.Intn(100)}
			atomic.StoreInt64(&doublebuffer.currentIdx, newIdx)
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				idx := atomic.LoadInt64(&doublebuffer.currentIdx)
				println(doublebuffer.confArr[idx].m[rand.Intn(100)])
			}
		}()
	}

	time.Sleep(time.Hour)

}
