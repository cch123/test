package main

import "fmt"

// Both types still need to have:
// same sequence of fields (the order matters)
// corresponding fields with same type.
type Person struct {
	age    int
	height int
	name   string
}

// the sequence of these two struct must be the same
type PersonVo struct {
	age    int    `json:"old"`
	height int    `json:"tall"`
	name   string `json:"person_name"`
}

func main() {
	v := PersonVo{1, 2, "aa"}
	// 1.7
	d := Person{v.age, v.height, v.name}
	// 1.8
	d2 := Person(v)
	fmt.Println("%#v", d)
	fmt.Println("%#v", d2)
}
