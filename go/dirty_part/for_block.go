package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go server()
	go printNum()
	var i = 1
	for {
		// will block here, and never go out
		i++
	}
	fmt.Println("for loop end")
	//fmt.Println(i)
	time.Sleep(time.Second * 3600)
}

func printNum() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func server() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
