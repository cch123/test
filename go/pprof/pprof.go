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
	for i := 0; i < 100000; i++ {
		go f()
	}

	http.ListenAndServe(":11181", nil)
}
