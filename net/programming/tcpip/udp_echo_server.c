#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define BUF_SIZE 1024

void error_handling(char * message);

int main(int argc, char * argv[]) {
    int socket_fd;
    char msg[BUF_SIZE];
    int str_len;
    socklen_t client_addr_size;
    struct sockaddr_in serv_addr, client_addr;
    if(argc != 2) {
        printf("Usage : %s <port>\n", argv[0]);
        exit(1);
    }

    socket_fd = socket(PF_INET, SOCK_DGRAM, 0);
    if(socket_fd == -1) {
        error_handling("udp socket create error");
    }

    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_addr.sin_port = htons(atoi(argv[1]));

    if(bind(socket_fd, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) == -1) {
        error_handling("bind error");
    }

    while(1) {
        client_addr_size = sizeof(client_addr);
        str_len = recvfrom(socket_fd, msg, BUF_SIZE, 0,
                (struct sockaddr*)&client_addr, &client_addr_size);
        sendto(socket_fd, msg, str_len, 0,
                (struct sockaddr*)&client_addr, client_addr_size);
    }
    close(socket_fd);
    return 0;

}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}
