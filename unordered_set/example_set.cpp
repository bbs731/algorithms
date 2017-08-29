//
// Created by chunlei zhang on 29/08/2017.
//
#include <unordered_set>
#include <string>
#include <iostream>

using namespace std;

// usuage of insert, erase, find

int main(){
    unordered_set<string> s = {"chunlei", "zhang"};
    s.insert({"wang", "lubo"});

    if (s.find("lubo") != s.end()) {
        cout << "found \"lubo\" is in the hashset!\n";
    }

    s.erase("zhang");

    unordered_set<string>::iterator pos;
    for (pos = s.begin(); pos != s.end(); pos++) {
        cout << *pos << endl;
    }


}

