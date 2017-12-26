package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 1000; i++ {
		x := rand.Intn(10)
		fmt.Println(x)
	}
}
