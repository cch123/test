package main

import (
	"fmt"
	"os"
)

func exit_this(f os.File) {
	f.Close()
	println("file closed")
}

func main() {
	f, _ := os.Open("./arrtest.go")
	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", *f)
	defer exit_this(*f)
}
