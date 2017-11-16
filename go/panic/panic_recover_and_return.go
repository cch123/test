package main

import "fmt"

func test(output *int) (output2 int) {
	defer func() {
		if err := recover(); err != nil {
			*output = 444
			output2 = 23323
		}
	}()
	*output = 2
	panic(1)
}

func main() {
	x := 0
	y := test(&x)
	fmt.Println(x, y)
}
