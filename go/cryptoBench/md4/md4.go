package main

import "golang.org/x/crypto/md4"

var hash []byte

func test() {

	s1 := "Hello, MD4 test text"
	ctx := md4.New()
	ctx.Write([]byte(s1))
	hash = ctx.Sum(nil)

}
