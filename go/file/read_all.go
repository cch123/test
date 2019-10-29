package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	contentBytes, err := ioutil.ReadFile("./list.txt")
	if err != nil {
		println(err)
		return
	}
	lines := strings.Split(string(contentBytes), "\n")
	for _, l := range lines {
	}
}
