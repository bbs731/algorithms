package dp

// 我靠， 从findTargetSumWays_4做降维的过程真是太爽了！ 有机会再试试， 第一遍就做对了
// 感受到编程的快乐了吗？有机会的话，讲给菠萝听
func findTargetSumWays(nums []int, target int) int {
	// f[i+1][c] = f[i][c] + f[i][c-nums[i]]
	// 让我们来做空间的降维， 变成单维度的数组

	n := len(nums)
	//var s int  可以重用 target
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	f[0] = 1

	for i := 0; i < n; i++ {
		//for j := 0; j <= target; j++ {  // 这里要改成倒序的， 看状态转移方程就知道是应该正序还是倒序（灵神给过方法）
		for j := target; j >= nums[i]; j-- {
			//if j >= nums[i] { // 这里的条件好难啊， 看一看，递归时候的逻辑
			f[j] = f[j] + f[j-nums[i]]
			//}
		}
	}
	return f[target]
}

func findTargetSumWays_4(nums []int, target int) int {
	// 把记忆化搜索backtracking 转换成递推
	// dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-nums[i])
	// dfs -> 数组
	// f[i][c] = f[i-1][c] + f[i-1][c-nums[i]]
	// f[i+1][c] = f[i][c] + f[i][c-nums[i]]

	n := len(nums)
	//var s int  可以重用 target
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, target+1)
	}

	f[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] { // 这里的条件好难啊， 看一看，递归时候的逻辑
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = f[i][j] + f[i][j-nums[i]]
			}
		}
	}
	return f[n][target]
}

func findTargetSumWays_3(nums []int, target int) int {

	// p stands for all positive nums's sum
	// s stands for total sum
	// p - (s-p)  = target  ==>  p = (target + s) / 2
	// 这是对 target 的特殊处理，以便方便写 dfs (不过下面的 zhangchunlei +1000 的做法也不复杂， 就是 +1000 不好看)
	// 题意就变成了， 从 nums[] 中, 选出一些数，让其和 = (target + sum(nums))/2

	n := len(nums)
	//var s int  可以重用 target
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	// i 可以从 0 到 n, 也可以从 n-1 到 -1
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}

		// 这道题，其实还是挺难的。
		if c < nums[i] {
			return dfs(i-1, c)
		}

		cache[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
		return cache[i][c]
	}
	return dfs(n-1, target)
}

func findTargetSumWays_2(nums []int, target int) int {

	// p stands for all positive nums's sum
	// s stands for total sum
	// p - (s-p)  = target  ==>  p = (target + s) / 2
	// 这是对 target 的特殊处理，以便方便写 dfs (不过下面的 zhangchunlei +1000 的做法也不复杂， 就是 +1000 不好看)

	n := len(nums)
	//var s int  可以重用 target
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	// i 可以从 0 到 n, 也可以从 n-1 到 -1
	dfs = func(i, c int) int {
		if i == n {
			if c == 0 {
				return 1
			}
			return 0
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}

		// 这道题，其实还是挺难的。
		if c < nums[i] {
			return dfs(i+1, c)
		}

		cache[i][c] = dfs(i+1, c) + dfs(i+1, c-nums[i])
		return cache[i][c]
	}
	return dfs(0, target)
}

// zhangchunlei 的想法。
func findTargetSumWays_1(nums []int, target int) int {
	n := len(nums)

	target += 1000

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= 2000; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i == n {
			if c == target {
				return 1
			}
			return 0
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}

		//ans := 0
		//// i 选择  +
		//ans = dfs(i+1, c-1)
		//
		////i 选择 -
		//ans += dfs(i+1, c)
		cache[i][c] = dfs(i+1, c+nums[i]) + dfs(i+1, c-nums[i])
		return cache[i][c]

	}

	return dfs(0, 1000)
}
