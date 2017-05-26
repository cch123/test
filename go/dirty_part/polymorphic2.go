package main

import (
	"encoding/json"
	"fmt"
)

type request struct {
	Operations map[string]op `json:"operations"`
}
type op struct {
	operation
	Test string `json:"test"`
}
type operation struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (o *operation) UnmarshalJSON(b []byte) error {
	type xoperation operation
	xo := &xoperation{Width: 500, Height: 500}
	if err := json.Unmarshal(b, xo); err != nil {
		return err
	}
	*o = operation(*xo)
	return nil
}

func main() {
	jsonStr := `{
            "operations": {
                "001": {
                     "test":"test",
                    "width": 100
                }
            }
        }`
	req := request{}
	json.Unmarshal([]byte(jsonStr), &req)
	fmt.Printf("%#v\n", req)
}
