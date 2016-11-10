#include <iostream>
using namespace std;

tuple<int, bool, string, double> return_mul() {
    return make_tuple(1, false, "fuck", 1.1);
}

int main() {
    int a;
    bool b;
    string c;
    double d;
    tie(a,b,c,d) = return_mul();
    cout << a << endl;
    cout << b << endl;
    cout << c << endl;
    cout << d << endl;
    return 0;
}
