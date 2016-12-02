package main

import (
	"fmt"
	"regexp"
)

func main() {
	isok, err := regexp.MatchString("^[0-9a-zA-Z_]+$", "g_service_worksheet ")
	fmt.Println(isok, err)
	isok, err = regexp.MatchString("^[0-9a-zA-Z_]+$", "中文英文ohyes")
	fmt.Println(isok, err)
	isok, err = regexp.MatchString("^[0-9a-zA-Z_]+$", "english_space ")
	fmt.Println(isok, err)
	isok, err = regexp.MatchString("^[0-9a-zA-Z_]+$", "normal_field")
	fmt.Println(isok, err)
}
