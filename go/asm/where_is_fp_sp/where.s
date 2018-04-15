#include "textflag.h"

TEXT ·where(SB),NOSPLIT, $8-32
	// MOVQ FP, AX
	MOVQ -8(SP), AX // 这是硬件寄存器
	MOVQ x-8(SP), BX // 这是伪寄存器
	MOVQ AX, ret+16(FP)
	MOVQ BX, ret+24(FP)
	RET
