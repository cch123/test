package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	res, err := hex.DecodeString("ff")
	fmt.Println(res, err)
	fmt.Println(string(res))

	ch := 'a'
	fmt.Println(int(ch))
	bytes := []byte{123, 123, 11, 23, 0, 23, 12}
	fmt.Println(string(bytes))
}
