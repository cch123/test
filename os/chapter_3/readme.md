`global _start` 在汇编里是一种约定，表示有一个全局的 symbol，这个 symbol 标记程序的入口位置

如果修改了 `_start` 的名字，或者`_start` 变成局部符号，那么都会导致链接失败，因为链接器找不到程序入口了。

nasm -f elf64 -o hello.o hello.s

readelf -s hello.o
```
  -s --syms              Display the symbol table
```

输出结果：
```
Symbol table '.symtab' contains 6 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
     1: 0000000000000000     0 FILE    LOCAL  DEFAULT  ABS hello.s
     2: 0000000000000000     0 SECTION LOCAL  DEFAULT    1
     3: 0000000000000000     0 SECTION LOCAL  DEFAULT    2
     4: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT    1 hello
     ;注意下面这行，有了这个GLOBAL的_start，链接器才能找到程序入口
     5: 0000000000000000     0 NOTYPE  GLOBAL DEFAULT    2 _start
```

把 `global _start` 注释掉，
```
Symbol table '.symtab' contains 7 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
     1: 0000000000000000     0 FILE    LOCAL  DEFAULT  ABS hello.s
     2: 0000000000000000     0 SECTION LOCAL  DEFAULT    1
     3: 0000000000000000     0 SECTION LOCAL  DEFAULT    2
     4: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT    1 hello
     5: 000000000000000e     0 NOTYPE  LOCAL  DEFAULT  ABS length
     6: 0000000000000000     0 NOTYPE  GLOBAL DEFAULT    2 _start
```

链接器指的是 ld，ld 的用法：

```
ld -o hello hello.o
```

程序的基本分段：
```
.data : 有初始化值的全局变量；定义常量。 .bss : 没有初始化值的全局变量。
.text : 代码段。
.rodata: 只读数据段。
```

.data 初始化值保存在 elf 数据段内。

.bss 被初始化为 0。

$: 当前内存地址；$$ 当前段内存地址。可在任意段使用，包括 .text。


```
ubuntu@ubuntu-xenial:~$ objdump -d  -M intel hello

hello:     file format elf64-x86-64


Disassembly of section .text:

00000000004000b0 <_start>:
  4000b0:	b8 01 00 00 00       	mov    eax,0x1
  4000b5:	bf 01 00 00 00       	mov    edi,0x1
  ;注意这里的 0x6000d8，和 readelf -s hello 读出来的符号表里的 hello 的内存地址是对应的
  4000ba:	48 be d8 00 60 00 00 	movabs rsi,0x6000d8
  4000c1:	00 00 00
  ; 注意这行，在编译的时候，length 已经被展开成 0xe 了
  4000c4:	ba 0e 00 00 00       	mov    edx,0xe
  4000c9:	0f 05                	syscall
  4000cb:	b8 3c 00 00 00       	mov    eax,0x3c
  4000d0:	48 31 ff             	xor    rdi,rdi
  4000d3:	0f 05                	syscall
```

```
ubuntu@ubuntu-xenial:~$ readelf -s hello

Symbol table '.symtab' contains 10 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
     1: 00000000004000b0     0 SECTION LOCAL  DEFAULT    1
     2: 00000000006000d8     0 SECTION LOCAL  DEFAULT    2
     3: 0000000000000000     0 FILE    LOCAL  DEFAULT  ABS hello.s
     4: 00000000006000d8     0 NOTYPE  LOCAL  DEFAULT    2 hello
     5: 000000000000000e     0 NOTYPE  LOCAL  DEFAULT  ABS length
     6: 00000000004000b0     0 NOTYPE  GLOBAL DEFAULT    1 _start
     7: 00000000006000e6     0 NOTYPE  GLOBAL DEFAULT    2 __bss_start
     8: 00000000006000e6     0 NOTYPE  GLOBAL DEFAULT    2 _edata
     9: 00000000006000e8     0 NOTYPE  GLOBAL DEFAULT    2 _end
```

以 string 方式查看 data 段中的存储内容：

```
ubuntu@ubuntu-xenial:~$ readelf -p .data ./rev

String dump of section '.data':
  [     1]  Hello World

```
