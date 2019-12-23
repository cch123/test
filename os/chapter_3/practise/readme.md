1. 从源代码执行调试

dlv debug aplusb.go

执行 disass
```
    aplusb.go:3    0x104f2c0    65488b0c25a0080000    mov rcx, qword ptr gs:[0x8a0]
    aplusb.go:3    0x104f2c9    483b6110        cmp rsp, qword ptr [rcx+0x10]
    aplusb.go:3    0x104f2cd    7655            jbe 0x104f324
    aplusb.go:3    0x104f2cf*    4883ec28        sub rsp, 0x28
    aplusb.go:3    0x104f2d3    48896c2420        mov qword ptr [rsp+0x20], rbp
    aplusb.go:3    0x104f2d8    488d6c2420        lea rbp, ptr [rsp+0x20]
    aplusb.go:4    0x104f2dd    48c744241801000000    mov qword ptr [rsp+0x18], 0x1
    aplusb.go:5    0x104f2e6    48c744241002000000    mov qword ptr [rsp+0x10], 0x2
    aplusb.go:6    0x104f2ef    488b442418        mov rax, qword ptr [rsp+0x18]
    aplusb.go:6    0x104f2f4    4883c002        add rax, 0x2
    aplusb.go:6    0x104f2f8    4889442408        mov qword ptr [rsp+0x8], rax
    aplusb.go:7    0x104f2fd    e85e3ffdff        call 0x1023260 runtime.printlock
    aplusb.go:7    0x104f302    488b442408        mov rax, qword ptr [rsp+0x8]
    aplusb.go:7    0x104f307    48890424        mov qword ptr [rsp], rax
    aplusb.go:7    0x104f30b    e83047fdff        call 0x1023a40 runtime.printint
    aplusb.go:7    0x104f310    e8fb41fdff        call 0x1023510 runtime.printnl
    aplusb.go:7    0x104f315    e8d63ffdff        call 0x10232f0 runtime.printunlock
```



2. 从可执行文件调试

dlv exec aplusb

```
    aplusb.go:3    0x104f2c0    65488b0c25a0080000    mov rcx, qword ptr gs:[0x8a0]
    aplusb.go:3    0x104f2c9    483b6110        cmp rsp, qword ptr [rcx+0x10]
    aplusb.go:3    0x104f2cd    7634            jbe 0x104f303
    aplusb.go:3    0x104f2cf*    4883ec10        sub rsp, 0x10
    aplusb.go:3    0x104f2d3    48896c2408        mov qword ptr [rsp+0x8], rbp
    aplusb.go:3    0x104f2d8    488d6c2408        lea rbp, ptr [rsp+0x8]
    aplusb.go:7    0x104f2dd    e87e3ffdff        call 0x1023260 runtime.printlock
    aplusb.go:7    0x104f2e2    48c7042403000000    mov qword ptr [rsp], 0x3 => 直接计算得到了这个 0x3，完全省略了计算过程
    aplusb.go:7    0x104f2ea    e85147fdff        call 0x1023a40 runtime.printint
    aplusb.go:7    0x104f2ef    e81c42fdff        call 0x1023510 runtime.printnl
    aplusb.go:7    0x104f2f4    e8f73ffdff        call 0x10232f0 runtime.printunlock
```

注意，编译后的结果已经被编译器优化了，特别是变量赋值和相加的部分
