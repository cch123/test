package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{43, 32, 12, 32, 44}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
