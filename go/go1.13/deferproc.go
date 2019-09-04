package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer func() {
			for {
				var a = make([]int, 128)
				fmt.Println(a)
			}
		}()
	}
}
