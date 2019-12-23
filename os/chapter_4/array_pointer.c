#include "stdio.h"
#include "stdlib.h"

int main() {
    long a[] = {1,2,3,4,5};
    long (*b)[5] = &a;
    // long (*b)[] = &a; will complain error
    b = b+1;
    // the pointer move to address after the end of a
    printf("%x\n", b);
    printf("%x\n", a);
    long (*c)[5] = &a;
    
    // 默认移动大小是指向目标的 sizeof 决定的
    // 如果想要以 int 方式移动，可对该数组指针进行强制转换
    c = (int *)c + 1;
    printf("%x\n", c);
}
