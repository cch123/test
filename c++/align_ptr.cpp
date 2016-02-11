#include <iostream>
#include "stdio.h"
#include "stdlib.h"
using namespace std;
typedef unsigned long uintptr_t;
#define ngx_align_ptr(p, a)                                                   \
          (u_char *) (((uintptr_t) (p) + ((uintptr_t) a - 1)) & ~((uintptr_t) a - 1))

int main() {
    void * a = malloc(sizeof(char)*100);
    a = (char *)a + 1;
    cout<< a<<endl;
    a = ngx_align_ptr(a, 8);
    cout<< a<<endl;
    return 0;
}
