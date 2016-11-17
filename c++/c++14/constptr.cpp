#include <iostream>
using namespace std;

int main() {
    const int a = 1;
    auto b = a;
    b = 4;
    auto c = &b;
    auto d = &a;

    cout << typeid(c).name() << endl;
    cout << typeid(d).name() << endl;
    return 0;
}
