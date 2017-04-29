package main

import (
	"log"
	"net/url"

	"github.com/go-playground/form"
)

// Address contains address information
type Address struct {
	Name  string
	Phone string
}

// User contains user information
type User struct {
	Name        string
	Age         uint8
	Gender      string
	Address     []Address
	Active      bool `form:"active"`
	MapExample  map[string]string
	NestedMap   map[string]map[string]string
	NestedArray [][]string
}

// use a single instance of Decoder, it caches struct info
var decoder *form.Decoder = form.NewDecoder()
var values = url.Values{
	"Name":                []string{"joeybloggs"},
	"Age":                 []string{"3"},
	"Gender":              []string{"Male"},
	"Address[0].Name":     []string{"26 Here Blvd."},
	"Address[0].Phone":    []string{"9(999)999-9999"},
	"Address[1].Name":     []string{"26 There Blvd."},
	"Address[1].Phone":    []string{"1(111)111-1111"},
	"active":              []string{"true"},
	"MapExample[key]":     []string{"value"},
	"NestedMap[key][key]": []string{"value"},
	"NestedArray[0][0]":   []string{"value"},
}

func test() {
	//decoder = form.NewDecoder()

	// this simulates the results of http.Request's ParseForm() function
	//values := parseForm()

	var user User

	// must pass a pointer
	err := decoder.Decode(&user, values)
	if err != nil {
		log.Panic(err)
	}

	//fmt.Printf("%#v\n", user)
}

// this simulates the results of http.Request's ParseForm() function
func parseForm() url.Values {
	return url.Values{
		"Name":                []string{"joeybloggs"},
		"Age":                 []string{"3"},
		"Gender":              []string{"Male"},
		"Address[0].Name":     []string{"26 Here Blvd."},
		"Address[0].Phone":    []string{"9(999)999-9999"},
		"Address[1].Name":     []string{"26 There Blvd."},
		"Address[1].Phone":    []string{"1(111)111-1111"},
		"active":              []string{"true"},
		"MapExample[key]":     []string{"value"},
		"NestedMap[key][key]": []string{"value"},
		"NestedArray[0][0]":   []string{"value"},
	}
}
