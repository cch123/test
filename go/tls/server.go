package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write(make([]byte, 16380))
	})

	err := http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
