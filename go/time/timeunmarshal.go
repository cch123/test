package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var a = `{"time" : "2015-01-01 00:00:00"}`
	var b struct {
		Time time.Time
	}
	err := json.Unmarshal([]byte(a), &b)
	fmt.Println(b, err)
}
