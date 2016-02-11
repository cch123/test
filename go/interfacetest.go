package main

import "fmt"

//面积和周长
type calc interface {
	girth() float32
	area() float32
}
type square struct {
	length float32
	width  float32
}

func (s square) girth() float32 {
	return s.length*2 + s.width*2
}
func (s square) area() float32 {
	return s.length * s.width
}

func main() {
	var s square = square{width: 1.1, length: 2.2}
	fmt.Printf("%+v\n", s)
	fmt.Printf("%.2f\n", s.area())
	fmt.Printf("%f\n", s.girth())
}
