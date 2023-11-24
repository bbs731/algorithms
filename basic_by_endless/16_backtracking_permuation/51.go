package backtracking

import "strings"

// what the fuck! 花了你半个小时的时间 // n =2 , n=3 时候是无解的。
func printAns(cols []int) (ans []string) {
	for i := 0; i < len(cols); i++ {
		pattern := make([]string, len(cols))
		for j := 0; j < len(cols); j++ {
			if cols[i] == j {
				pattern[j] = "Q"
			} else {
				pattern[j] = "."
			}
		}
		ans = append(ans, strings.Join(pattern, ""))
	}
	return
}

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	cols := make([]int, n)
	vcols := make([]bool, n)

	diag1 := make([]bool, 2*n) // i+j
	diag2 := make([]bool, 2*n) // i-j

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, printAns(cols))
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
