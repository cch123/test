#include "textflag.h"
#include "go_tls.h"

TEXT ·getg(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVQ g(CX), AX
	LEAQ 192(AX), BX // this is the offset
	MOVQ BX, ret+0(FP)
	RET
