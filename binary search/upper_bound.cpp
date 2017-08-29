//
// Created by chunlei zhang on 29/08/2017.
//

#include <cstdio>

// 当 v 存在时， 返回最后一个位置的后面一个位置。
// 如果所找 v 不存在， 返回一个下标位置，在此处插入 v, 数组还是有序的
int upper_bound(int *A, int x, int y, int v) {
    int m;
    while (x < y) {  //search [x, y)
        m = x + (y - x) / 2;

        if (A[m] <= v) {
            x = m + 1;
        } else {
            y = m;
        }
    }
    return x;
}


int upper_bound_implicit(int *A, int x, int y, int v) {
    int m;
    while (x < y) {
        m = x + (y - x) / 2;

        if (A[m] > v) {
            y = m;
        } else if (A[m] == v) {
            x = m + 1;
        } else if (A[m] < v) {
            x = m + 1;
        }
    }

    return x;
}

int main() {
    int f;
    int B[8] = {0, 0, 2, 2, 2, 4, 5};
    f = upper_bound(B, 0, 8, 2);
    printf(" return value is %d\n", f);

}