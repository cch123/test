#include "sds.h"
#include "stdio.h"
#include "string.h"

int main() {
    sds a = sdsnewlen("set a \"ohnot \\x00hisisthe value\"", 100);
    int *cnt;
    sds * arr = sdssplitargs(a, cnt);
    printf("argc is %d\n", *cnt);
    printf("argv[0] is %s\n", arr[0]);
    printf("argv[1] is %s\n", arr[1]);
    printf("argv[2] is %s\n", arr[2]);
    printf("sdslen is %zu\n", sdslen(arr[2]));
    printf("strlen is %lu\n", strlen(arr[2]));
}
