go build -gcflags "-N -l" xxx.go

gdb xxx

l=>列出代码

b `line_number` 列出行号

r 运行程序，在断点处会自动停止

i locals
