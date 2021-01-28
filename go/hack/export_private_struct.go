package main

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"unsafe"
)

const BlockSize = 64

// digest represents the partial evaluation of a checksum.
type digest struct {
	s   [4]uint32
	x   [BlockSize]byte
	nx  int
	len uint64
}

func main() {
	var md5s = md5.New()
	fmt.Printf("%#v\n", md5s)
	var ptr = reflect.ValueOf(md5s).Pointer()
	var x = (*digest)(unsafe.Pointer(ptr))
	fmt.Printf("%#v\n", x)
}

