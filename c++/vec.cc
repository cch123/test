#include<iostream>
#include<vector>

using namespace std;

void change_v(vector<int> v) {
    v.push_back(4);
    v.push_back(4);
    v.push_back(4);
    cout << v.size() << endl;
}

int main() {
    auto v = vector<int>();
    v.push_back(1);
    v.push_back(1);
    v.push_back(1);
    cout << v.size() << endl;
    change_v(v);
    cout << v.size() << endl;
}