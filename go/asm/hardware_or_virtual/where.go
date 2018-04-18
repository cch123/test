package main

import "fmt"

func where() (int, int, int, int)

func main() {
	a, b, c, d := where()
	fmt.Println(a, b, c, d)
}
