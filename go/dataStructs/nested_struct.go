package main

import "fmt"

type gas struct {
	northeast struct {
		lat float32
		lng float32
	}
}

func main() {
	var b = gas{}
	b.northeast.lat = 1.1
	var a = gas{
		northeast: struct {
			lat float32
			lng float32
		}{
			lat: 1.1,
			lng: 1.1,
		},
	}
	fmt.Println(a)
	fmt.Println(b)
}
