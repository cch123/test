package main

import (
	"net/http"
	"runtime/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {
	for i := 0; i < 100000; i++ {
		go f()
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}
