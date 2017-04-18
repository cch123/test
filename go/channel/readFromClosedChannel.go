package main

import "fmt"

func main() {
	var a = make(chan int, 4)
	close(a) // 如果注释掉这行会panic, 思考一下golang为什么要这么做呢(这个goroutine死锁了~
	data, ok := <-a
	fmt.Println(data, ok)
	data, ok = <-a
	fmt.Println(data, ok)
	data, ok = <-a
	fmt.Println(data, ok)
	data, ok = <-a
	fmt.Println(data, ok)
	data, ok = <-a
	fmt.Println(data, ok)
}
