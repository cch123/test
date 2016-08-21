package main

import (
	"fmt"
	"net/http"
	"time"
)

type Foo struct {
	content string
}

type FooSlice []*Foo

func updateFooSlice(fooSlice FooSlice) {
	for {
		foo := &Foo{content: "new"}
		fooSlice[0] = foo
		time.Sleep(time.Second)
	}
}

func installHttpHandler(fooSlice FooSlice) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		for _, foo := range fooSlice {
			if foo != nil {
				fmt.Fprintf(w, "foo: %v ", (*foo).content)
			}
		}

	}
	http.HandleFunc("/", handler)
}

func main() {
	foo1 := &Foo{content: "hey"}
	foo2 := &Foo{content: "yo"}
	fooSlice := FooSlice{foo1, foo2}

	installHttpHandler(fooSlice)

	go updateFooSlice(fooSlice)

	http.ListenAndServe(":8080", nil)
}
