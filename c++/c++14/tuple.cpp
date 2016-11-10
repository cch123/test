#include <iostream>
using namespace std;

int main() {
    //可以直接搞出来很多变量！
    tuple<int, bool, float, string> result = make_tuple(128, true, 1.1f, "no");
    auto value = get<0>(result);
    //通过下标去获取对应的元素！
    //有点像void*类型的数组，但明显方便太多了
    auto value1 = get<1>(result);
    auto value2 = get<2>(result);

    int obj1;
    bool obj2;
    float obj3;
    tie(obj1, obj2, obj3) = make_tuple(128, true, 1.2f);
    cout << obj1 << " " << obj2 << " " << obj3 << endl;
    cout << value << endl;
    cout << value1 << endl;
    cout << value2 << endl;
}
