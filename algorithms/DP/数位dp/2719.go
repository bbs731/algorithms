package 数位dp

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
