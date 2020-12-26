#include <stdio.h>
#include <unistd.h>
#include <signal.h>
 
void timeout(int sig)
{
    if(sig==SIGALRM){
        puts("time out");
    }
    alarm(2);
}

int main(void)
{
    int i;
    struct sigaction act;  //声明结构体类型
    act.sa_handler = timeout;  //保存函数指针
    sigemptyset(&act.sa_mask);  //将sa_mask初始化为0
    act.sa_flags = 0;
    sigaction(SIGALRM, &act, 0);
    alarm(2);
    for(i=0; i<3; i++){
        puts("wait...");
        sleep(100);
    }
    return 0;
}

