package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var content = make([]byte, 16000)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	wr.Header()["Content-Length"] = []string{fmt.Sprint(len(content))}
	wr.Header()["Content-Type"] = []string{"application/json"}
	wr.Write(content)
}

func main() {
	go func() {
		http.ListenAndServe(":3333", nil)
	}()
	http.HandleFunc("/", sayhello)

	err := http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
