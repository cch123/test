package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var js = `{"a" : {"a":"b", "c" :"d"}}`
	var j struct {
		A json.RawMessage `json:"a"`
	}

	var res struct {
		D json.RawMessage
	}

	json.Unmarshal([]byte(js), &j)
	res.D = j.A

	resStr, err := json.Marshal(res)
	fmt.Println(string(resStr), err)

}
