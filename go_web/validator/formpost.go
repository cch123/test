// schema库会在内部绑定的时候把错误返回来
// 如果自己用，感觉需要一些定制
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type Person struct {
	Name   string
	Phone  int64
	Height []float64
	Salary float64
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Handle error
	}

	var person Person

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&person, r.PostForm)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", person)
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
