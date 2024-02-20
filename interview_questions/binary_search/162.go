package binary_search


/***
哎， 这么简单的问题想半天， 有问题啊！
 */
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	if nums[n-1] > nums[n-2] {
		return n-1
	}
	l, r := -1, n-1
	for l+1 < r {
		mid := (l + r)>>1
		if nums[mid] < nums[mid+1] {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 == r
	return r
}

