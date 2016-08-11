package main

import "encoding/json"
import "fmt"

func main() {
	a := `{"a" : {"b" : 1}}`
	var b map[string](map[string]interface{})
	json.Unmarshal([]byte(a), &b)
	// var bb interface{}
	//json.Unmarshal([]byte(a), &bb)
	fmt.Println(b)
	fmt.Println(b["a"]["b"])
	// var c = bb.(map[string](map[string]interface{}))
	// 这样做断言是不行的
}
