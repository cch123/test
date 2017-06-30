package main

import (
	"sync"

	"github.com/skoo87/ringbuffer"
)

func main() {
	ring := ringbuffer.NewRing(100, 1000)
	go func() {
		var wbuf *ringbuffer.Buffer
		for i := 0; i < 10000000; i++ {
			wbuf = ring.Write(wbuf, "json encode this is the last of the message laneghete is very long everytttt")
		}
		ring.Stop(wbuf)
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var rbuf *ringbuffer.Buffer
			var e interface{}
			for {
				if e, rbuf = ring.Read(rbuf); rbuf == nil {
					break
				}
				_ = e
				//fmt.Println(e.(int))
			}
		}()
	}
	wg.Wait()

}
