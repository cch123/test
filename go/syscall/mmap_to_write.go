package main

import (
	"fmt"
	"io/ioutil"
	"os"

	mmap "github.com/edsrzf/mmap-go"
)

func openFile(flags int) *os.File {
	f, err := os.OpenFile("./file.dat", flags, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func main() {
	f, err := os.OpenFile("./file.dat", os.O_RDWR, 0644)
	defer f.Close()
	mmap, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		fmt.Printf("error mapping: %s", err)
	}
	fmt.Println(len(mmap))
	defer mmap.Unmap()

	mmap[0] = '3'
	mmap.Flush()

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error reading file: %s", err)
	}
	fmt.Println(string(fileData))

	// leave things how we found them
	mmap[9] = '9'
	mmap.Flush()

}
