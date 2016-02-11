package main

func main() {
	a := []string{}
	println(a)
	println(len(a))
	println(cap(a))
	a = append(a, "fuck")
	println(a)
	println(len(a))
	println(cap(a))
}
