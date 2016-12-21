#include "sds.h"
#include "stdio.h"

int main() {
    char a[] = "君の名";
    sds mySDS = sdsnew(a);
    printf("%s", mySDS);
}
