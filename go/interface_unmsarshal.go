package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var a struct {
	A interface{} `json:"a"`
}

func main() {
	jsonStr := `{"a" : false}`
	json.Unmarshal([]byte(jsonStr), &a)
	fmt.Println(a)
	fmt.Println(a.A == nil)
	fmt.Println(reflect.TypeOf(a.A))
	fmt.Println(reflect.ValueOf(a.A))
}
