package main

import "os"

var user = os.Getenv("USER")

func init() {
	println("init is automatically called")
	if user == "" {
		panic("no value for $USER")
	}
	println("the value of user is " + user)
}

func main() {
}
