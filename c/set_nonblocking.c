#include "stdio.h"
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>

int main() {
    //STDIN_FILENO
}

int setnonblocking(int fd) {
    int old_option = fcntl(fd, F_GETFL);
    int new_option = old_option | O_NONBLOCK;
    fcntl(fd, F_SETFL, new_option);
    return old_option;
}

int
setnonblocking2(int fd) {
    int  nb;
    nb = 1;
    return ioctl(fd, FIONBIO, &nb);
}
