package main

import "fmt"

func where(a int, b int) (int, int)

func main() {
	a, b := where(1, 2)
	fmt.Println(a, b)
}
