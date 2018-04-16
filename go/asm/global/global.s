#include "textflag.h"

DATA pi+0(SB)/8, $3.1415926
DATA one+8(SB)/8, $12345
DATA two+16(SB)/8, $22222
GLOBL pi(SB), RODATA, $8
GLOBL one(SB), RODATA, $16
GLOBL two(SB), RODATA, $24

// 注意这里最后的值，如果 rodata 里已经有其它变量了
// 就不是所谓的变量 size 了，官方这里严重误导

// func output() (int, float64)
TEXT ·output(SB),NOSPLIT,$0-16
    MOVQ one+8(SB), AX
    MOVQ AX, ret+0(FP)
    MOVQ pi+0(SB), AX
    MOVQ AX, ret+8(FP)
    MOVQ two+16(SB), AX
    MOVQ AX, ret+16(FP)
    RET
