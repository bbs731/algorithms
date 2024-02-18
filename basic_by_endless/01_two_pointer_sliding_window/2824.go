package sliding_window

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	cnt := 0

	left := 0
	right := n - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			cnt += right - left // 这里犯了个错误，不是 cnt++, 这里的计数是个技巧！
			left++
		} else {
			right--
		}
	}
	return cnt
}
