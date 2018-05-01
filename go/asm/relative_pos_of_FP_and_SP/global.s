#include "textflag.h"

// func output(int) (int, int)
// frame size is 0 or 8 will affect the result
TEXT ·output(SB), $8-48
	MOVQ 24(SP), DX
	MOVQ DX, ret3+24(FP)
	MOVQ argx+16(SP), BX
	MOVQ BX, ret2+16(FP)
	MOVQ arg0+0(FP), AX
	MOVQ AX, ret1+8(FP)
	RET

TEXT ·output2(SB), $0-48
	MOVQ 8(SP), DX
	MOVQ DX, ret3+24(FP)
	MOVQ argx+8(SP), BX
	MOVQ BX, ret2+16(FP)
	MOVQ arg0+0(FP), AX
	MOVQ AX, ret1+8(FP)
	RET

