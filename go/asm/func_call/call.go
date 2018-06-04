package main

import _ "strconv"

func callOther() string

func main() {
	x := callOther()
	println(x, len(x))
}
