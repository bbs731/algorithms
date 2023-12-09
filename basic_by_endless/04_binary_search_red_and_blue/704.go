package binary_search

import "sort"

func search(nums []int, target int) int {
	pos := sort.SearchInts(nums, target)
	if pos == len(nums) || nums[pos] != target {
		return -1
	}
	return pos
}
