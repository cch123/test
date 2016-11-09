package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var a = `g_service_1323`
	a = strings.TrimRightFunc(a, unicode.IsDigit)
	a = strings.TrimRight(a, "_")
	fmt.Println(a)
}
