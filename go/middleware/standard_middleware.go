package main

import "net/http"

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yes in hook"))
}

func main() {
	http.ListenAndServe(":8080", mainHandler)
}
