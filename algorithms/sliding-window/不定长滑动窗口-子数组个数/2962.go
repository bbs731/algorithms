package sliding_window

import "slices"

// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/
// 灵神的题解，就是这个套路， 就是这个味道！
func countSubarrays(nums []int, k int) int64 {
	// slices.Max 这个是新学到的知识， 开心
	maxn := slices.Max(nums)

	cnts := 0
	left :=0
	ans := 0

	for _, v := range nums {
		if v == maxn {
			cnts++
		}
		for cnts >=k {
			// remove left element from window
			if nums[left] == maxn {
				cnts--
			}
			left++
		}
		ans += left
	}
	return int64(ans)
}
