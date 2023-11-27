package dp

import "math"

/*
让我们来降维: 根据上一种解法
f[i+1][c] = min (f[i][c], f[i+1][c- coin[i]] + 1)

降维：去掉第一个维度， 并且 loop c 的时候，是正序的loop
f[c] = min(f[c] + f[c - coin[i]] + 1
*/
// 这降维写起来， 太TMD爽了
func coinChange(coins []int, amount int) int {
	n := len(coins)
	inf := math.MaxInt / 2 //这样初始化 inf 也可以。

	f := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		f[j] = inf
	}
	f[0] = 0

	for i := 0; i < n; i++ {
		//for j := coins[i]; j <= amount; j++ {  // 这样写是错误的， 因为需要 f[i+1][j] = f[i][j]  用 f[i][j] 来更新 f[i+1][j]
		for j := coins[i]; j <= amount; j++ {
			//if j >= coins[i] { //这里还是挺难的， 第一次写错。好好考虑一下
			f[j] = min(f[j], f[j-coins[i]]+1)
			//}
		}
	}
	if n == 0 {
		return 0
	}
	ans := f[amount]
	if ans == inf {
		return -1
	}
	return ans
}

/*
dfs(i, c) = min(dfs(i-1, c), dfs(i, c-coins[i])+1)
翻译成 dp 递推的数组
f[i][c] = min ( f[i-1][c] , f[i][c-coin[i]] + 1)
i changed to i+1
f[i+1][c] = min (f[i][c], f[i+1][c- coin[i]] + 1)
 */

func coinChange_dp(coins []int, amount int) int {
	n := len(coins)
	inf := int(1e9)

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			f[i][j] = inf
		}
	}
	f[0][0] = 0

	for i := 0; i < n; i++ {
		//for j := coins[i]; j <= amount; j++ {  // 这样写是错误的， 因为需要 f[i+1][j] = f[i][j]  用 f[i][j] 来更新 f[i+1][j]
		for j := 0; j <= amount; j++ {
			if j >= coins[i] { //这里还是挺难的， 第一次写错。好好考虑一下
				f[i+1][j] = min(f[i][j], f[i+1][j-coins[i]]+1)
			} else {
				f[i+1][j] = f[i][j]
			}
		}
	}
	if n == 0 {
		return 0
	}
	ans := f[n][amount]
	if ans == inf {
		return -1
	}
	return ans
}

func coinChange_dfs_cache(coins []int, amount int) int {
	n := len(coins)
	var dfs func(int, int) int
	inf := int(1e9)

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
				return 0
			}
			return inf
		}
		if c < 0 {
			return inf
		}

		if cache[i][c] != -1 {
			return cache[i][c]
		}

		// 选择 coin i, 不选择 coin i
		ans := min(dfs(i-1, c), dfs(i, c-coins[i])+1)
		cache[i][c] = ans
		return ans
	}

	if n == 0 {
		return 0
	}
	ans := dfs(n-1, amount)
	if ans == inf {
		return -1
	}
	return ans
}

// 这道题是个完全knapsack 的问题。
// 因为是选择最小的方案， 所以，用 min，还是恰好为 amount的类型。
// dfs(i,c) = min(dfs(i-1, c), dfs(i, c-coins[i]) + 1)
func coinChange_basic_dfs(coins []int, amount int) int {
	n := len(coins)
	var dfs func(int, int) int
	inf := int(1e9)

	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return inf
		}
		if c < 0 {
			return inf
		}

		// 选择 coin i, 不选择 coin i
		ans := min(dfs(i-1, c), dfs(i, c-coins[i])+1)
		return ans
	}

	if n == 0 {
		return 0
	}
	ans := dfs(n-1, amount)
	if ans == inf {
		return -1
	}
	return ans
}
