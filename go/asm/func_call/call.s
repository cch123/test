#include "textflag.h"

// func callOther() string
TEXT ·callOther(SB), NOSPLIT, $24-16
    MOVQ $323232323, (SP)
    CALL strconv·Itoa(SB)
    MOVQ 8(SP), AX
    MOVQ 16(SP), BX
    MOVQ AX, ret+0(FP)
    MOVQ BX, ret+8(FP)
    RET
