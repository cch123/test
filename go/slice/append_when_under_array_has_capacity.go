package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	// will use the cap of array underneath
	b := a[1:3]
	fmt.Println(len(b), cap(b))
}
