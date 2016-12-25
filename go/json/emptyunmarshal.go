package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B bool `json:"b,omitempty"`
}

func main() {
	d1 := `{"b":false}`
	d3 := `{"b":t}`
	d2 := `{}`
	var a A
	var b A
	var c A
	var d = make(map[string]interface{})
	json.Unmarshal([]byte(d1), &a)
	json.Unmarshal([]byte(d2), &b)
	json.Unmarshal([]byte(d3), &c)
	json.Unmarshal([]byte(d1), &d)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
