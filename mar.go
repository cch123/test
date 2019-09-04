package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B
}
type B struct {
	Age int
}

func main() {
	var x = A{}
	y, _ := json.Marshal(x)
	fmt.Println(string(y))
}
