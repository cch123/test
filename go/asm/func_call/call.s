#include "textflag.h"

TEXT ·callOther(SB), NOSPLIT, $16-8
    MOVQ $4, (SP)
    CALL math·Inf(SB)
    MOVQ 8(SP), AX
    MOVQ AX, ret+0(FP)
    RET
