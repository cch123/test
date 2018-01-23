#include "stdio.h"
#include "stdlib.h"

int main() {
    char * x = NULL;
    void * a  = x;
    int c = a == x? 1:0;
    printf("%d",c);
}


