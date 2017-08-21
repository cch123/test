package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"

	"net/http"
)

func sayhello(httpRespWriter http.ResponseWriter, request *http.Request) {
	panic("normal fuck")
}

func sayhello2(httpRespWriter http.ResponseWriter, request *http.Request) {
	go func() {
		panic("deep fuck")
	}()
}

func main() {
	http.HandleFunc("/", sayhello)
	http.HandleFunc("/a", sayhello2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
