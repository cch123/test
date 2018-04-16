#include "textflag.h"

DATA pi+0(SB)/8, $3.1415926
GLOBL pi(SB), RODATA, $8
DATA two+0(SB)/8, $22222
GLOBL two(SB), RODATA, $24
DATA one+0(SB)/8, $12345
GLOBL one(SB), RODATA, $8
DATA three+0(SB)/8, $99
GLOBL three(SB), RODATA, $8

// 注意这里最后的值，如果 rodata 里已经有其它变量了
// 就不是所谓的变量 size 了，官方这里严重误导

// func output() (int, float64)
TEXT ·output(SB),NOSPLIT,$0-16
    MOVQ one+0(SB), AX
    // 注意，是没有办法直接从内存将数据搬到另一块内存的
    // 如果要做这种内存搬运，必须把数据搬到 cpu，再从 cpu 写入到内存
    MOVQ AX, ret+0(FP)
    // 这个 symbol 的名字 pi 需要和 GLOBL 的 symbol 一致
    // 否则会提示 relocation xxx not found
    MOVQ pi+0(SB), AX
    MOVQ AX, ret+8(FP)
    MOVQ three+0(SB), AX
    MOVQ AX, ret+16(FP)
    RET
