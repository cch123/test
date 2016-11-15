#include<iostream>
using namespace std;
template<typename T>
auto add(const T a, const T b) {
    cout <<"add" <<typeid(a).name() <<typeid(b).name()<< endl;
    return a + b;
}

//不知道在C艹里有没有更好的办法来实现type打印的？
int main() {
    int a = 1, b = 100;
    auto c = add(a, b);
    cout << c << endl;
    cout << typeid(c).name() << endl;
    double ax = 1.1, bx = 100.0;
    auto cx = add(ax, bx);
    cout << cx << endl;

    cout << typeid(ax).name() << endl;
    cout << typeid(double).name() << endl;
    cout << typeid(cx).name() << endl;

    //判断是否是某个变量的type
    if(typeid(cx) == typeid(double)) {
        cout << "yes it is double" << endl;
    }
    return 0;
}
