package main

import "context"
import "time"
import _ "net/http/pprof"
import "net/http"

func init() {
	go http.ListenAndServe(":9999", nil)
}

func main() {
	ctx, _ := context.WithCancel(context.TODO())
	for i := 0; i < 100; i++ {
		go func() {
			select {
			case <-ctx.Done():
			}
		}()
	}
	time.Sleep(time.Hour)
}
