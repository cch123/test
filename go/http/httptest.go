package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("oh yeah")
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
