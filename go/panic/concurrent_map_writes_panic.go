package main

import "fmt"

var a = map[string]int{}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered")
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered")
			}
		}()
		for {
			a["a"] = 1
		}
	}()

	for {
		a["a"] = 1
	}

}
