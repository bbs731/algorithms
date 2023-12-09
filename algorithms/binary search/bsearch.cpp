//
// Created by chunlei zhang on 29/08/2017.
//
#include <cstdio>


int bsearch(int *A, int x, int y, int v){

    int m;
    while (x < y) {   // search [x, y)
        m = x + (y - x) / 2;
        if (A[m] == v) {
            return m;
        }else if (A[m] > v) {
            y = m;
        }else{
            x = m + 1;
        }
    }
    return -1; // not found
}

