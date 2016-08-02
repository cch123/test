package main

//Go 在多个 crypto/* 包中实现了一系列散列函数。
import "crypto/sha1"

var bs []byte

func test() {
	s := "Hello, MD4 test text"
	h := sha1.New()
	h.Write([]byte(s))
	bs = h.Sum(nil)
}
