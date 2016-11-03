#include <stdio.h>

int a() {
    goto ERROR_HANDLE;
    return 1;
ERROR_HANDLE:
    printf("aaa");
    return 0;
}

int main() {
    goto ERROR_HANDLE;
ERROR_HANDLE:
    printf("1");

   return 0;
}

