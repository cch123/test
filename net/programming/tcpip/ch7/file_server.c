#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define BUF_SIZE 30

void error_handling(char * message);

int main(int argc, char * argv[]) {
    int serv_fd, conn_fd;

    FILE * fp;
    char buf[BUF_SIZE];
    int read_cnt;

    struct sockaddr_in serv_addr, clnt_addr;

    socklen_t clnt_addr_size;

    if(argc!=2) {
        printf("Usage: %s <port>\n",argv[0]);
        exit(1);
    }

    fp = fopen("file_server.c", "rb");
    serv_fd = socket(PF_INET, SOCK_STREAM, 0);

    memset(&serv_addr, 0 , sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_addr.sin_port = htons(atoi(argv[1]));

    bind(serv_fd, (struct sockaddr *)&serv_addr, sizeof(serv_addr));
    listen(serv_fd, 5);

    clnt_addr_size = sizeof(clnt_addr);
    conn_fd = accept(serv_fd, (struct sockaddr *)&clnt_addr, &clnt_addr_size);

    while(1) {
        read_cnt = fread((void*)buf, 1, BUF_SIZE, fp);
        if(read_cnt < BUF_SIZE) {
            write(conn_fd, buf, read_cnt);
            break;
        }
        write(conn_fd, buf, BUF_SIZE);
    }

    shutdown(conn_fd, SHUT_WR);
    read(conn_fd, buf, BUF_SIZE);
    printf("Message from client: %s\n", buf);

    fclose(fp);
    close(conn_fd);close(serv_fd);
    return 0;
}

void error_handling(char * message) {
    fputs(message, stderr);
    fputc('\n',stderr);
    exit(1);
}
