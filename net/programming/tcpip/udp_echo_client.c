#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define BUF_SIZE 1024

void error_handling(char * message);

int main(int argc, char* argv[]) {
    int socket_fd;
    char msg[BUF_SIZE];
    int str_len;
    socklen_t addr_size;

    struct sockaddr_in serv_addr, from_addr;

    if(argc!=3) {
        printf("Usage : %s <IP> <port>\n", argv[0]);
        exit(1);
    }

    socket_fd = socket(PF_INET, SOCK_DGRAM, 0);

    if(socket_fd == -1) {
        error_handling("socket get error");
    }

    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = inet_addr(argv[1]);
    serv_addr.sin_port = htons(atoi(argv[2]));

    while(1) {
        fputs("Insert msg(q to quit): ", stdout);
        /*
         *char *fgets(string, count, stream) - input string from a stream
         *
         *Purpose:
         * get a string, up to count-1 chars or '\n', whichever comes first,
         * append '\0' and put the whole thing into string. the '\n' IS included
         * in the string. if count<=1 no input is requested. if EOF is found
         * immediately, return NULL. if EOF found after chars read, let EOF
         * finish the string as '\n' would.
         */
        fgets(msg, sizeof(msg), stdin);
        if(!strcmp(msg, "q\n") || !strcmp(msg, "Q\n")) {
            break;
        }

        sendto(socket_fd, msg, strlen(msg), 0,
                (struct sockaddr*)&serv_addr, sizeof(serv_addr));
        addr_size = sizeof(from_addr);
        str_len = recvfrom(socket_fd, msg, BUF_SIZE, 0,
                (struct sockaddr*)&from_addr, &addr_size);
        msg[str_len] = 0;
        printf("msg from server : %s", msg);
    }
    close(socket_fd);
    return 0;

}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}
