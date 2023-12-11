package binary_search

//A peak element is an element that is strictly greater than its neighbors.
//
// Given a 0-indexed integer array nums, find a peak element, and return its ind
//ex. If the array contains multiple peaks, return the index to any of the peaks.
//
//
// You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is a
//lways considered to be strictly greater than a neighbor that is outside the arra
//y.
//
// You must write an algorithm that runs in O(log n) time.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,1]
//Output: 2
//Explanation: 3 is a peak element and your function should return the index num
//ber 2.
//
// Example 2:
//
//
//Input: nums = [1,2,1,3,5,6,4]
//Output: 5
//Explanation: Your function can return either index number 1 where the peak ele
//ment is 2, or index number 5 where the peak element is 6.
//
//
// Constraints:
//
//
// 1 <= nums.length <= 1000
// -231 <= nums[i] <= 231 - 1
// nums[i] != nums[i + 1] for all valid i.

func findPeakElement(nums []int) int {
	// 定义， 红色是山峰的左侧， 蓝色是山峰或者山峰的右侧
	n := len(nums)
	left, right := 0, n-2 // [left, right]  // 这里初始化成 n-2 因为 n-1 一定是蓝色的。 因为 n-1 要么是山峰，要么是山峰的右侧
	// 这个初始化很有技巧， 它简化了，下面 nums[mid+1] 是否越界的判断。

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left // left = right+1
}
