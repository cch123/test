package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var err = errors.New("fuck")
	err = errors.Wrap(err, "this **** happened here")
	fmt.Printf("%+v\n", err)
	err = errors.WithStack(err)
	fmt.Printf("%+v\n", err)

	x := fmt.Sprintf("%+v", err)
	fmt.Println(x)
}
