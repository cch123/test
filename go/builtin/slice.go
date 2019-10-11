package main

func p(s []int) {
	panic(1)
}

func main() {
	var s = []int{1, 2, 3}
	p(s)
}
