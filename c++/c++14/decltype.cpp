#include <iostream>
using namespace std;

int main() {
    int a = 1, b = 2;
    decltype(a) c = a + b;
    cout << c << endl;
    cout << typeid(c).name() << endl;
    auto d = a + c;
    cout << d << endl;
    cout << typeid(d).name() << endl;
    return 0;
}
