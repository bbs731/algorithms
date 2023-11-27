package dp

import "math"

func numSquares_dfs(target int, nums []int) int {
	n := len(nums)
	var dfs func(int, int) int
	inf := math.MaxInt / 2

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			cache[i][j] = -1
		}
	}

	dfs = func(i int, c int) int {
		if i < 0 { // 这里的条件
			if c == 0 {
				return 0 //这里的返回值， 0， 还是 1?  容易出错的地方
			}
			return inf
		}

		if c < 0 {
			return inf
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}

		ans := min(dfs(i, c-nums[i])+1, dfs(i-1, c)) // 忘了+1， 容易出错的地方
		cache[i][c] = ans
		return ans
	}

	return dfs(n-1, target)
}

/*
	dfs(i, c) = min( dfs(i, c-nums[i]) +1 , dfs(i-1, c))

	f[i][c] = min(f[i][c-nums[i]] + 1, f[i-1][c])

	f[i+1][c] = min(f[i+1][c-nums[i]] + 1 , f[i][c])

	f[c] = min (f[c-nums[i]]+1, f[c]) // 根据上面的公式， loop c 的顺序是： 正序
 */
func numSquares_dp(target int, nums []int) int {
	n := len(nums)
	dp := make([]int, target+1)
	inf := math.MaxInt / 2

	for i := 0; i <= target; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := nums[i]; j <= target; j++ {
			dp[j] = min(dp[j], dp[j-nums[i]]+1)
		}
	}
	return dp[target]
}

func numSquares(n int) int {
	nums := make([]int, 0)

	square := 1
	for i := 1; square <= n; {
		nums = append(nums, square)
		i++
		square = i * i
	}
	return numSquares_dp(n, nums)
}
