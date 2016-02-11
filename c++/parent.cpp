#include <iostream>
using namespace std;

class Father{
public:
    virtual void test(){}
    void tt() {
        test();
    }
};

class Son : public Father{
public:
    void test() { cout << "son test" << endl;}

};

int main() {
    Father * a = new Son();
    a->tt();
}

