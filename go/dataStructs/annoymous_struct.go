package main

import "fmt"

func main() {
	var str = struct {
		settings struct {
			name string
		}
	}{
		settings: struct{ name string }{
			name: "fuck",
		},
	}
	fmt.Printf("%+v\n", str)
}
