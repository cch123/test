#include <iostream>
#include <array>
using namespace std;

int main(int argc, char * argv[]) {
    array<int, 5> arr = {3, 1, 2, 5, 4};
    sort(begin(arr), end(arr));
    for(int i=0; i<arr.size();i++) {
        cout << arr[i] << endl;
    }

    sort(begin(arr), end(arr), greater<int>{});
    for(int i=0; i<arr.size();i++) {
        cout << arr[i] << endl;
    }
    return 0;
}
