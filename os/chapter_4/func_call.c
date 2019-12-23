#include <stdio.h>
#include <stdlib.h>

int add(int a, int b) {
    return a + b;
}

int main(int argv, char ** argc) {
    int a = 1;
    int b = 2;
    printf("%d\n", add(a,b));
}
