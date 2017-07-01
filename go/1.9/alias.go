package main

type myint int

func main() {
	var a myint
	var b = 1
	// will cause error below 1.9
	a = b
}
