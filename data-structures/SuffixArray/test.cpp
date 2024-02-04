//
// Created by chunlei zhang on 29/08/2017.
//

#include <iostream>
#include "suffix_array.h"

using namespace std;
extern int ranks[MAX_N];

void print_sa(int *sa, int n) {

    cout << "suffix array: " << endl;
    for (int i = 0; i < n; i++) {
        cout << sa[i] << "\t";
    }
    cout << endl;

}

void print_ranks(int *ranks, int n) {
    cout << "ranks: " << endl;
    for (int i = 0; i < n; i++) {
        cout << ranks[i] << "\t";
    }
    cout << endl;

}

void print_height(int *lcp, int n) {
    cout << "lcp: " << endl;
    for (int i = 0; i < n; i++) {
        cout << lcp[i] << "\t";
    }
    cout << endl;
}

int main() {
    string S = "banana";
    int n = S.length() + 1;
    // to-do adding construct_sa() function instead of hardcoded suffix array
//    int sa[] = {6, 5, 3, 1, 0, 4, 2}; // "banana" 's suffix array
    int *sa = (int *) calloc(n, sizeof(int));
    int *lcp = (int *) calloc(n, sizeof(int));
    construct_sa(S, sa);
    construct_rank(sa, n);
    construct_lcp_implicit(S, sa, lcp);

    print_sa(sa, n);
    print_ranks(ranks, n);
    print_height(lcp, n);

    construct_lcp(S, sa, lcp);
    print_height(lcp, n);

}
