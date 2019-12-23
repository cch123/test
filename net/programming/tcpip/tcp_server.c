#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
void error_handling(char *msg);

int main(int argc, char *argv[]) {
    int socket_fd, connect_fd;

    struct sockaddr_in serv_addr;
    struct sockaddr_in client_addr;
    socklen_t client_addr_size;

    char msg[] = "hello world";

    if(argc!=2) {
        printf("Usage: %s <port>\n", argv[0]);
        exit(1);
    }

    socket_fd = socket(PF_INET, SOCK_STREAM, 0);
    if(socket_fd == -1) {
        error_handling("socket() error");
    }

    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_addr.sin_port=htons(atoi(argv[1]));

    if(bind(socket_fd, (struct sockaddr*)&serv_addr, sizeof(serv_addr)) == -1) {
        error_handling("bind error");
    }

    if(listen(socket_fd, 5) == -1) {
        error_handling("listen error");
    }

    client_addr_size = sizeof(client_addr);

    connect_fd = accept(socket_fd, (struct sockaddr*)&client_addr, &client_addr_size);
    if(connect_fd == -1) {
        error_handling("accept error");
    }

    write(connect_fd, msg, sizeof(msg));
    close(connect_fd);
    close(socket_fd);
    return 0;
}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}

