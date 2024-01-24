package binary_search

/***
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。



示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10


提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105
 */
/***
https://leetcode.cn/problems/single-element-in-a-sorted-array/solutions/1264496/gong-shui-san-xie-er-duan-xing-fen-xi-yu-17nv/

宫三的题解，写的非常的好！ 里面还是有好多技巧。
 */
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	l, r := -1, n

	for l+1 < r {
		mid := (l + r) >> 1
		if mid&1 == 0 {
			if mid+1 <= n-1 && nums[mid+1] == nums[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			if mid-1 >= 0 && nums[mid-1] == nums[mid] {
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	// l + 1 = r
	return nums[r]
}

// 感觉，写的还是有点问题. 你的假设是，题目给的一定符合条件，有一个不满足的数，但是他骗了你咋办？ 你写的循环能结束不？
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	l, r := -1, len(nums)

	for l+1 < r {
		mid := (l + r) >> 1
		if mid-1 >= 0 && nums[mid-1] == nums[mid] {
			if (n-1-mid)&1 == 0 {
				r = mid - 1
			} else {
				l = mid
			}
		} else if mid+1 <= n-1 && nums[mid] == nums[mid+1] {
			if (n-mid-2)&1 == 0 {
				r = mid
			} else {
				l = mid + 1
			}

		} else {
			return nums[mid]
		}
	}
	return nums[l]
}
