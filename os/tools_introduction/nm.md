man nm
```
    For each symbol, nm shows:

       ·   The symbol value, in the radix selected by options (see below), or hexadecimal by default.

       ·   The symbol type.  At least the following types are used; others are, as well, depending on
       the object file format.  If lowercase, the symbol is usually local; if uppercase, the
       symbol is global (external).  There are however a few lowercase symbols that are shown for
       special global symbols ("u", "v" and "w").
       注意，大小写只是表明该变量是 external 还是 internal，在 c 里文件作用域级的变量是 static，即 internal 的
       

          "B"
           "b" The symbol is in the uninitialized data section (known as BSS).

           "C" The symbol is common.  Common symbols are uninitialized data.
               When linking, multiple common symbols may appear with the same
               name.  If the symbol is defined anywhere, the common symbols
               are treated as undefined references.

           "D"
           "d" The symbol is in the initialized data section.
```

在 C 语言中，全局变量，初始化全局变量，静态变量所在的内存区可以使用 nm 来方便的查看：
```
#include "stdio.h"
#include "stdlib.h"

int x = 1;
int y;
int main() {
static int z=1;
}
```

nm a.out

```
0000000000601030 D x
000000000060103c B y
0000000000601034 d z.280
```

除了 nm 的简单判断之外，可以结合 readelf -S 来判断该变量的 address 是不是“确实”落在某一个段上了。

```
  [23] .got              PROGBITS         0000000000600ff8  00000ff8
       0000000000000008  0000000000000008  WA       0     0     8
  [24] .got.plt          PROGBITS         0000000000601000  00001000
       0000000000000020  0000000000000008  WA       0     0     8
  [25] .data             PROGBITS         0000000000601020  00001020
       0000000000000018  0000000000000000  WA       0     0     8
  [26] .bss              NOBITS           0000000000601038  00001038
       0000000000000008  0000000000000000  WA       0     0     4

```