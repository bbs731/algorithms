package one_day_exercise

/*

给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x 满足以下条件，我们称它是一个好整数：

num1 <= x <= num2
min_sum <= digit_sum(x) <= max_sum.
请你返回好整数的数目。答案可能很大，请返回答案对 109 + 7 取余后的结果。

注意，digit_sum(x) 表示 x 各位数字之和。



示例 1：

输入：num1 = "1", num2 = "12", min_num = 1, max_num = 8
输出：11
解释：总共有 11 个整数的数位和在 1 到 8 之间，分别是 1,2,3,4,5,6,7,8,10,11 和 12 。所以我们返回 11 。
示例 2：

输入：num1 = "1", num2 = "5", min_num = 1, max_num = 5
输出：5
解释：数位和在 1 到 5 之间的 5 个整数分别为 1,2,3,4 和 5 。所以我们返回 5 。


提示：

1 <= num1 <= num2 <= 10^22
1 <= min_sum <= max_sum <= 400
 */

/***

想一想，时间复杂度，是多少？
动态规划的时间复杂度 === 状态个数 × 单个状态的计算时间

*/
const mod int = 1e9 + 7

func solve(n string, min_sum, max_sum int) int {
	m := len(n)

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, 9*m+1)
		for j := 0; j < 9*m+1; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool) int
	dfs = func(i int, sum int, isLimit bool) int {
		if i == m {
			if sum >= min_sum && sum <= max_sum {
				return 1
			}
			return 0
		}

		if !isLimit {
			if memo[i][sum] != -1 {
				return memo[i][sum]
			}
		}

		up := 9
		if isLimit {
			up = int(n[i] - '0')
		}

		res := 0
		for d := 0; d <= up; d++ {
			res += dfs(i+1, sum+d, isLimit && d == up)
			res %= mod
		}

		if !isLimit {
			memo[i][sum] = res
		}
		return res
	}

	return dfs(0, 0, true)
}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	// 22 位 10进制数字，已经超过了 int64 的范围。
	sum := 0
	complement := 0
	for _, d := range num1 {
		sum += int(d - '0')
	}
	if sum >= min_sum && sum <= max_sum {
		complement = 1
	}

	return (solve(num2, min_sum, max_sum) - solve(num1, min_sum, max_sum) + complement + mod) % mod
}
