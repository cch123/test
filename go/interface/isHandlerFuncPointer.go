package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func getHandler() http.Handler {
	var a = map[string]http.HandlerFunc{}
	h := a["abc"]
	return h
}

func main() {
	h := getHandler()
	fmt.Println(h)
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.ValueOf(h))
	fmt.Println(h == nil)

	var d interface{} = http.Handler(nil)
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.ValueOf(d))
	fmt.Println(d == nil)

	var f = map[string]http.Handler{}
	g := f["def"]
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.ValueOf(g))
	fmt.Println(g == nil)
}
