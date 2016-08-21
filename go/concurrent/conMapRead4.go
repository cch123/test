package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/a/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Println(count)
	})

	http.HandleFunc("/b/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Println(count)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
