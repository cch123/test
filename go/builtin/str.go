package main

func p(s string) {
	panic(1)
}

func main() {
	var s = "hello"
	p(s)
}
