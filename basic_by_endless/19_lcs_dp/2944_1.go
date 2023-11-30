package dp

/*
 题解
https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/solutions/2542044/dpcong-on2-dao-onpythonjavacgo-by-endles-nux5/
 */

/*
用单调队列，来优化时间复杂度， 从 O(n^2) 到 O（n)
 */
func minimumCoins(prices []int) int {
	n := len(prices)
	type pair struct{ i, value int }
	q := []pair{{n + 1, 0}}

	for i := n; i >= 1; i-- {
		for q[0].i > 2*i+1 {
			q = q[1:]
		}

		f := prices[i-1] + q[0].value
		for f <= q[len(q)-1].value {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f})
	}
	return q[len(q)-1].value
}

/*
	f[i] = price[i-1] + min(f[j]) j from i+1 to 2i+1
	f[i >n] = 0 // 初始化条件
 */
func minimumCoins_dp(prices []int) int {
	n := len(prices)
	inf := int(1e9)

	f := make([]int, 2*n+2)

	for i := n; i >= 1; i-- {
		ans := inf
		for j := i + 1; j <= 2*i+1; j++ {
			ans = min(ans, f[j])
		}
		f[i] = ans + prices[i-1]
	}
	return f[1]
}

/*
灵神的版本
 */
func minimumCoins_dfs2(prices []int) int {
	n := len(prices)
	inf := int(1e9)

	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i > n {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		ans := inf
		for j := i + 1; j <= i+i+1; j++ {
			ans = min(ans, dfs(j))
		}
		ans += prices[i-1]
		cache[i] = ans
		return ans
	}
	return dfs(1)
}

/*
我的版本：
 */
func minimumCoins_dfs1(prices []int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)
	inf := int(1e9)
	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i > n {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		ans := inf
		//for j := i + 1; j <= min(i+i+1, n-1); j++ {
		for j := i + 1; j <= i+i+1; j++ {
			ans = min(ans, prices[i]+dfs(j))
		}
		cache[i] = ans
		return ans
	}
	return dfs(1)
}
