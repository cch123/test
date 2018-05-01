package main

import (
	"fmt"
)

func output(int) (int, int, int)
func output2(int) (int, int, int)

func main() {
	a, b, c := output(987654321)
	fmt.Println(a, b, c)
	a, b, c = output2(98765)
	fmt.Println(a, b, c)
}
