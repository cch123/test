#include "textflag.h"

TEXT ·main(SB), $24-0
	MOVQ $0, a-8*2(SP) // a = 0
	MOVQ $0, b-8*1(SP) // b = 0

	// 将新的值写入a对应内存
	MOVQ $10, AX       // AX = 10
	MOVQ AX, a-8*2(SP) // a = AX

	// 以a为参数调用函数
	MOVQ AX, 0(SP)
	CALL	runtime·printint(SB)
	CALL	runtime·printnl(SB)

	// 函数调用后, AX/BX 寄存器可能被污染, 需要重新加载
	MOVQ a-8*2(SP), AX // AX = a
	MOVQ b-8*1(SP), BX // BX = b

	// 计算b值, 并写入内存
	MOVQ AX, BX        // BX = AX  // b = a
	ADDQ BX, BX        // BX += BX // b += a
	IMULQ AX, BX        // BX *= AX // b *= a
	MOVQ BX, b-8*1(SP) // b = BX

	// 以b为参数调用函数
	MOVQ BX, 0(SP)
	CALL runtime·printint(SB)
	CALL runtime·printnl(SB)

	RET
