package main

import (
	"flag"
	"fmt"
)

var a = flag.String("c", "", "flag a ")

func main() {
	// 不parse不出结果
	flag.Parse()
	fmt.Println(*a)
}
