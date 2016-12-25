#include <iostream>
#include <algorithm>
#include <vector>
#include <cmath>

using namespace std;
class Person {
 public:
  Person(string name, int age) {
    this->age = age;
    this->name = name;
  }
  int age;
  string name;
};

int main() {
  cout << "hello world" << endl;
  vector<int> arr;
  arr.push_back(1);
  arr.push_back(100);
  cout << arr.size() << endl;
  Person *p = new Person("alex", 23);
  cout << p->age << endl;
  cout << p->name << endl;
  arr.push_back(32);

  cout << arr[0] << endl;
  cout << arr.at(1) << endl;

  cout << "hello" << endl;
}