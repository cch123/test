#include "textflag.h"

DATA pi+0(SB)/8, $3.1415926
DATA one+8(SB)/8, $12345
GLOBL pi(SB), RODATA, $8
GLOBL one(SB), RODATA, $16

// 注意这里最后的值，如果 rodata 里已经有其它变量了
// 不能直接就用 8 的，网上的资料有误导性

// func output() (int, float64)
TEXT ·output(SB),NOSPLIT,$0-16
    MOVQ one+8(SB), AX
    MOVQ AX, ret+0(FP)
    MOVQ pi+0(SB), AX
    MOVQ AX, ret+8(FP)
    RET
