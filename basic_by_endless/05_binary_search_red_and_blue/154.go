package binary_search

//Suppose an array of length n sorted in ascending order is rotated between 1 an
//d n times. For example, the array nums = [0,1,4,4,5,6,7] might become:
//
//
// [4,5,6,7,0,1,4] if it was rotated 4 times.
// [0,1,4,4,5,6,7] if it was rotated 7 times.
//
//
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
//in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
// Given the sorted rotated array nums that may contain duplicates, return the m
//inimum element of this array.
//
// You must decrease the overall operation steps as much as possible.
//
//
// Example 1:
// Input: nums = [1,3,5]
//Output: 1
// Example 2:
// Input: nums = [2,2,2,0,1]
//Output: 0
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// nums is sorted and rotated between 1 and n times.
//
//
//
// Follow up: This problem is similar to Find Minimum in Rotated Sorted Array, b
//ut nums may contain duplicates. Would this affect the runtime complexity? How an
//d why?
//
//

/*
灵神的答案。
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/solutions/2131553/zhi-yao-ni-hui-153-jiu-neng-kan-dong-pyt-qqc6/
 */
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-2

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[right+1] { // 这里是画龙点睛的地方， 为什么要和  right+1 的位置比较？
			right = mid - 1
		} else if nums[mid] > nums[right+1] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}

/*
对的，最坏的情况下，会退化成 O(n) 的时间复杂度。

这个答案也是对的，就是题解稍复杂！
 */

func findMin_chunlei(nums []int) int {
	n := len(nums)
	left, right := 0, n-2 //[left, right]

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else if nums[mid] < nums[n-1] {
			right = mid - 1
		} else {
			//看左边  这里都是自己的想法，思路还是挺严谨的。就是代码偏多了。 看看灵神给的答案。
			if nums[left] > nums[mid] {
				left++
				right = mid - 1
			} else if nums[left] < nums[mid] {
				return nums[left]
			} else {
				// left == mid
				found := false
				for i := left + 1; i < mid; i++ {
					if nums[i] != nums[left] {
						found = true
						break
					}
				}
				if found {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	// l = r + 1
	return nums[left]
}
