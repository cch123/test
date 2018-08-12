package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("0123456789ABCDEF")
	ciphertext, _ := hex.DecodeString("08f24c28f087fa30c8e31355c5cda423d26130a3905979fcbb1350956bf02866cb9ac1ea65905bd0baf6ae3309613792e99579b18en")

	nonce, _ := hex.DecodeString("000000000000000000000000")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(plaintext), len(plaintext))
}
