package dp

func maxProfit_dp(k int, prices []int) int {
	n := len(prices)
	inf := int(1e10)
	f := make([][][]int, n+1)

	// 这里的， 初始化才是最难得!
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			f[i][j] = make([]int, 2)
			f[i][j][1] = -inf
		}
	}
	for j := 0; j < k+1; j++ {
		f[0][k][0] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			f[i+1][j+1][1] = max(f[i][j+1][1], f[i][j][0]-prices[i])
			f[i+1][j+1][0] = max(f[i][j+1][0], f[i][j+1][1]+prices[i])
		}
	}

	return f[n][k][0]
}

func maxProfit(prices []int) int {
	return maxProfit_dp(2, prices)
}
