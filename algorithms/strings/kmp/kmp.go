package main

import "fmt"

/*
KMP 的原理：  https://www.zhihu.com/question/21923021/answer/37475572
帮助记忆的两个点：
1. 当 text 和 pattern 发生不匹配的时候，  text 不回退， 只有 Pattern 回退。
2. 既然需要 Pattern 回退，那么 Pattern 需要自己和自己 match (Pattern 自己的 真前缀和真后缀 的 match)
   得到回退的 pos

3. 分析时间复杂度的技巧，如果证明，得到 Pattern 回退数组的时间复杂度是线性了 O（m)
	Text 和 Pattern match 的时间复杂度 是 O（n) 的。

KMP 的实现：



可以参考：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/strings.go#L94-L109
*/

func calcMaxMatchLengths(pattern string) []int {
	n := len(pattern)
	match := make([]int, n)

	for i, c := 1, 0; i < n; i++ {
		v := pattern[i]
		for c > 0 && pattern[c] != v {
			c = match[c-1]
		}
		if pattern[c] == v {
			c++
		}
		match[i] = c
	}
	return match
}

// return all the match positions
func kmpSearch(text, pattern string) []int {
	match := calcMaxMatchLengths(pattern)
	pos := make([]int, 0)
	lenP := len(pattern)
	c := 0

	for i, v := range text {
		for c > 0 && pattern[c] != byte(v) {
			c = match[c-1]
		}
		if pattern[c] == byte(v) {
			c++
		}
		if c == lenP {
			pos = append(pos, i-lenP+1)
			c = match[c-1]
		}
	}
	return pos
}

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
