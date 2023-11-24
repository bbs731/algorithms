package backtracking

import "strings"

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	cols := make([]int, n)
	vcols := make([]bool, n)

	diag1 := make([]bool, 2*n) // i+j
	diag2 := make([]bool, 2*n) // i-j

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 打印 ans 的这段代码，真牛！
			pattern := make([]string, n)
			for k, c := range cols {
				pattern[k] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, pattern)
			return
		}

		for j := 0; j < n; j++ {
			if vcols[j] == false && diag1[i+j] == false && diag2[i-j+n] == false {
				vcols[j] = true
				diag1[i+j] = true
				diag2[i-j+n] = true
				cols[i] = j
				dfs(i + 1)
				vcols[j] = false
				diag1[i+j] = false
				diag2[i-j+n] = false
			}
		}
	}
	dfs(0)

	return ans
}
