#include "textflag.h"
#include "go_tls.h"

TEXT ·getg(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVQ g(CX), AX
	LEAQ 152(AX), BX   // this is the offset
	MOVQ (BX), DX
	MOVQ DX, ret+0(FP)
	RET
