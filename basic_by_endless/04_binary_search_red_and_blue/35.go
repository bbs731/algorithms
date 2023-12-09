package binary_search

import "sort"

func searchInsert(nums []int, target int) int {
	//return lower_bound(nums, target)
	return sort.SearchInts(nums, target)
}
