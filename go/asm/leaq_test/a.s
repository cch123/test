#include "textflag.h"

TEXT ·run(SB), NOSPLIT, $0-8
    XORQ BX, BX
    XORQ AX, AX

    // scale 只能是 0、2、4、8
    // LEAQ (BX)(AX*3), CX
    // ./a.s:6: bad scale: 3

    LEAQ (BX)(AX*8), CX

    // 用 LEAQ 的话，即使是两个寄存器值直接相加，也必须提供 scale
    // 下面这样是不行的
    // LEAQ (BX)(AX), CX
    // asm: asmidx: bad address 0/2064/2067

    // 三个寄存器做运算，还是别想了
    // LEAQ DX(BX)(AX*8), CX
    // ./a.s:13: expected end of operand, found (

    MOVQ CX, ret+0(FP)
    RET
