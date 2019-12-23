global _start

section .text

_start:
	mov rcx, 236
	mov rax, 123

s:
	add  rax, 123
	loop s; 在rcx 变为 0 之前，loop 会一直执行

	mov rax, 60; 'exit' syscall number
	xor rdi, rdi
	syscall

