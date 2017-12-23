#include "point.h"
#include <stdlib.h>

// .c 中定义的结构体，在 .h 中无法访问，且对外部不可见
struct Point {
    double x;
    double y;
};

struct Point * makePoint(double x, double y) {
    struct Point * p = malloc(sizeof(struct Point));
    p->x = x;
    p->y = y;
    return p;
}
