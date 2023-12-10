package sliding_window

import "sort"


func abs(a, b int) int {
	if a > b {
		return  a-b
	}
	return b -a
}
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	ans := int(1e9)

	for i := 0; i < n-2; i++ {
		x := nums[i]

		left := i + 1
		right := n - 1

		for left < right {
			sum := x + nums[left] + nums[right]
			if abs(sum, target)  < abs(ans, target) {
				ans = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
