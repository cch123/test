package main

import "github.com/go-martini/martini"

func hello() string {
	return "Hello world"
}

func main() {
	m := martini.Classic()
	m.Get("/", hello)
	m.Run()
}
