package dp

/*

		f[i+1][1] = max(f[i][1], f[i][0] - price[i])
		f[i+1][0] = max(f[i][0], f[i][1] + prices[i])
初始化： f[0][0] = 0 , f[0][1] = -inf


因为  f[i+1] 用到 f[i] 所以是正序哈!

降维：
		f[1] = max(f[1], f[0] -prices[i])
		f[0] = max(f[0], f[1] + prices[i])
初始化： f[0] = 0 f[1] = -inf

 */
func maxProfit(prices []int) int {
	inf := int(1e11)
	var f0, f1 int
	f1 = -inf

	for i := range prices {
		//old_f1 := f1  存还是不存 f1 好像都行， 因为下一轮更新 f0 的时候，需要用最新的 f1 更新。
		f1 = max(f1, f0-prices[i])
		//f0 = max(f0, old_f1+prices[i])
		f0 = max(f0, f1+prices[i])
	}
	return f0
}

/*
		dfs(i, 0) = max(dfs(i-1, 0), dfs(i-1, 1) + price[i-1])
		dfs(i, 1) = max(dfs(i-1, 1), dfs(i-1, 0) - price[i-1])
初始化：	dfs(0, 0) = 0 ,  dfs(0, 1) = -inf

翻译：
		f[i+1][1] = max(f[i][1], f[i][0] - prices[i])
		f[i+1][0] = max(f[i][0], f[i][1] + prices[i])
初始化： f[0][0] = 0 , f[0][1] = -inf

 */
func maxProfit_dp(prices []int) int {
	inf := int(1e10)
	n := len(prices)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
	}
	f[0][1] = -inf

	for i := range prices {
		f[i+1][1] = max(f[i][1], f[i][0]-prices[i])
		f[i+1][0] = max(f[i][0], f[i][1]+prices[i])
	}
	return f[n][0]
}

/*
dfs(i, 0)  第 i 天结束的时候， 不持有股票的最大利润
dfs(i, 1)  第 i 天结束的时候， 持有股票时的最大利润

转态转移

dfs(i, 0) = max(dfs(i-1, 0), dfs(i-1, 1) + price[i-1])
dfs(i, 1) = max(dfs(i-1, 1), dfs(i-1, 0) - price[i-1])

初始化条件  		dfs(-1, 0) = 0,  dfs(-1, 1) = -inf  允许 i = 0 那么 price[i-1] 会越界， 所以把  i + 1
初始化 变成      dfs(0, 0) = 0 ,  dfs(0, 1) = -inf

要求的答案是 dfs(n-1, 0)
变成  dfs(n, 0)
 */
func maxProfit_dfs(prices []int) int {
	inf := int(1e10)
	n := len(prices)
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i int, hold int) int {
		if i == 0 {
			if hold == 0 {
				return 0
			}
			return -inf
		}

		if cache[i][hold] != -1 {
			return cache[i][hold]
		}

		var ans int
		if hold == 1 {
			ans = max(dfs(i-1, 1), dfs(i-1, 0)-prices[i-1])
		} else {
			ans = max(dfs(i-1, 0), dfs(i-1, 1)+prices[i-1])
		}
		cache[i][hold] = ans
		return ans

	}
	return dfs(n, 0)
}
