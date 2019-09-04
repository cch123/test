# Go 1.13 study

## defer 在栈上分配

```go
	0x003a 00058 (deferstack.go:4)	PCDATA	$0, $1
	0x003a 00058 (deferstack.go:4)	LEAQ	""..autotmp_1+8(SP), AX
	0x003f 00063 (deferstack.go:4)	PCDATA	$0, $0
	0x003f 00063 (deferstack.go:4)	MOVQ	AX, (SP)
	0x0043 00067 (deferstack.go:4)	CALL	runtime.deferprocStack(SB)
	0x0048 00072 (deferstack.go:4)	TESTL	AX, AX
	0x004a 00074 (deferstack.go:4)	JNE	92
	0x004c 00076 (deferstack.go:5)	XCHGL	AX, AX
	0x004d 00077 (deferstack.go:5)	CALL	runtime.deferreturn(SB)
	0x0052 00082 (deferstack.go:5)	MOVQ	64(SP), BP
	0x0057 00087 (deferstack.go:5)	ADDQ	$72, SP
	0x005b 00091 (deferstack.go:5)	RET
```

deferproc 变成了 deferprocStack

什么条件下使用 deferproc，什么条件下使用 deferprocStack？

TODO

