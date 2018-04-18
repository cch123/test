#include "textflag.h"

// func output(int) (int, int)
// 不管这里的 frame size 是 0 还是 8
// 这个程序的输出都不会变
// 说明伪寄存器 SP 和 伪寄存器 FP 就是差了 16 个字节？？
TEXT ·output(SB), $8-24
	MOVQ arg0+0(FP), AX
	MOVQ AX, ret1+8(FP)
	MOVQ argx+16(SP), BX
	MOVQ BX, ret2+16(FP)
	MOVQ 24(SP), DX
	MOVQ DX, ret3+24(FP)
	RET
