package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"io"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w,docgen.JSONRoutesDoc(r))
	})
	http.ListenAndServe(":3000", r)
}
