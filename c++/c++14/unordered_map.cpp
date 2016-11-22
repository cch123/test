#include <iostream>
#include <unordered_map>
using namespace std;

int main() {
    unordered_map<int, string>m;
    m.insert(make_pair(100, "fuck"));
    m.insert(make_pair(1, "fuck1"));
    m.insert(make_pair(20, "xxj"));
    m.insert(make_pair(10, "jjno"));
    for(auto p : m) {
        cout << p.first << endl;
        cout << p.second<< endl;
    }
    cout << "bucket_count" << m.bucket_count() << endl;
    cout << "max bucket count" << m.max_bucket_count() << endl;
    cout << "load factor" << m.load_factor() << endl;
    cout << "max load factor" << m.max_load_factor() << endl;
    return 0;

}
