#include <iostream>
#include <vector>
using namespace std;

int main(int argc, char * argv[]) {
    vector<int> a = {1,2,3,4};
    vector<int> b = {5,6,7,8};
    // 两个容器必须有相同的大小
    cout << "swap one by one use swap_ranges" << endl;
    swap_ranges(begin(a), end(a), begin(b));

    for(int i=0;i<a.size();i++) {
        cout << a[i] << endl;
    }

    for(int i=0;i<b.size();i++) {
        cout << b[i] << endl;
    }

    swap(a,b);
    cout << "swap at once" << endl;
    for(int i=0;i<a.size();i++) {
        cout << a[i] << endl;
    }

    for(int i=0;i<b.size();i++) {
        cout << b[i] << endl;
    }

    return 0;
}
