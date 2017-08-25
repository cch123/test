package main

import (
	"fmt"
	"sync"
)

var codisManager sync.Map

func main() {
	fmt.Printf("%#v\n", codisManager)
	codisManager.Store("abc", "1")
	res, err := codisManager.Load("abc")
	fmt.Println(res, err)
}
