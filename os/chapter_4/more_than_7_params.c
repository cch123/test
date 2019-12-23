#include "stdio.h"
#include "stdlib.h"

void test(long a, long b, long c, long d, long e, long f, long g, long h) { 
    printf( "%ld\n", a+b+c+d+e+f+g+h);
}

int main() {
    test(1,2,3,4,5,6,7,8);
}
