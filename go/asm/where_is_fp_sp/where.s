#include "textflag.h"

TEXT ·where(SB), NOSPLIT, $8-32
	// MOVQ FP, AX
	MOVQ -8(SP), AX     // this is hardwhere SP
	MOVQ x-8(SP), BX    // this is the virtual register SP
	MOVQ AX, ret+16(FP)
	MOVQ BX, ret+24(FP)
	RET

