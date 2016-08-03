package main

import (
	"net/http"
	_ "net/http/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	http.ListenAndServe(":8080", nil)
}
