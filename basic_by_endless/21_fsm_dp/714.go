package dp

func maxProfit_dp(prices []int, fee int) int {
	inf := int(1e10)
	n := len(prices)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
	}
	f[0][1] = -inf

	for i := range prices {
		f[i+1][1] = max(f[i][1], f[i][0]-prices[i]-fee)
		f[i+1][0] = max(f[i][0], f[i][1]+prices[i])
	}
	return f[n][0]
}

func maxProfit(prices []int, fee int) int {

	return maxProfit_dp(prices, fee)
}
