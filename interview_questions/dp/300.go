package dp

import "sort"

/****
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1
 */

/***
O（n^2) 的解法。 还有贪心 O(n*logn) 的解法

看 1691 3D 的版本
 */

func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += 1
		ans = max(ans, f[i])
	}
	return ans
}

/***
优化的版本, 属于， 用单调栈，优化DP 的范围吗？

感觉，脑袋不是特别的好使了。
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	g := []int{}

	for i := 0; i < n; i++ {
		p := sort.SearchInts(g, nums[i])
		if p == len(g) {
			g = append(g, nums[i])
		} else {
			g[p] = nums[i]
		}
	}
	return len(g)
}
