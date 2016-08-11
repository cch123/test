package main

import "encoding/json"
import "fmt"

func main() {
	a := `{"a" : {"b" : 1}}`
	var b interface{}
	json.Unmarshal([]byte(a), &b)
	fmt.Println(b)
	//c, ok := b.(map[string](map[string]int))
	c := map[string](map[string]int)(b)
	fmt.Println(c)
}
