package main

import (
	"fmt"
)

func output() (int, float64, int)
func output2() (int, float64, int)

func main() {
	a, b, c := output()
	fmt.Println(a, b, c)
	a, b, c =output2()
	fmt.Println(a, b, c)
}
