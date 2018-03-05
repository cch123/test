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
	/*
			PUT /user/installations/:installation_id/repositories/:repository_id

		GET /search
		GET /status
		GET /support
		GET /user/installations/:installation_id/repositories
		GET /installation/repositories
		GET /repos/:owner/:repo/commits
		GET /repos/:owner/:repo/issues/events
	*/
	//r.GET("", Hello)
	r.PUT("/src/:filename", Hello)
	r.PUT("/src/*abc", Hello)
	//r.PUT("/search", Hello)
	//r.PUT("/status", Hello)
	//r.PUT("/support", Hello)
	//r.PUT("/user/installations/:installation_id/repositories/:reposit/aaa", Hello)
	//r.GET("/user/:info", Hello)

	httprouter.WalkPrint(r, "PUT")
}
