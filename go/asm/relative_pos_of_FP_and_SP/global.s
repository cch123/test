#include "textflag.h"

// func output(int) (int, int)
// frame size 是 0 还是 8
// 会影响程序的输出结果
TEXT ·output(SB), $8-48
	MOVQ 32(SP), DX
	MOVQ DX, ret3+24(FP)
	MOVQ argx+16(SP), BX
	MOVQ BX, ret2+16(FP)
	MOVQ arg0+0(FP), AX
	MOVQ AX, ret1+8(FP)
	RET
