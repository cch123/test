#include "textflag.h"

TEXT ·where(SB), NOSPLIT, $8-32
	// MOVQ FP, AX
	MOVQ -8(SP), AX     // this is hardwhere SP
	MOVQ x-8(SP), BX    // this is the virtual register SP
	MOVQ 8(SP), DX
	MOVQ x+8(SP), CX
	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)
	MOVQ DX, ret+16(FP)
	MOVQ CX, ret+24(FP)
	RET
