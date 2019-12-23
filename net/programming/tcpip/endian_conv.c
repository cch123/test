#include <stdio.h>
#include <arpa/inet.h>
/*
 * 因为intel和amd的系统大多是小端系统，所以和网络字节序是反着的
 * 而需要关注大小端，做字节序转换只有当填充sock_addr和sock_addr_in里的
 * address和port的时候需要关注，
 * 接收和发送数据系统已经帮助完成了这个过程
 */


int main(int argc, char *argv[]) {
    unsigned short host_port = 0x1234;
    unsigned short net_port;
    unsigned long host_addr = 0x12345678;
    unsigned long net_addr;

    net_port = htons(host_port);
    net_addr = htonl(host_addr);
    printf("Host ordered port is %#x\n", host_port);
    printf("Network ordered port is %#x\n", net_port);
    printf("Host ordered address is %#lx\n", host_addr);
    printf("Network ordered address is %#lx\n", net_addr);
}
