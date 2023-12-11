package binary_search

//There is an integer array nums sorted in ascending order (with distinct values
//).
//
// Prior to being passed to your function, nums is possibly rotated at an unknow
//n pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k]
//, nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For
//example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0
//,1,2].
//
// Given the array nums after the possible rotation and an integer target, retur
//n the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1
// Example 3:
// Input: nums = [1], target = 0
//Output: -1
//

/*
这是好题， 太锻炼思维了。


红蓝染色的意义是什么？
找到目标元素，目标元素的右边都是蓝色, 目标元素的左边都是红色。 目标元素，是红是蓝？


红蓝染色，太难考虑了。不过学会了红蓝染色法，这个将是一劳永逸的方法。 如果按照红蓝染色的方法分析，那么，如何定义蓝色红色？ 蓝色是指，大于等于 target 的区间是蓝色吗？不对。
是 target 元素和 target 右边的元素被染成蓝色。 （右边的元素，可以小于 target)

https://leetcode.cn/problems/search-in-rotated-sorted-array/solutions/1987503/by-endlesscheng-auuh/

用灵神的做法，再做一遍， 再看一遍视频。 二分，永远的痛！

 */

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-2 // [left, right]

	var isBlue func(int) bool
	isBlue = func(mid int) bool {
		if nums[mid] > nums[n-1] {
			return target > nums[n-1] && nums[mid] >= target
		}
		// nums[mid] < nums[n-1]
		return target <= nums[mid] || target > nums[n-1]
	}

	for left <= right {
		mid := (left + right) / 2

		if isBlue(mid) {
			right = mid - 1
		} else {
			left = left + 1
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func search_chunlei(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1 // [left, right]

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target == nums[left] {
			return left
		}
		if target == nums[right] {
			return right
		}
		if nums[mid] > nums[n-1] {
			if target < nums[mid] && target > nums[n-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // mid < n-1
			if target > nums[mid] && target < nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
