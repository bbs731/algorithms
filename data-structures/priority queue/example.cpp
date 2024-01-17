//
// Created by chunlei zhang on 29/08/2017.
//

#include <queue>
#include <cstdio>

using namespace std;

int main(){

    // default priority_queue using
    // priority_queue<int, vector<int>, less<int> > pque;
    // max heap, 取值是取最大值。
    // 如果想要 min heap, use greater<int> instead
    priority_queue<int> pque;

    pque.push(3);
    pque.push(5);
    pque.push(1);

    while (!pque.empty()) {
        printf("%d\n", pque.top());
        pque.pop();
    }

}


