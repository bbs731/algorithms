package main

import "fmt"

/*
KMP 的实现：

可以参考：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/strings.go#L94-L109
*/

func prefix_function(pattern string) []int {
	n := len(pattern)
	pi := make([]int, n)

	j := 0 // j 记录的是 pi[i-1], 初始化为 pi[0]  即为 0
	for i := 1; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func kmp(text, pattern string) []int {

	pi := prefix_function(pattern)
	pos := make([]int, 0)

	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = pi[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			pos = append(pos, i-len(pattern)+1)
			j = pi[j-1]
		}
	}
	return pos
}

func main() {

	fmt.Println(prefix_function("ababaca"))
	fmt.Println(kmp("bacbababacabcbab", "ababaca"))

}
