Overview
---

ensure that the stack frame layout is like the following
```
sym+8(FP)
-------
sym+0(FP)
------ FP
 return address 8 bytes
------
 frame pointer 8 bytes
------ SP
sym-8(SP)
-----
sym-16(SP)
-----
```
