package interview_questions

/***
这道题，再次证明了， 写二分查找的问题，
(l, r) 和 [l, r] 等区间的选择，对于，结果没有区别。
 */

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	if nums[n-1] == target {
		return n - 1
	}

	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[n-1] {
			// 这里面的判断，太复杂了， 有好几种写法。
			if target < nums[mid] && target > nums[n-1] {
				r = mid - 1
			} else {
				// 这个情况是复杂的， target > nums[mid] 和   target < nums[n-1] < nums[mid] 的情况下， 都在 mid 左边。
				l = mid + 1
			}
		} else {
			// nums[mid] <  nums[n-1]
			if target < nums[mid] || target > nums[n-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[n-1] {

			// 这里面的判断，太复杂了， 有好几种写法。
			//if target < nums[mid] && target > nums[n-1] {
			//	r = mid - 1
			//} else {
			//	// 这个情况是复杂的， target > nums[mid] 和   target < nums[n-1] < nums[mid] 的情况下， 都在 mid 左边。
			//	l = mid + 1
			//}

			// 这里面的判断，太复杂了， 有好几种写法。
			//if target < nums[mid] && target > nums[n-1] {
			if target > nums[mid] || target <= nums[n-1] { // 这个是 <= 这个是不是太难了？
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else {
			// nums[mid] <  nums[n-1]
			if target < nums[mid] || target > nums[n-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := -1, n

	for l+1 < r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[n-1] {
			if target < nums[mid] && target > nums[n-1] {
				r = mid
			} else {
				// 这个情况是复杂的， target > nums[mid] 和   target < nums[n-1] < nums[mid] 的情况下， 都在 mid 左边。
				l = mid
			}
		} else {
			// nums[mid] <  nums[n-1]
			if target < nums[mid] || target > nums[n-1] {
				r = mid
			} else {
				l = mid
			}
		}
	}

	return -1
}
