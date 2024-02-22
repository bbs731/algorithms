package dp

import "math"

/***

给你两个字符串 str1 和 str2，返回同时以 str1 和 str2 作为 子序列 的最短字符串。如果答案不止一个，则可以返回满足条件的 任意一个 答案。

如果从字符串 t 中删除一些字符（也可能不删除），可以得到字符串 s ，那么 s 就是 t 的一个子序列。



示例 1：

输入：str1 = "abac", str2 = "cab"
输出："cabac"
解释：
str1 = "abac" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 的第一个 "c"得到 "abac"。
str2 = "cab" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 末尾的 "ac" 得到 "cab"。
最终我们给出的答案是满足上述属性的最短字符串。
示例 2：

输入：str1 = "aaaaaaaa", str2 = "aaaaaaaa"
输出："aaaaaaaa"

 */

/***
2024-02-22 怎么，感觉你的功力，还在呢， 呵呵。

very good job!  only make a mistake on initialization


f[i][j] = min (f[i-1][j], f[i][j-1]) + 1  a[i] != b[j]
		= f[i-1][j-1] + 1 if a[i] == b[j]

 */

func shortestCommonSupersequence(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)
	inf := math.MaxInt32
	if n == 0 {
		return str2
	}
	if m == 0 {
		return str1
	}

	f := make([][]int, n+1)

	// 这种题目， 应该是初始化是最难的吧！
	for i := range f {
		f[i] = make([]int, m+1)
		for j := range f[i] {
			f[i][j] = inf
			if i == 0 {
				f[0][j] = j
			}
			if j == 0 {
				f[i][0] = i
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}

	// build the answer
	l := f[n][m]
	ans := make([]byte, l)
	p1, p2 := n-1, m-1
	for k := l - 1; k >= 0; k-- {
		if p1 < 0 {
			ans[k] = str2[p2]
			p2--
			continue
		}
		if p2 < 0 {
			ans[k] = str1[p1]
			p1--
			continue
		}

		if str1[p1] == str2[p2] {
			ans[k] = str1[p1]
			p1--
			p2--
			continue
		}

		if f[p1+1][p2+1] == f[p1+1][p2]+1 {
			// select from str2
			ans[k] = str2[p2]
			p2--

		} else {
			// select from str1
			ans[k] = str1[p1]
			p1--
		}
	}
	return string(ans)
}
