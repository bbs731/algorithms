package dp

import "math"

/*
	dfs(i, c) = max(dfs(i-1, c - nums[i]) + 1, dfs(i-1, c))
	terminate condition  i<0  and c == 0  then return 0 else return -inf

	f[i+1][c] = max( f[i][c-nums[i]] + 1, f[i][c] )

	f[c] = max(f[c-nums[i]] +1, f[c])  // 应该倒序的 loop c
	initialized f[0] = 0 others -inf

 */
func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	inf := math.MaxInt / 2

	f := make([]int, target+1)
	for i := 0; i <= target; i++ {
		f[i] = -inf
	}
	f[0] = 0

	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			f[j] = max(f[j], f[j-nums[i]]+1)
		}
	}

	//if f[target] == -inf {  // 这样写错误， 为什么呢？
	// test case : [1,1,5,4,5] 3
	// 知道为什么了， 因为在算 f[3] 的时候， 用到了 f[3] = f[2] + 1，  这个时候， f[2] = -inf， f[3]更新为 = -inf +1
	// 所以最后判断 f[3] 的时候， 应该用 <=0 来判断
	if f[target] <= 0 {
		return -1
	}
	return f[target]
}
