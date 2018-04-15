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
TEXT ·sum(SB),NOSPLIT, $0-24
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
	MOVQ SI, ret+16(FP)
	RET

// 汇编最后需要有一个空行

