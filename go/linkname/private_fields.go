package main

import (
	"fmt"
	"unsafe"
	_ "unsafe"
)

//go:linkname allgs runtime.allgs
var allgs []*g

//go:linkname gentraceback runtime.gentraceback
func gentraceback(pc0, sp0, lr0 uintptr, gp *g, skip int, pcbuf *uintptr, max int, callback func(uintptr, unsafe.Pointer) bool, v unsafe.Pointer, flags uint) int

type StackRecord struct {
	Stack0 [32]uintptr // stack trace for this record; ends at first 0 entry
}

type g struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       [2]uintptr // offset known to runtime/cgo
	stackguard0 uintptr    // offset known to liblink
	stackguard1 uintptr    // offset known to liblink

	_panic       uintptr // innermost panic - offset known to liblink
	_defer       uintptr // innermost defer
	m            uintptr // current m; offset known to arm liblink
	sched        [7]uintptr
	syscallsp    uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc    uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp     uintptr        // expected sp at top of stack, to check in traceback
	param        unsafe.Pointer // passed parameter on wakeup
	atomicstatus uint32
	stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid         int64
}

func main() {
	fmt.Println(allgs)
	fmt.Println(allgs[0].stack[0], allgs[0].stack[1])

	for _, gp := range allgs {
		var x StackRecord
		if gp.goid != 1 {
			gentraceback(gp.sched[1], gp.sched[0], 0, gp, 0, &x.Stack0[0], len(x.Stack0), nil, nil, 0)
			fmt.Println(x)
		}
	}
}
