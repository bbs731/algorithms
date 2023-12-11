package sliding_window

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。



示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

 */

/*
 参考 15题， 把优化的小技巧 later 加上
掌握了解题技巧之后，就考验， 编码准确性和熟练性。 (在去重的处理上，还是犯了2次错误)
 */

func threeSum(nums []int, target int) [][]int {
	// nums is sorted
	n := len(nums)
	ans := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		x := nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			sum := x + nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				ans = append(ans, []int{x, nums[left], nums[right]})
				left++
				for left < n && nums[left] == nums[left-1] {
					left++
				}
				for right+1 < n && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := threeSum(nums[i+1:], target-nums[i])
		if len(t) > 0 {
			for _, l := range t {
				ans = append(ans, append(l, nums[i]))
			}
		}
	}
	return ans
}
