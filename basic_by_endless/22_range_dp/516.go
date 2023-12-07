package dp

import "fmt"

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

func longestPalindromeSubseq(s string) int {
	rs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		rs[len(s)-1-i] = s[i]
	}
	fmt.Println(string(rs))
	return longestCommonSubsequence_dp_basic(s, string(rs))
}
