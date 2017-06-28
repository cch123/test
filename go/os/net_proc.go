package main

import (
	"fmt"
	"syscall"
)

func main() {
	res, err := syscall.Sysctl("security.mac.device_enforce")
	fmt.Printf("%#v\n", res)
	fmt.Println(err)
}
