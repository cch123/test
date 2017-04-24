package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/gops/agent"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	// 标准库没有办法在写入后，再从responsewriter里读出来
	// 蛋疼
	wr.WriteHeader(404)
	io.WriteString(wr, "hello")
	fmt.Println(wr.Header())
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
