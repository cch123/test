package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"
)

type person struct {
	X *int `json:"x" form:"x" binding:"required"`
}

func contentType(r *http.Request) string {
	return filterFlags(requestHeader(r, "Content-Type"))
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func requestHeader(r *http.Request, key string) string {
	if values, ok := r.Header[key]; ok {
		return values[0]
	}
	return ""
}

func bind(r *http.Request, dst interface{}) error {
	contentType := contentType(r)
	if len(contentType) == 0 {
		// do some thing?
	}
	fmt.Println(contentType)

	b := binding.Default(r.Method, contentType)
	return b.Bind(r, dst)
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	p := person{}
	err := bind(r, &p)
	fmt.Printf("%#v\n", p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
