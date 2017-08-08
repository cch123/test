package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Dislike string        `msgpack:"dislike"`
	Age     int           `msgpack:"age"`
	CallDad func() string `msgpack:"-"`
}

func main() {
	var a = Person{Dislike: "abc", Age: 1}
	res, err := msgpack.Marshal(a)
	fmt.Println(string(res))
	fmt.Println(err)

	var b Person
	err = msgpack.Unmarshal(res, &b)
	fmt.Printf("%#v\n", b)
}
