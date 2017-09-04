package main

import (
	"fmt"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

func main() {
	mu, _ := uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
	// mov eax, 1234
	code := []byte{184, 210, 4, 0, 0}
	mu.MemMap(0x1000, 0x1000)
	mu.MemWrite(0x1000, code)
	if err := mu.Start(0x1000, 0x1000+uint64(len(code))); err != nil {
		panic(err)
	}
	err := mu.RegWrite(uc.X86_REG_R10, 10101010101010101010)
	fmt.Println(err)
	eax, _ := mu.RegRead(uc.X86_REG_EAX)
	r10, err := mu.RegRead(uc.X86_REG_R10)
	fmt.Printf("EAX is now: %d\n", eax)
	fmt.Printf("r10 is now: %d\n", r10)
	fmt.Println(err)
}
