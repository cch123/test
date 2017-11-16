package main

import "fmt"

func main() {
	var a = [10]int{1, 3, 4}
	var b = a[2:4:5]
	fmt.Println(b)
}
