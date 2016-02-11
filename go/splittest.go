package main

import "fmt"
import "strings"
import "errors"

func get_uid_and_content(str string) ([]string, error) {
	arr := strings.SplitN(str, "\r\n", 3)
	if len(arr) < 3 {
		return arr, errors.New("输入参数有误")
	}
	return arr, nil
}

func main() {
	str := "uid1\r\nuid2\r\n"
	//fmt.Printf("%q\n", get_uid_and_content("uid1\r\nuid2\r\ncontent"))
	arr, _ := get_uid_and_content("uid1\r\nuid2\r\ncontent")
	fmt.Printf("%q\n", arr)

	println(str)
}
