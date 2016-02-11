package main

import "fmt"

type Person struct {
	name    string
	age     int
	history [4]string
}

func main() {
	var i = 5
	var x = "ststs"
	var y = 'x'
	var z = [5]int{1, 2, 3, 4, 5}
	fmt.Println(z)
	per := Person{name: "cch", age: 12, history: [4]string{"1", "2", "3", "4"}}
	fmt.Println(per)
	fmt.Printf("%+v\n", per)
	fmt.Printf("%v\n", per)
	fmt.Printf("%#v\n", per)
	println(i, x, y)
}
