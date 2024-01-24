package binary_search

import "sort"

/***

给你一个按 非递减顺序 排列的数组 nums ，返回正整数数目和负整数数目中的最大值。

换句话讲，如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值。
注意：0 既不是正整数也不是负整数。



示例 1：

输入：nums = [-2,-1,-1,1,2,3]
输出：3
解释：共有 3 个正整数和 3 个负整数。计数得到的最大值是 3 。
示例 2：

输入：nums = [-3,-2,-1,0,0,1,2]
输出：3
解释：共有 2 个正整数和 3 个负整数。计数得到的最大值是 3 。
示例 3：

输入：nums = [5,20,66,1314]
输出：4
解释：共有 4 个正整数和 0 个负整数。计数得到的最大值是 4

 */

/***

灵神的答案，也是太强了！
func maximumCount(nums []int) int {
   return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}


对，你和灵神考虑这道题的差距是， 你不了解， SearchInts() 返回的边界掉件是什么？

[0,0,0,0,0]
sort.SearchInts(nums, 1) 的时候，会返回数组的长度 n

最后可以总结为，  sort.SearchInts(nums, 0) 在计算，所有负数的长度。
n- sort.SearchInts(nums, 1) 在计算正数的长度。

边界条件都满足。

 */
func maximumCount(nums []int) int {
	n := len(nums)
	nl, pl := 0, 0

	if nums[n-1] > 0 {
		// find pos, that nums[pos] > 0
		a := sort.SearchInts(nums, 1)
		pl = n - a
	}

	if nums[0] < 0 {
		b := sort.SearchInts(nums, 0) - 1
		nl = b + 1
	}
	return max(nl, pl)
}
