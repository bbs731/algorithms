package dp

import "sort"

// 这都能错吗？ 和 LC300 LIS 用到的方法类似！
// 复杂度是 n*logn
func maxProfit(prices []int) int {
	q := make([]int, 0) // 存值
	ans := 0

	for _, x := range prices {
		pos := sort.SearchInts(q, x)
		if pos == len(q) {
			q = append(q, x)
		} else {
			q[pos] = x
		}
		ans = max(ans, x-q[0])
	}
	return ans
}

/*
	f[i] = f[i] - min(f[j])  for j < i
dp 数组都不用创建。 时间复杂度
n^2 这道题会超时
how to optimize?
 */
func maxProfit_dp(prices []int) int {
	//n := len(prices)
	//f := make([]int, n)
	ans := 0

	for i, x := range prices {
		for j := 0; j < i; j++ {
			if x > prices[j] {
				ans = max(ans, x-prices[j])
			}
		}
	}
	return ans
}
