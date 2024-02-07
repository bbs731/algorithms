package bits_operation

// 用集合的解法。 （还有 dfs 的解法）
// 集合的解法， 套用了两个板子：  1. 枚举集合，  2. 遍历集合里面的元素
func subsets(nums []int) [][]int {

	n := len(nums)
	ans := make([][]int, 0)

	for s := 0; s < 1<<n; s++ {
		// 处理s 的逻辑, 集合 s 里面的元素范围是 0 到 n-1
		p := []int{}
		//遍历集合： 假设元素包括 从 0 到 n-1 个元素. 判断每个元素是否在集合 s 中
		for i := 0; i < n; i++ {
			if s>>i&1 == 1 { // i 在 s 中
				// 处理 i 的逻辑
				p = append(p, nums[i])
			}
		}
		ans = append(ans, p)
	}

	//ans = append(ans, []int{})
	return ans
}
