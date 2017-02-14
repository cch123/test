package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("whoops")
	fmt.Println(err)
	fmt.Printf("%+v", err)
}
