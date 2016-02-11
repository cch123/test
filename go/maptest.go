package main

import "fmt"

func main() {
	var m map[string]int
	//下面这句是必须的
	m = make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Printf("%+v", m)
	delete(m, "x")
	delete(m, "one")

	ma := make(map[string]int)
	ma["x"] = 1
	fmt.Printf("%+v", ma)
}
