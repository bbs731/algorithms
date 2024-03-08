package one_day_exercise

/***
超时，todo 去看题解，的分组讨论把。
 */

func maxValueAfterReverse(nums []int) int {
	d := make([]int, 0)
	sum := 0
	for i := 1; i <= len(nums)-1; i++ {
		d = append(d, nums[i]-nums[i-1])
		sum += nums[i] - nums[i-1]
	}

	add := 0
	n := len(nums)
	ans := sum
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i > 0 {
				add += abs(nums[i-1], nums[j]) - abs(nums[i-1], nums[i])
			}
			if j < n-1 {
				add += abs(nums[j+1], nums[i]) - abs(nums[j+1], nums[j])
			}
			ans = max(ans, sum+add)
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
