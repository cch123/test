#include "textflag.h"


// func output(int) (int, int)
TEXT ·output(SB),$0-24
    MOVQ arg0+0(FP), AX
    MOVQ AX, ret1+8(FP)
    MOVQ argx+16(SP), BX
    MOVQ BX, ret2+16(FP)
    RET
