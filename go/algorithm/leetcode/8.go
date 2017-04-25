package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("  -0012a42"))
}

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	validBytes := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-'}
	newStr := []byte{}
	for _, b := range []byte(str) {
		exist := false
		for _, v := range validBytes {
			if b == v {
				exist = true
			}
		}

		if len(newStr) > 0 && exist == false {
			break
		} else {
			newStr = append(newStr, b)
		}
	}

	i, _ := strconv.Atoi(string(newStr))
	if i > 2147483648 {
		i = 2147483647
	}

	if i <= -2147483649 {
		i = -2147483648
	}
	return i
}
