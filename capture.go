package main

import "fmt"

func main() {
	var a = []*int{}
	for i := 0; i < 10; i++ {
		func() {
			defer println(i)
		}()
	}
	fmt.Println(a)
}
