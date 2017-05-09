package main

import "fmt"

func main() {

	var a = map[string]int{
		"1": 1,
		"2": 2,
	}

	var counter int
	for counter < 100 {
		for key, _ := range a {
			fmt.Printf("%s ", key)
		}
		println()
	}
}
