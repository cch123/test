# About registers

Go has defined 4 pseudo registers: 

> There are four predeclared symbols that refer to pseudo-registers. These are not real registers, but rather virtual registers maintained by the toolchain, such as a frame pointer. The set of pseudo-registers is the same for all architectures:

> FP: Frame pointer: arguments and locals.

> PC: Program counter: jumps and branches.

> SB: Static base pointer: global symbols.

> SP: Stack pointer: top of stack.

The name of pseudo register SP is the same with the hardware one in plan9 asm code.

How to differ the pseudo register SP from the real hardware one?

The official doc said that the pseudo one has an offset negative, if you use positive offset, then it's hardware SP.

But surely, this is not true. As this code show, whether the register is pseudo doesn't depend on the offset. But the symbol before it.

```go
MOVQ -8(SP), AX  // hardware SP
MOVQ symname-8(SP), BX // pseudo SP
```

```go
MOVQ 8(SP), AX   // hardware SP
MOVQ symname+8(SP), BX // pseudo SP
```

The key is symname in asm code~
