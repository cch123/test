#include <iostream>
#include <unordered_map>
using namespace std;

int main() {
    unordered_map<int,int> m = {{1,2}, {3,2}};
    m[12] = 1;
    m[14] = 1;
    for(auto &x: m) {
        cout << x.first << " " << x.second << endl;
    }
}