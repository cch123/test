#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
void error_handling(char * msg);

int main(int argc, char * argv[]) {
    int socket_fd;
    struct sockaddr_in serv_addr;
    char msg[30];
    int str_len;

    if(argc!=3) {
        printf("Usage : %s <IP> <port>\n", argv[0]);
        exit(1);
    }

    socket_fd = socket(PF_INET, SOCK_STREAM, 0);
    if(socket_fd == -1) {
        error_handling("socket error");
    }

    memset(&serv_addr, 0, sizeof(serv_addr));

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = inet_addr(argv[1]);
    serv_addr.sin_port = htons(atoi(argv[2]));

    if(connect(socket_fd, (struct sockaddr*)&serv_addr, sizeof(serv_addr)) == -1) {
        error_handling("connect error");
    }

    str_len = read(socket_fd, msg, sizeof(msg)-1);
    if(str_len == -1) {
        error_handling("read error");
    }

    printf("Message from server : %s\n", msg);
    close(socket_fd);
    return 0;

}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}
