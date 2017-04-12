package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	var flag = false
	if x < 0 {
		flag = true
		x = -x
	}
	var str = fmt.Sprintf("%d", x)
	byteStr := []byte(str)
	for i := 0; i < len(byteStr)/2; i++ {
		byteStr[i], byteStr[len(byteStr)-1-i] = byteStr[len(byteStr)-1-i], byteStr[i]
	}

	res, _ := strconv.Atoi(string(byteStr))
	if res > 2147483647 {
		res = 0
	}
	if flag {
		res = -res
	}
	return res
}
