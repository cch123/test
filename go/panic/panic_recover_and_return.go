package main

import "fmt"

func test(output *int) {
	defer func() {
		if err := recover(); err != nil {
			*output = 444
		}
	}()
	*output = 2
	panic(1)
}

func main() {
	x := 0
	test(&x)
	fmt.Println(x)
}
