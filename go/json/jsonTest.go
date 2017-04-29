package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type person struct {
	Age        *int `validate:"required"`
	PageNumber int `json:"page_number"`
}

func main() {
	p := person{}
	json.Unmarshal([]byte(`{"page_number":10}`), &p)
	res, _ := json.Marshal(p)
	fmt.Printf("%#v\n", p)
	fmt.Println(string(res))
	validate := validator.New()
	errs:=validate.Struct(p)
	fmt.Println(errs)
	fmt.Println(p.Age)
}
