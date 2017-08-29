//
// Created by chunlei zhang on 29/08/2017.
//

#include <string>
#include <cstdlib>
#include "suffix_array.h"

using namespace std;

int ranks[MAX_N + 1];


// 怎么能看出来 construct_lcp 的 time complexity is O(n）？
// h[i] = lcp[rank[i]] we have property that:
// h[i] >= h[i-1] - 1  so we can calculate by sequence of h[1], h[2], h[3], .....
// see the construct_lcp_implicit() implementation
void construct_lcp(string S, int *sa, int *lcp) {
    int n = S.length();
    for (int i = 0; i <= n; i++) ranks[sa[i]] = i;
    int h = 0;  // see we define h as h[i] = lcp[rank[i]] here
    // we can define h as array but when we loop i from 0 to n , only one variable will be sufficient
    lcp[0] = -1;
    for (int i = 0; i < n; i++) {
        int j = sa[ranks[i] - 1];

        if (h) h--;
        for (; j + h < n && i + h < n; h++) {
            if (S[j + h] != S[i + h]) break;
        }
        lcp[ranks[i]] = h;
    }
}


void construct_lcp_implicit(string S, int *sa, int *lcp) {
    int n = S.length();
    // we can acutally replace h[] array with single h variable as did in construct_lcp()
    int *h = (int *) calloc(n+1, sizeof(int));
    for (int i = 0; i <= n; i++)
        ranks[sa[i]] = i;

    lcp[0] = -1;
    for (int i = 0; i < n; i++) {
        int j = sa[ranks[i] - 1];

        if (h[i - 1])
            h[i] = h[i - 1] - 1;
        for (; j + h[i] < n && i + h[i] < n; h[i]++) {
            if (S[j + h[i]] != S[i + h[i]])
                break;
        }
        lcp[ranks[i]] = h[i];
    }
}