// https://en.wikipedia.org/wiki/FLAGS_register

package main

import "fmt"

/*
0    CF    Carry flag    Status
1        Reserved, always 1 in EFLAGS [2]
2    PF    Parity flag    Status
3        Reserved
4    AF    Adjust flag    Status
5        Reserved
6    ZF    Zero flag    Status
7    SF    Sign flag    Status
8    TF    Trap flag (single step)    Control
9    IF    Interrupt enable flag    Control
10    DF    Direction flag    Control
11    OF    Overflow flag    Status

[cf:0, zf:1, of:0, sf:0, pf:1, af:0, df:0]]
*/

func main() {
	a := 44
	//if a>>2&1 > 0 {
	if a>>5&1 > 0 {
		fmt.Println("yes, shift two, then pos is >0")
	} else {
		fmt.Println("oh no")
	}
}
