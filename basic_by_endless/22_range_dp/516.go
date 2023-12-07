package dp

import "fmt"

/*
 range dp
 dfs(i, j) = max{ dfs(i+1, j-1) if s[i] == s[j],  dfs(i+1, j) , dfs(i, j-1))
边界条件： dfs(i+1, i) = 0 , dfs(i, i) = 1
 	f[i][j+1] = max( f[i+1][j]+2 , f[i+1][j+1], f[i][j])    // 因为 f[i] 需要依赖 f[i+1] 所以， i 需要倒序排列。
初始化条件  f[i][i] = 1

另外这一道题，比较特殊的是, i 和 j 的维度是不一样的一个是 n 维度， 一个是 n+1 维度。
 */
func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n+1)
		f[i][i+1] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j+1] = f[i+1][j] + 2
			} else {
				f[i][j+1] = max(f[i+1][j+1], f[i][j])
			}
		}
	}
	return f[0][n]
}

// 需要加 cache 记忆化搜索
func longestPalindromeSubseq_dfs(s string) int {
	n := len(s)

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		ans := 0
		if s[i] == s[j] {
			ans = 2 + dfs(i+1, j-1)
		} else {
			ans = max(dfs(i+1, j), dfs(i, j-1))
		}
		return ans
	}
	return dfs(0, n-1)
}

func longestCommonSubsequence_dp_basic(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

func longestPalindromeSubseq_lcs(s string) int {
	rs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		rs[len(s)-1-i] = s[i]
	}
	fmt.Println(string(rs))
	return longestCommonSubsequence_dp_basic(s, string(rs))
}
