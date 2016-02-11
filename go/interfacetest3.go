package main

import "math"

import "fmt"

func abs(i interface{}) interface{} {
	var res interface{}
	switch v := i.(type) {
	case int:
		res = v
		if v < 0 {
			res = -v
		}
	case float32:
		res = v
		if v < 0 {
			res = -v
		}
	case float64:
		res = v
		if v < 0 {
			res = -v
		}
	}
	return res
}

func main() {
	var i = -1
	fmt.Println(abs(i))
	var j = -3.1
	fmt.Println(abs(j))
	//println(abs(i))
	//fmt.Println(abs(i))
	fmt.Println(math.Sqrt2)
}
