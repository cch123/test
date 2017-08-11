package main

import (
	"context"
	"fmt"
	"net/http"
)

func testContext(ctx context.Context) {
	if ctx == nil {
		println("nil")
		fmt.Printf("%#v\n", ctx)
		return
	}
	println("not nil")
	fmt.Printf("%#v\n", ctx)
}

func pointerContext(ctx *context.Context) {
	if ctx == nil {
		println("nil")
		fmt.Printf("%#v\n", ctx)
		return
	}
	println("not nil")
	fmt.Printf("%#v\n", ctx)
}

func httpHandler(h *http.Handler) {
	if h == nil {
		println("nil handler")
		fmt.Printf("%#v\n", h)
		return
	}
	println("not nil handler")
	fmt.Printf("%#v\n", h)
}

func interfa(h *interface{}) {
	if h == nil {
		println("nil handler")
		fmt.Printf("%#v\n", h)
		return
	}
	println("not nil handler")
	fmt.Printf("%#v\n", h)
}

func main() {
	testContext(context.TODO())
	testContext(nil)
	var x *context.Context
	pointerContext(x)
	var y *http.Handler
	var z interface{} = y
	httpHandler(y)
	interfa(&z)
}
