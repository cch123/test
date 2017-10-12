package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	size := 1000
	file, err := os.Open("./file.dat")
	if err != nil {
		fmt.Println(err)
		return
	}

	mmap, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err == nil {
		println(string(mmap))
	}

}
