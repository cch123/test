#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
void error_handling(char * msg);

#define BUFSIZE 1024

int main(int argc, char * argv[]) {
    int socket_fd;
    struct sockaddr_in serv_addr;
    char msg[BUFSIZE];
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
    } else {
        puts("Connected......\n");
    }


    while(1) {
        fputs("Input message(Q to quit):", stdout);
        fgets(msg, BUFSIZE, stdin);

        if(!strcmp(msg, "q\n") || !strcmp(msg, "Q\n")) {
            break;
        }

        //tcp没有数据边界，这么发和读是一种取巧的方式
        write(socket_fd, msg, strlen(msg));
        str_len = read(socket_fd, msg, BUFSIZE - 1);
        msg[str_len] = 0;
        printf("Message from server : %s\n", msg);
    }
    close(socket_fd);
    return 0;

}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}
