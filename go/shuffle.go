package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x, b, c, d int
	for i := 0; i < 1000; i++ {
		var a = []int{1, 2, 3, 4}
		SliceShuffle(a)
		x += a[0]
		b += a[1]
		c += a[2]
		d += a[3]
	}
	fmt.Println(x, b, c, d)
}

func SliceShuffle(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		seed := rand.Int63n(100000)
		rand.Seed(seed)
		a := rand.Intn(len(slice))

		seed = rand.Int63n(1000000)
		rand.Seed(seed)
		b := rand.Intn(len(slice))

		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
