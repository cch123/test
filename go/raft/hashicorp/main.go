package main

import "net/http"

func main() {
	config(3)

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
