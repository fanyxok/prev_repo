// #include <utility>
// #include <functional>
// #include "skiplist.hpp"
#include "skiplist.h"
#include <iostream>

using namespace std;

int main() {
	int a = 10;
	skiplist<int, int, std::less<int>> skipl(a);
	cout << "=========== test for instantiation of skiplist passed ============" << endl;
	// skipl.insert(2, 1);
	// cout << "insert(1,1): " <<  skipl.insert(1, 1).second << endl;
	// cout << "insert(1,1): " <<  skipl.insert(1, 1).second << endl;
	// skipl.find(1);
	// cout << "erase(1): " <<  skipl.erase(1) << endl;
	// cout <<skipl.level() << endl;
	// cout << skipl.size() << endl;
	// cout << "=========== test for skiplist::insert(1) passed ===================" << endl;
	// skipl.insert(2, 1);
	// cout << "=========== test for skiplist::insert(2) passed ===================" << endl;
	// skipl.insert(3, 1);
	// cout << "=========== test for skiplist::insert(3) passed ===================" << endl;
	// skipl.insert(4, 1);
	// cout << "=========== test for skiplist::insert(4) passed ===================" << endl;
	// skipl.insert(5, 1);
	// cout << "=========== test for skiplist::insert(5) passed ===================" << endl;
	// skipl.insert(6, 1);
	// cout << "=========== test for skiplist::insert(6) passed ===================" << endl;

	for (int i=0;i<1000;i++) {
		cout << skipl.insert(i, i*2).second <<endl;
	}
	cout << skipl.level() << endl;
	// skiplist<int, int, std::less<int>>::iterator b = skipl.insert(0,0).first;
	// for (int i=0;i<10;i++) {
	// 	cout << b++->first <<endl;
	// }
	cout << "=========== test for skiplist::insert()        passed =============" << endl;


	cout << "=========== test for skiplist::find()          passed ============" << endl;
	for (int i=0;i<1000;i++) {
		cout << skipl.erase(i) << endl;
	}
	for (int i=0;i<10;i++) {
		cout << skipl.erase(i) << endl;
	}
	skipl.find(3);
	cout << "=========== test for skiplist::erase()         passed =============" << endl;
	cout << skipl.size() << endl;
	cout << "+++++++++++++++++++++++++ end of all tests ++++++++++++++++++++++++" << endl;
	return 0;
}
