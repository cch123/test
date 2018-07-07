package main

import (
	"fmt"
	"math"
)

func main() {
	n1 := math.NaN()
	n2 := math.NaN()
	fmt.Println(n1, n2)
	fmt.Println(n1 == n2)
}
