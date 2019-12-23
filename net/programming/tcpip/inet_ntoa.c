#include <stdio.h>
#include <arpa/inet.h>
#include <string.h>

int main(int argc, char *argv[]) {
    struct sockaddr_in addr1, addr2;
    char *str_ptr;
    char str_arr[20];
    addr1.sin_addr.s_addr = htonl(0x1020304);
    addr2.sin_addr.s_addr = htonl(0x1010101);

    str_ptr = inet_ntoa(addr1.sin_addr);
    strcpy(str_arr, str_ptr);
    printf("Dotted-Decimal notation : %s\n", str_ptr);

    //这是为了演示，其它系统或者函数调用返回的char*类型
    //一定要在自己的内存空间里做一次拷贝，否则无法保证安全性
    inet_ntoa(addr2.sin_addr);
    printf("Dotted-Decimal notation2: %s\n", str_ptr);
    printf("Dotted-Decimal notation3: %s\n", str_arr);
    return 0;
}
