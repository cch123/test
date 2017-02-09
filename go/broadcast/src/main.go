package main

import (
	"broadcast"
	"fmt"
)

var b = broadcast.NewBroadcaster()

func main() {
	//b.Write("abc")
	r := b.Listen()
	b.Write("a")
	b.Write("b")
	b.Write("c")
	b.Write("d")
	fmt.Println(r.Read())
	fmt.Println(r.Read())
	fmt.Println(r.Read())
	fmt.Println(r.Read())
	fmt.Println(r)
}
