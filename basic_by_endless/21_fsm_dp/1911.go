package dp

/*
	f[0] = max(f[1] - nums[i], f[0])
	f[1] = max(f[0] + nums[i], f[1])
f[0] = f[1] = 0
return: f[1]

算法的力量吗？ 最后的代码这么的简洁吗？
 */
func maxAlternatingSum(nums []int) int64 {
	var f0, f1 int
	for _, x := range nums {
		new_f0 := max(f1-x, f0)
		f1 = max(f0+x, f1)
		f0 = new_f0
	}
	return int64(f1)
}

/*
偶数    dfs(i, 0) =   max( dfs(i-1, 1) - nums[i] , dfs(i-1, 0) )   // 选或者不选
奇数    dfs(i, 1) =   max( dfs(i-1, 0) + nums[i], dfs(i-1, 1) )

边界条件： dfs(-1, 0) = dfs(-1, 1) = 0

	f[i+1][0] = max(f[i][1] - nums[i], f[i][0])
	f[i+1][1] = max(f[i][0] + nums[i], f[i][1])

f[0][0] = f[0][1] = 0

return: f[n][1]
 */
func maxAlternatingSum_dp(nums []int) int64 {
	n := len(nums)
	f := make([][2]int, n+1)

	for i, x := range nums {
		f[i+1][0] = max(f[i][1]-x, f[i][0])
		f[i+1][1] = max(f[i][0]+x, f[i][1])
	}
	return int64(f[n][1])
}

/*
nums = [6,2,1,2,4,5]

定义   dfs(i, 0) 为， 到 Nums[i] 位置， 组成偶数个 subseq 的最大值。
       dfs(i, 1) 为， 到 nums[i] 为止， 组成的奇数个 subseq 的最大值。
偶数    dfs(i, 0) =   max( dfs(i-1, 1) - nums[i] , dfs(i-1, 0) )   // 选或者不选
奇数    dfs(i, 1) =   max( dfs(i-1, 0) + nums[i], dfs(i-1, 1) )

边界条件： dfs(-1, 0) = dfs(-1, 1) = 0

return : dfs(n-1, 1)
 */
func maxAlternatingSum_dfs(nums []int) int64 {
	n := len(nums)
	cache := make([][2]int64, n)
	for i := range cache {
		cache[i][0] = -1
		cache[i][1] = -1
	}
	var dfs func(int, int) int64
	dfs = func(i int, odd int) int64 {
		if i < 0 {
			return 0
		}

		if cache[i][odd] != -1 {
			return cache[i][odd]
		}

		var ans int64
		if odd == 1 {
			ans = max(dfs(i-1, 0)+int64(nums[i]), dfs(i-1, 1))
		} else {
			ans = max(dfs(i-1, 1)-int64(nums[i]), dfs(i-1, 0))
		}
		cache[i][odd] = ans
		return ans
	}

	return dfs(n-1, 1)
}
