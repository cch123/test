#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

int main() {
    vector<int> a;
    a.push_back(122);
    auto sz = a.size();
    cout << sz << endl;

    unordered_map<int, string>m;
    m.insert(make_pair(100, "fuck"));
    m.insert(make_pair(1, "fuck1"));
    m.insert(make_pair(20, "xxj"));
    m.insert(make_pair(10, "jjno"));
    for(auto p : m) {
        cout << p.first << endl;
        cout << p.second<< endl;
    }
    return 0;
}
