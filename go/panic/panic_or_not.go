// 还是会 panic 的，请注意
package main

import "fmt"

func main() {
	var ch = make(chan int)
	var ch2 = make(chan int, 1)
	ch2 <- 1
	close(ch)
	select {
	case ch <- 1:
	case n := <-ch2:
		fmt.Println("not panic", n)
	}
}
