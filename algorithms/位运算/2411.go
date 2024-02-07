package bits_operation

/***
给你一个长度为 n 下标从 0 开始的数组 nums ，数组中所有数字均为非负整数。对于 0 到 n - 1 之间的每一个下标 i ，你需要找出 nums 中一个 最小 非空子数组，它的起始位置为 i （包含这个位置），同时有 最大 的 按位或运算值 。

换言之，令 Bij 表示子数组 nums[i...j] 的按位或运算的结果，你需要找到一个起始位置为 i 的最小子数组，这个子数组的按位或运算的结果等于 max(Bik) ，其中 i <= k <= n - 1 。
一个数组的按位或运算值是这个数组里所有数字按位或运算的结果。

请你返回一个大小为 n 的整数数组 answer，其中 answer[i]是开始位置为 i ，按位或运算结果最大，且 最短 子数组的长度。

子数组 是数组里一段连续非空元素组成的序列。



示例 1：

输入：nums = [1,0,2,1,3]
输出：[3,3,2,2,1]
解释：
任何位置开始，最大按位或运算的结果都是 3 。
- 下标 0 处，能得到结果 3 的最短子数组是 [1,0,2] 。
- 下标 1 处，能得到结果 3 的最短子数组是 [0,2,1] 。
- 下标 2 处，能得到结果 3 的最短子数组是 [2,1] 。
- 下标 3 处，能得到结果 3 的最短子数组是 [1,3] 。
- 下标 4 处，能得到结果 3 的最短子数组是 [3] 。
所以我们返回 [3,3,2,2,1] 。
示例 2：

输入：nums = [1,2]
输出：[2,1]
解释：
下标 0 处，能得到最大按位或运算值的最短子数组长度为 2 。
下标 1 处，能得到最大按位或运算值的最短子数组长度为 1 。
所以我们返回 [2,1] 。

 */

/***
怎么，题都这么难。
这题好难啊！

先写 O(n^2) 的答案， 然后再继续优化。

太难了，  O(n^2) 的写法，都写不对，哎！
*/

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, x := range nums {
		ans[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j]|x != nums[j] {
				nums[j] = nums[j] | x
				ans[j] = i - j + 1
			}
		}
	}
	return ans
}

/**
优化，利用了 nums 的值域   <= 1e9  (=2^29)

所以时间复杂度是 O (n*29)
 */
func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, x := range nums {
		ans[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j]|x != nums[j] {
				nums[j] = nums[j] | x
				ans[j] = i - j + 1
			} else {
				// 优化的地方： 如果 x 是 nums[j] 的子集， 可以提前结束。
				// 优化，优化！优化，就是不等式， 优化就是可以剪枝，可以提前结束。
				break
			}
		}
	}
	return ans
}

/***
https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solutions/1830911/by-endlesscheng-zai1/

下面的这个模板代码， 实际上，自己推到的时候，想了1个多小时，过程是推演过的， 但是不知道用什么数据结构存储,
对于，需要存储的规模没有概念。 估计过一段时间，也是看不懂的。


该模板可以做到

1. 求出所有子数组的按位或的结果，以及值等于该结果的子数组的个数。
2. 求按位或结果等于任意给定数字的子数组的最短长度/最长长度。

 */

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	type pair struct {
		or int // 按位于的值
		i  int // 对应子数组的右端点的最小值。
	}
	ors := make([]pair, 0)
	for i := n - 1

}
