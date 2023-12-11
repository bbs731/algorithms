package sliding_window

/*
//Given an integer array nums, return the number of triplets chosen from the arr
//ay that can make triangles if we take them as side lengths of a triangle.
//
//
// Example 1:
//
//
//Input: nums = [2,2,3,4]
//Output: 3
//Explanation: Valid combinations are:
//2,3,4 (using the first 2)
//2,3,4 (using the second 2)
//2,2,3
//
//
// Example 2:
//
//
//Input: nums = [4,2,3,4]
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
//
 */

/*
时间复杂度是 O（n^2)
 */
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0

	for i := 0; i < n-2; i++ {
		x := nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			if x+nums[left] > nums[right] {
				// valid triangle
				ans += right - left
				right--
			} else {
				// 在这个分支里面，可以选择更小的 right, 也可以选择更大的 left. 两种选择。
				// 1. 我们计算更小的 right
				for k := right - 1; k > left; k-- {
					if x+nums[left] > nums[k] {
						ans += k - left
						break
					}
				}
				// 2. 更大的 left, 但是 ans 放在外层Loop 去统计。
				left++
			}
		}
	}
	return ans
}
