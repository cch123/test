#include "stdio.h"

int main() {
    setlinebuf(stdout);
    printf("abc");
    printf("abc");
    printf("abc");
    printf("abc");
    sleep(10);
    return 1;
}
