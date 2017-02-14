package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("oh no")
	fmt.Println(err)
	fmt.Printf("%+v", err)
}
