## Global Data

Combine DATA directive and GLOBL directive to declare a global variable.

```go
DATA pi+0(SB)/8, $3.1415926
GLOBL pi(SB), RODATA, $8
```

This global data pi, is a package level global data, you can use it in another file.

If you write it this way:

```go
DATA pi<>+0(SB)/8, $3.1415926
GLOBL pi<>(SB), RODATA, $8
```

Then it means file level global data, cannot be referenced in another file. If you refer to it, will produce an error:

```go
relocation target not found
```

## Data Composition

If you want to declare some data like slices or long strings, you can split it to multiple data area, and one var area, as the following code list:

```go
DATA divtab<>+0x00(SB)/4, $1
DATA divtab<>+0x04(SB)/4, $0xe6eaedf0
DATA divtab<>+0x08(SB)/4, $0x81828384
GLOBL divtab<>(SB), RODATA, $12

DATA world<>+0(SB)/8, $"hello wo"
DATA world<>+8(SB)/4, $"rld "
GLOBL world<>+0(SB), RODATA, $12
```
