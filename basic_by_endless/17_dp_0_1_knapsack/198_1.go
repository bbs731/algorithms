package dp


// f[i+2] = max（f[i+1], f[i] + nums[i])
// 我们可以把空间优化成 O（1）， 具体就需要 f0, 和 f1 两个变量
func rob(nums []int) int {
	n := len(nums)
	//f0, f1 :=nums[0], nums[1] 这个初始化也是错误的， 这里的 f0 相当于 f[-2], f1 相当于 f[-1] 应该初始化为0
	var f0, f1 int
	for i:=0; i<n; i++ {
		new_f := max(f1, f0 + nums[i])
		f0 = f1
		f1= new_f
	}
	return f1
}

// dfs(i) = max(dfs(i-1) , dfs(i-2) + nums[i])
// 一比一， 把 递归改成 数组 递推。
// f[i] = max(f[i-1], f[i-2] + nums[i])
// f[i+2] = max（f[i+1], f[i] + nums[i])
func rob2(nums []int) int {
	n := len(nums)

	f := make([]int, n+2)
	// 这里不需要初始话，因为我们把 i 都变成了 i+2， 因此， f[0] 相当于 f[-2] 为 0 正常， f[2] 才相当于 nums[0]
	//f[0] = nums[0] 这个初始化是错误的。

	for i:=0; i<n; i++ {
		f[i+2] = max(f[i+1], f[i] +nums[i])
	}
	return f[n+1]
}

// 子集问题回溯的递归解法：
// dfs(i) = max(dfs(i-1) , dfs(i-2) + nums[i])
func rob1(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i :=0; i <n; i++ {
		// input nums[] 可能全是 0, 测试用例的 67
		cache[i]= -1
	}

	var dfs func(int) int

	dfs = func(i int) int{
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		cache[i] = max(dfs(i-1), dfs(i-2) + nums[i])
		return cache[i]
	}
	return dfs(n-1)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}