package main

import "net/http"

type mainHandler struct{}

func (m mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func hookHandler(next http.Handler) http.Handler {
	//是HandlerFunc而不是HandleFunc，注意。。
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("in hook"))
		next.ServeHTTP(w, r)
		w.Write([]byte("in hook after"))
	})
}

func main() {
	var a mainHandler
	http.ListenAndServe(":8080", hookHandler(a))
}
