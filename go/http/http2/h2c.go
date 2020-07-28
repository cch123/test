package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	h2s := &http2.Server{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello") //"Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	})

	server := &http.Server{
		Addr:    "0.0.0.0:10002",
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:10002]...\n")
	server.ListenAndServe()
}
