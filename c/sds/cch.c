#include "sds.h"
#include "stdio.h"
#include "string.h"

int main() {
    sds a = sdsnewlen("set a \"ohnot \\x00hisisthe value\"", 100);
    int *cnt;
    sds * arr = sdssplitargs(a, cnt);
    printf("%d\n", *cnt);
    printf("%s\n", arr[0]);
    printf("%s\n", arr[1]);
    printf("%s\n", arr[2]);
    printf("%zu\n", sdslen(arr[2]));
    printf("%lu\n", strlen(arr[2]));
}
