//
// Created by chunlei zhang on 29/08/2017.
//

#include <unordered_map>
#include <iostream>
#include <string>

using namespace std;

int main(){

    unordered_map<string, int> m;
    m["zhang"] = 37;
    m.insert(make_pair("lubo", 34));

    if (m.find("zhang") != m.end()) {
        cout << "found key zhang with value: " << m.find("zhang")->second << endl;
        m.erase("zhang");
    }

    for (auto &pos : m) {
        cout << pos.first << " : " << pos.second << endl;
    }
}