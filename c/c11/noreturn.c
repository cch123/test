#include <stdio.h>
#include <stdnoreturn.h>
// c11 增加了 _Alignof _Alignas _Atomic _Generic _Noreturn _Static_assert
// _Thread_local 几个新关键字

_Noreturn fun() {
    printf("in fun\n");
}

int main() {
    printf("start main\n");
    fun();
    printf("end main\n");
}
