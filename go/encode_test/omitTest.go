package main

import (
	"encoding/json"
	"fmt"
)

type queryRequest struct {
	Query string `json:"query,omitempty"`
	From  string `json:"from,omitempty"`
	Size  string `json:"size,omitempty"`
	Sort  string `json:"sort,omitempty"`
}

func main() {
	var a queryRequest
	a.Query = "fuck"
	byteRes, _ := json.Marshal(a)
	fmt.Println(string(byteRes))
}
