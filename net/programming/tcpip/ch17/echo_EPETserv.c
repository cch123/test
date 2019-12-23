#include <fcntl.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/epoll.h>

#define BUF_SIZE 1024
#define EPOLL_SIZE 50

void error_handling(char *);
void setnonblockingmode(int);

int main(int argc, char * argv[]) {
    int sock_fd, conn_fd;
    struct sockaddr_in serv_addr, client_addr;
    socklen_t addr_size;
    int str_len, i;
    char buf[BUF_SIZE];

    struct epoll_event *ep_events = malloc(sizeof(struct epoll_event) * EPOLL_SIZE);

    struct epoll_event event;

    int epfd, event_cnt;
    if(argc != 2) {
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
    //也可以像下面这样，注意，如果flag填错了的话可能会导致epoll_wait失败
    //epfd = epoll_create1(EPOLL_CLOEXEC);
    setnonblockingmode(sock_fd);
    event.events = EPOLLIN;
    event.data.fd = sock_fd;
    //有个疑问，感觉内部应该是对event做过一次拷贝
    //要不然反复修改event肯定是会有问题的
    epoll_ctl(epfd, EPOLL_CTL_ADD, sock_fd, &event);

    while(1) {
        //最后一个参数-1表示无限等待，本身的含义是timeout
        event_cnt = epoll_wait(epfd, ep_events, EPOLL_SIZE, -1);
        if(event_cnt == -1) {
            puts("epoll_wait failed");
            break;
        }

        puts("epoll_wait returned");
        for(i=0;i<event_cnt;i++) {
            if(ep_events[i].data.fd == sock_fd) {
                addr_size = sizeof(client_addr);
                conn_fd = accept(sock_fd, (struct sockaddr*)&client_addr, &addr_size);
                //连接的fd也要设置成nonblocking
                setnonblockingmode(conn_fd);
                //边缘触发
                event.events = EPOLLIN|EPOLLET;
                event.data.fd = conn_fd;
                epoll_ctl(epfd, EPOLL_CTL_ADD, conn_fd, &event);
                printf("connected client: %d\n", conn_fd);
            } else {
                //认为是要传输数据
                //其实这里的判断不严谨，还缺少很多错误条件判断
                while(1) {
                    //因为边缘模式只触发一次EPOLLIN，所以需要一直读到完毕为止
                    str_len = read(ep_events[i].data.fd, buf, BUF_SIZE);
                    if(str_len == 0) { //说明需要关闭了
                        epoll_ctl(epfd, EPOLL_CTL_DEL, ep_events[i].data.fd, NULL);
                        close(ep_events[i].data.fd);
                        printf("closed client : %d\n", ep_events[i].data.fd);
                        break;
                    } else if(str_len < 0){
                        //说明出错了，需要判断是不是EAGAIN
                        if(errno == EAGAIN) {
                            break;
                        }
                    } else {
                        //str_len > 0
                        write(ep_events[i].data.fd, buf, str_len);
                    }
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

void setnonblockingmode(int fd) {
    int flag = fcntl(fd, F_GETFL, 0);
    fcntl(fd, F_SETFL, flag|O_NONBLOCK);
}
