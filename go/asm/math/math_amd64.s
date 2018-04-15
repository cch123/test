#include "textflag.h"

// plan9 和 intel IA64 的寄存器使用顺序是反着的
TEXT ·add(SB),NOSPLIT, $0-24
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), BX
	ADDQ AX, BX
	MOVQ BX, ret+16(FP)
	RET

TEXT ·minus(SB),NOSPLIT, $0-24
	MOVQ x+0(FP), DX
	MOVQ y+8(FP), CX
	SUBQ CX, DX

	// if using FP register, must give it a symbol
	// or will cause : cannot reference FP without a symbol
	MOVQ DX, ret+16(FP)
	RET

// func sum(sl []int64) int64
TEXT ·sum(SB),NOSPLIT, $0-32
	MOVQ $0, SI
	MOVQ sl+0(FP), BX // &sl[0], addr of the first elem
	MOVQ sl+8(FP), CX // len(sl)
	INCQ CX // CX++，为了让 len = 1 的时候也能进循环

start:
	DECQ CX       // CX--
	JZ   done
	ADDQ (BX), SI // SI += *BX
	ADDQ $8, BX   // 指针移动
	JMP  start

done:
    // 返回地址是 24 是怎么得来的呢？
    // 可以通过 go tool compile -S math.go 得知
    // 在调用 sum1 函数时，会传入三个值，分别为:
    // slice 的首地址、slice 的 len， slice 的 cap
    // 不过我们这里的求和只需要 len，但 cap 依然会占用参数的空间
    // 就是 16(FP)
	MOVQ SI, ret+24(FP)
	RET

// 汇编最后需要有一个空行

