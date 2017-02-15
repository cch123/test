package main

import (
	"io"
	"log"
	"net/http"

	"github.com/google/gops/agent"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	io.WriteString(wr, "hello")
}

func main() {
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
