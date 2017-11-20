// 不管 panic 不 panic，defer 是一定可以执行的
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(10)
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("oh paniced")
			}
			fmt.Printf("%#v\n", wg)
		}()
	*/
	defer fmt.Printf("%#v\n", wg)
	defer wg.Done()
	defer fmt.Println("oh nonono")
	panic(1)
}
