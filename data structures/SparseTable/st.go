package SparseTable

/*
时间复杂度，
构建表（预处理）：  O(nlgn)
查询： 			  O(1)
 */
const N = 2000001
const LogN = 21 /* log2000001 = 20.9  = 21  take the seiling */

func main() {

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, LogN+1)
	}
	logn := make([]int, N+1)

	logn[1] = 0
	logn[2] = 1
	for i := 3; i <= N; i++ {
		logn[i] = logn[i/2] + 1
	}

	// ST 表的构建  f[i][j] = max(f[i][j-1], f[i+ 1 <<(j-1]][j-1])
	for j := 1; j <= LogN; j++ {
		for i := 1; i+1<<uint(j-1) <= N; i++ {
			dp[i][j] = max(dp[i][j-1], dp[i+1<<uint(j-1)][j-1])
		}
	}

	var l, r int // query left and right point
	s := logn[r-l+1]
	ans := max(dp[l][s], dp[r-1<<uint(s)+1][s])
	_ := ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
