//
// Created by chunlei zhang on 29/08/2017.
//
#include <cstdio>


// in case we may have multiple v in the array A,
// lower bound is going to find the first occurrence of v, if not found return such index i if v is inserted, the A
// array is still in order.

// 需要知道我们在搜索的范围虽然是 [x, y) 闭开区间， 但是 bsearch 返回的结果却可以是 [x,y] 闭闭区间
// 试想一下，如果 v 比 A[y-1] 的值还要大，我们应该返回的位置应该是 y

int lower_bound(int *A, int x, int y, int v) {
    int m;
    while (x < y) {
        m = x + (y - x) / 2;
        if (A[m] >= v) {
            y = m;
        } else {
            x = x + 1;
        }
    }
    return x;
}

int lower_bound_implicit(int *A, int x, int y, int v) {
    int m;
    while (x < y) {  // [x, y) search range but return range is [x, y]
        m = x + (y - x) / 2;
        if (A[m] > v) {
            y = m;
        } else if (A[m] == v) {
            y = m;
        } else if (A[m] < v) {
            x = m + 1;
        }
    }

    return x;
}

int main() {
    int A[7] = {-1, 2, 5, -3, -9, 10, 4};
    int f = lower_bound_implicit(A, 0, 7, 13);
    printf(" return value is %d\n", f);

    int B[8] = {0, 0, 2, 2, 2, 4, 5};
    f = lower_bound(B, 0, 8, 2);
    printf(" return value is %d\n", f);

    f = lower_bound(B, 0, 8, 1);
    printf(" return value is %d\n", f);
}
