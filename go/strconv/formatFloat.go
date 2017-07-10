package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 0.00000001
	f := byte('e')
	fmt.Println(strconv.FormatFloat(float64(val), f, -1, 64))
}
