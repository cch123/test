package main

import (
	"fmt"
	"regexp"
)

func main() {
	isok, err := regexp.MatchString("^[0-9a-zA-Z_]+$", "g_service_worksheet ")
	fmt.Println(isok, err)
}
