package dp

/*
	dfs(i, c) = dfs(i, c-coins[i]) + dfs(i-1, c)
改成递推
	f[i][c] = f[i][c-coins[i]] + dfs[i-1][c]
=>
	f[i+1][c] = f[i+1][c-coins[i]] + f[i][c]
降维
	f[c] = f[c-coins[i]] + f[c]  // 正序 loop C 就行。
*/

// 牛啊！ 一遍写对了。
func change(amount int, coins []int) int {
	n := len(coins)
	f := make([]int, amount+1)

	f[0] = 1

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] += f[j-coins[i]]
		}
	}
	return f[amount]
}

func change_dp(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				f[i+1][j] = f[i+1][j-coins[i]] + f[i][j]
			} else {
				f[i+1][j] = f[i][j]
			}
		}
	}
	return f[n][amount]
}

func change_cache(amount int, coins []int) int {
	n := len(coins)
	var dfs func(int, int) int

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			cache[i][j] = -1
		}
	}

	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}

		if c < 0 {
			return 0
		}

		if cache[i][c] != -1 {
			return cache[i][c]
		}

		// 选和不选 coins[i]
		ans := dfs(i, c-coins[i]) + dfs(i-1, c)
		cache[i][c] = ans
		return ans
	}
	return dfs(n-1, amount)
}

func change_dfs(amount int, coins []int) int {
	n := len(coins)

	var dfs func(int, int) int

	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}

		if c < 0 {
			return 0
		}

		// 选和不选 coins[i]
		ans := dfs(i, c-coins[i]) + dfs(i-1, c)
		return ans
	}
	return dfs(n-1, amount)
}
