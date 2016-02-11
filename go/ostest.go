package main

import "os/exec"
import "reflect"

import "fmt"
import "bytes"

func main() {
	dateCmd := exec.Command("date")
	dateOut, _ := dateCmd.Output()
	println(dateOut)
	println(string(dateOut))
	buf1 := bytes.NewBufferString("fuck")
	fmt.Println(reflect.TypeOf(buf1))
}
