#include "textflag.h"

DATA px+0(SB)/8, $2.1
DATA ox+8(SB)/8, $5555
DATA tx+16(SB)/8, $333
GLOBL px(SB), RODATA, $8
GLOBL ox(SB), RODATA, $16
GLOBL tx(SB), RODATA, $24
// GLOBL 级别的变量，名字和其它包里的不能冲突

// 注意这里最后的值，如果 rodata 里已经有其它变量了
// 就不是所谓的变量 size 了，官方这里严重误导

// func output2() (int, float64)
TEXT ·output2(SB),NOSPLIT,$0-16
    MOVQ ox+8(SB), AX
    // 注意，是没有办法直接从内存将数据搬到另一块内存的
    // 如果要做这种内存搬运，必须把数据搬到 cpu，再从 cpu 写入到内存
    MOVQ AX, ret+0(FP)
    MOVQ px+0(SB), AX
    MOVQ AX, ret+8(FP)
    MOVQ two+16(SB), AX
    MOVQ AX, ret+16(FP)
    RET
