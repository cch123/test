package main

import (
	"encoding/json"
	"fmt"
)

type queryRequest struct {
	// 0，空字符串，nil指针
	Query string `json:"query,omitempty"`
	From  string `json:"from,omitempty"`
	Size  string `json:"size,omitempty"`
	Sort  string `json:"sort,omitempty"`
	X     int    `json:"x,omitempty"`
	Y     []int  `json:"y,omitempty"`
}

func main() {
	var a queryRequest
	a.Query = "fuck"

	// 加不加这句都不会输出的
	a.Y = []int{}

	byteRes, _ := json.Marshal(a)
	fmt.Println(string(byteRes))
	fmt.Printf("%#v\n", a)
}
