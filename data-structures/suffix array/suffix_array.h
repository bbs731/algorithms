//
// Created by chunlei zhang on 29/08/2017.
//

#ifndef ALGORITHMS_SUFFIX_ARRAY_H_H
#define ALGORITHMS_SUFFIX_ARRAY_H_H
#include <string>

const int MAX_N = 1001 * 1000 + 100;

void construct_lcp(std::string S, int *sa, int *lcp);
void construct_lcp_implicit(std::string S, int *sa, int *lcp);

void construct_sa(std::string S, int *sa);
void construct_rank(int *sa, int n);
#endif //ALGORITHMS_SUFFIX_ARRAY_H_H
