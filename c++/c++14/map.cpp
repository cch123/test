#include <iostream>
#include <map>
using namespace std;

int main(int argc, char * argv[]) {
    map<string, int> m = {{"a", 1}, {"b", 2}, {"c", 3}};
    cout << m.count("b") << endl;
    cout << m["b"] << endl;
    return 0;
}
