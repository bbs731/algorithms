package dp

func rob_knapsack(nums []int) int {
	n := len(nums)
	//f0, f1 :=nums[0], nums[1] 这个初始化也是错误的， 这里的 f0 相当于 f[-2], f1 相当于 f[-1] 应该初始化为0
	var f0, f1 int
	for i := 0; i < n; i++ {
		new_f := max(f1, f0+nums[i])
		f0 = f1
		f1 = new_f
	}
	return f1
}

// 恩， 这里的 n == 1, n ==2 的判断，有更简洁的写法， 看看灵神的答案
//	https://leetcode.cn/problems/house-robber-ii/solutions/2445622/jian-ji-xie-fa-zhi-jie-diao-yong-198-ti-qhvri/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 这里属于枚举了吧。 （子集问题, 思路， 选 nums[0] 不选 nums[0])
	// 选 nums[0]
	ans := rob_knapsack(nums[2:n-1]) + nums[0]

	//不选 nums[0]
	return max(ans, rob_knapsack(nums[1:n]))
}
