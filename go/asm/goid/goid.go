package main

import (
	"fmt"
	"time"
)

func getg() uint64

var offset = map[string]int{}

func main() {
	g := getg()
	fmt.Println(g)
	go func() {
		g := getg()
		fmt.Println(g)
	}()
	time.Sleep(time.Second)
}
