#include <fcntl.h>
void main() {
    int dev_random = open("/dev/urandom", O_RDONLY);
    printf("%d", dev_random);
    int NUMBYTES = 8;
    unsigned char rand_data[NUMBYTES];
    int r = read(dev_random, &rand_data, NUMBYTES);
    printf("%d\n", r);
    for(int i=0;i<NUMBYTES;i++) {
        printf("%02X\n", rand_data[i]);
    }

}
