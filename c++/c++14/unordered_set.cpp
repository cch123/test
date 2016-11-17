#include <iostream>
#include <unordered_set>
using namespace std;

int main() {
    unordered_set<int> foo;
    foo.insert(1);
    foo.insert(2);
    foo.insert(4);
    foo.insert(11);
    for(auto p: foo) {
        cout << p << endl;
    }

    cout << "bucket_count" << foo.bucket_count() << endl;
    cout << "max bucket count" << foo.max_bucket_count() << endl;
    cout << "load factor" << foo.load_factor() << endl;
    cout << "max load factor" << foo.max_load_factor() << endl;
    return 0;
}
