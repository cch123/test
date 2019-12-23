#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/epoll.h>

#define BUF_SIZE 1024
#define EPOLL_SIZE 50

void error_handling(char * msg);

int main(int argc, char * argv[]) {
    int sock_fd, conn_fd;
    struct sockaddr_in serv_addr, client_addr;
    socklen_t addr_size;
    int str_len, i;
    char buf[BUF_SIZE];
    struct epoll_event * ep_events;
    struct epoll_event event;

    int epfd, event_cnt;

    if(argc!=2) {
        printf("Usage : %s <port>\n", argv[0]);
        exit(1);
    }

    sock_fd = socket(PF_INET, SOCK_STREAM, 0);
    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_addr.sin_port = htons(atoi(argv[1]));

    if(bind(sock_fd, (struct sockaddr*)&serv_addr, sizeof(serv_addr)) == -1) {
        error_handling("bind error");
    }

    if(listen(sock_fd, 5) == -1) {
        error_handling("listen error");
    }

    epfd = epoll_create(EPOLL_SIZE);
    ep_events = malloc(sizeof(struct epoll_event) * EPOLL_SIZE);

    event.events = EPOLLIN;
    event.data.fd = sock_fd;
    epoll_ctl(epfd, EPOLL_CTL_ADD, sock_fd, &event);

    while(1) {
        event_cnt = epoll_wait(epfd, ep_events, EPOLL_SIZE, -1);
        if(event_cnt == -1) {
            puts("epoll_wait error");
            break;
        }

        for(i=0; i<event_cnt; i++) {
            if(ep_events[i].data.fd == sock_fd) {
                addr_size = sizeof(client_addr);
                conn_fd = accept(sock_fd, (struct sockaddr*)&client_addr, &addr_size);
                event.events = EPOLLIN;
                event.data.fd = conn_fd;
                epoll_ctl(epfd, EPOLL_CTL_ADD, conn_fd, &event);
                printf("connected client : %d\n", conn_fd);
            } else {
                str_len = read(ep_events[i].data.fd, buf, BUF_SIZE);
                if(str_len == 0) { // close request EOF?
                    epoll_ctl(epfd, EPOLL_CTL_DEL, ep_events[i].data.fd, NULL);
                    close(ep_events[i].data.fd); //close(conn_fd);
                    printf("closed client: %d\n", ep_events[i].data.fd);
                } else {
                    write(ep_events[i].data.fd, buf, str_len); // echo result!
                }
            }
        }
    }
    close(sock_fd);
    close(epfd);
    return 0;

}

void error_handling(char * msg) {
    fputs(msg, stderr);
    fputc('\n', stderr);
    exit(1);
}
