global _start

section .data
    SYS_WRITE equ 1
    STD_IN equ 1
    SYS_EXIT equ 60
    EXIT_CODE equ 0

    NEW_LINE db 0xa
    WRONG_ARG db "Must be two cmdline arguments", 0xa

section .text
_start:
	pop rcx
	cmp rcx, 3
	jne argcError

	add rsp, 8

	pop rsi
	call str_to_int
	mov r10, rax

	pop rsi
	call str_to_int
	mov r11, rax

	add r10, r11
    mov rax, r10
    xor r12, r12
    jmp int_to_str

argcError:
    mov rax, 1
    mov rdi, 1
    mov rsi, WRONG_ARG
    mov rdx, 34
    syscall
    jmp exit

str_to_int:
    xor rax, rax
    mov rcx, 10

next:
    cmp [rsi], byte 0
    je return_str
    mov bl,[rsi]
    sub bl, 48
    mul rcx
    add rax, rbx
    inc rsi
    jmp next

return_str:
    ret

int_to_str:
		mov rdx, 0
		mov rbx, 10
		div rbx
		add rdx, 48
		add rdx, 0x0
		push rdx
		inc r12
		cmp rax, 0x0
		jne int_to_str
		jmp print

print:
	;;;; calculate number length
	mov rax, 1
	mul r12
	mov r12, 8
	mul r12
	mov rdx, rax

	;;;; print sum
	mov rax, SYS_WRITE
	mov rdi, STD_IN
	mov rsi, rsp
	;; call sys_write
	syscall

    jmp exit


exit:
	mov rax, SYS_EXIT
	mov rdi, EXIT_CODE
	syscall