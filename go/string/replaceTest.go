package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = `('a','b','c')`
	fmt.Println(strings.Replace(a, `'`, `"`, -1))
}
