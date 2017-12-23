#include "point.h"
#include <stdio.h>

int main() {
    struct Point * p = makePoint(1.1, 2.2);
    printf("%p", p);
    // printf("%f", p->x);
    // 这样的无法访问到 p->x 的，编译会报错
    return 0;
}