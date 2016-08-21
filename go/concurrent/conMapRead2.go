package main

import (
	"fmt"
	"net/http"
)

//var apiCounts = make(map[string]map[string]int)
var publicMap = make(map[string]int)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	publicMap["a"] = 1
	fmt.Println(publicMap, "a")
}

func handleProxy2(w http.ResponseWriter, r *http.Request) {
	_ = publicMap["a"]
	fmt.Println(publicMap, "b")
}

func main() {
	http.HandleFunc("/", handleProxy)
	http.HandleFunc("/a", handleProxy2)
	http.ListenAndServe(":8080", nil)
}
