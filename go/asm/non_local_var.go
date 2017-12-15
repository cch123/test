package main

/*
"".xxx STEXT nosplit size=31 args=0x30 locals=0x0
    0x0000 00000 (non_local_var.go:3)    TEXT    "".xxx(SB), NOSPLIT, $0-48
    0x0000 00000 (non_local_var.go:3)    FUNCDATA    $0, gclocals·26c19b003b4032a46d3e8db29831f3fe(SB)
    0x0000 00000 (non_local_var.go:3)    FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (non_local_var.go:3)    MOVQ    "".a+8(SP), AX //把 a 赋值到 AX 寄存器
    0x0005 00005 (non_local_var.go:5)    MOVQ    AX, "".e+32(SP)  // 把 a 赋值给 e
    0x000a 00010 (non_local_var.go:3)    MOVQ    "".b+16(SP), AX // 把 b 赋值到 AX 寄存器
    0x000f 00015 (non_local_var.go:5)    MOVQ    AX, "".f+40(SP) // 以此类推
    0x0014 00020 (non_local_var.go:3)    MOVQ    "".c+24(SP), AX
    0x0019 00025 (non_local_var.go:5)    MOVQ    AX, "".g+48(SP)
    0x001e 00030 (non_local_var.go:5)    RET
    0x0000 48 8b 44 24 08 48 89 44 24 20 48 8b 44 24 10 48  H.D$.H.D$ H.D$.H
    0x0010 89 44 24 28 48 8b 44 24 18 48 89 44 24 30 c3     .D$(H.D$.H.D$0.
*/
func xxx(a, b, c int) (e, f, g int) {
	e, f, g = a, b, c
	return
}
func main() {

}
