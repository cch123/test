package main

import "sync"

func main() {
	ch := make(chan string, 1000)
	go func() {
		for i := 0; i < 10000000; i++ {
			ch <- "json encode this is the last of the message laneghete is very long everytttt"
		}
		close(ch)
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var ok bool
			var e interface{}
			for {
				e, ok = <-ch
				if !ok {
					break
				}
				_ = e
			}
		}()
	}
	wg.Wait()
}
