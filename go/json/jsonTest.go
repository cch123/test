package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Age        int
	PageNumber int `json:"page_number"`
}

func main() {
	p := person{}
	json.Unmarshal([]byte(`{"page_number":10}`), &p)
	fmt.Printf("%#v\n", p)
}
