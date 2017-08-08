// 在struct末尾插入变长类型的做法
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int length;
    unsigned char data[];
} str;

int main() {
    str *s = malloc(sizeof(str) + 10);
    s->length = 10;
    memset(s->data, 0, 10);

    memcpy(s->data, "hello", 5);
    printf("%d,%s\n", s->length, s->data);
    printf("%lu \n", sizeof(str));
    return 0;
}

