package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	ch := make(chan struct{}, 100)
	go func() {
		for {
			ch <- struct{}{}
		}
	}()

	for i := 0; i < 5000; i++ {
		go func() {
			for {
				select {
				case _ = <-ch:
				}
			}
		}()
	}

	http.ListenAndServe(":8080", nil)

}
