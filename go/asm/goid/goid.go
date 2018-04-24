package main

import "fmt"

func getg() uint64

var offset = map[string]int{}

func main() {
	g := getg()
	fmt.Println(g)
}
