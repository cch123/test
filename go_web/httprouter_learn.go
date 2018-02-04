package main

import (
	"fmt"
	"net/http"

	"github.com/cch123/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	r := httprouter.New()
	//r.GET("/hello/over/:fuc/set/:name/:id/ohyea/", Hello)
	r.GET("/user/:id/get/:addr/abc/:id/", Hello)
	r.GET("/user/:id/gets/:addr/abc/:id/", Hello)
	//r.GET("/hello/:fuc/:name", Hello)
	//	r.GET("/hello/fuci/:name", Hello)
	httprouter.WalkPrint(r, "GET")
}
