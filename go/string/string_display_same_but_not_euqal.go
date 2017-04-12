package main

import "fmt"

func main() {
	var a = []byte{49, 0, 50}
	var b = []byte{49, 50, 0}
	fmt.Println(string(a), string(b))
	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)
	fmt.Println(len(a), len(b))
	fmt.Println(string(a) == string(b))
}
