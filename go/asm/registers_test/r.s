#include "textflag.h"

TEXT ·output(SB), NOSPLIT, $0-128
    MOVQ $1, AX
    MOVQ $2, BX
    MOVQ $3, CX
    MOVQ $4, DX
    MOVQ $5, DI
    MOVQ $6, SI
    MOVQ $7, R8
    MOVQ $8, R9
    MOVQ $9, R10
    MOVQ $10, R11
    MOVQ $11, R12
    MOVQ $12, R13
    MOVQ $13, R14
    MOVQ $14, R15
    MOVQ AX, ret1+0(FP)
    MOVQ BX, ret1+8(FP)
    MOVQ CX, ret1+16(FP)
    MOVQ DX, ret1+24(FP)
    MOVQ DI, ret1+32(FP)
    MOVQ SI, ret1+40(FP)
    MOVQ R8, ret1+48(FP)
    MOVQ R9, ret1+56(FP)
    MOVQ R10, ret1+64(FP)
    MOVQ R11, ret1+72(FP)
    MOVQ R12, ret1+80(FP)
    MOVQ R13, ret1+88(FP)
    MOVQ R14, ret1+96(FP)
    MOVQ R15, ret1+104(FP)
    RET
