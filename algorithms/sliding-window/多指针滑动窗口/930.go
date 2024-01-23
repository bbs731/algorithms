package sliding_window

/***

给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。

子数组 是数组的一段连续部分。


示例 1：

输入：nums = [1,0,1,0,1], goal = 2
输出：4
解释：
有 4 个满足题目要求的子数组：[1,0,1]、[1,0,1,0]、[0,1,0,1]、[1,0,1]
示例 2：

输入：nums = [0,0,0,0,0], goal = 0
输出：15


 */
// 真是，这种没见过的套路， 面试第一次见到肯定谁死翘翘
// 题目太打击人， 想了半天做不出来，但是题解却特别简单。

/***
https://leetcode.cn/problems/binary-subarrays-with-sum/solutions/864087/he-xiang-tong-de-er-yuan-zi-shu-zu-by-le-5caf/


官网上给的 滑动窗口的题解，足够清晰。 利用了: 固定右端点的情况下， 题目的解（左端点）是在一个连续的区间内。 这一个非常重要的性质。 因此解题的关键，就是如何求出这个解区间的两个端点了， 这也是为什么这里题目，被划分为
多指针窗口问题， 解的空间也是一段窗口。

930, 1248 是一个套路。

  */

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0

	for right, v := range nums {
		sum1 += v
		sum2 += v

		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}

		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}

		ans += left2 - left1
	}
	return ans
}
