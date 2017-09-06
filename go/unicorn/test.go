package main

import (
	"encoding/hex"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

func main() {
	mu, _ := uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
	// mov rax, 0x1234
	// code, _ := hex.DecodeString("48c7c034120000")

	// push rax
	//code, _ := hex.DecodeString("50")
	code, _ := hex.DecodeString("58")
	// push rax
	mu.MemMap(0x1000, 0x1000)
	mu.MemWrite(0x1000, code)
	mu.RegWrite(uc.X86_REG_RBP, 0x2000)
	mu.RegWrite(uc.X86_REG_RSP, 0x1200)
	if err := mu.Start(0x1000, 0x1000+uint64(len(code))); err != nil {
		panic(err)
	}
}
