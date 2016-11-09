#include <iostream>
using namespace std;

int main() {
    int && ref = 1+2;
    ref +=1;
    int a = 1;
    cout << "value is " << ref << " and ptr is " << &ref << endl;
    cout << "value is " << a << " and ptr is " << &a<< endl;
    return 0;
}
