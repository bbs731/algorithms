package digit_dp

// 时间复杂度,  O(pos * 10)  状态数 x 状态转移数
// Digig DP template
func problem(n int) int {
	var a []int
	var dp []int

	var dfs func(int, bool, bool) int
	dfs = func(pos int, limit bool, lead bool) int {
		if pos == -1 {
			return 1 // 各位枚举， 根据题目的情况, 不一定返回 1
		}

		if !limit && !lead && dp[pos] != -1 {
			return dp[pos]
		}

		up := 9
		if limit {
			up = a[pos]
		}
		ans := 0

		for i := 0; i <= up; i++ {
			if true {

			} else {

			}
			ans += dfs(pos-1, limit && i == up, lead && i == 0)
		}

		if !limit && !lead {
			dp[pos] = ans
		}
		return ans
	}
	pos := 0

	for n != 0 {
		a[pos] = n % 10
		pos++
		n /= 10
	}
	return dfs(pos-1, true, true)
}
