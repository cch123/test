package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "abc@def@gh"
	res := strings.SplitN(a, "@", 2)
	fmt.Println(res)
}
