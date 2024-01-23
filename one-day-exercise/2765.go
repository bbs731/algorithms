package one_day_exercise

/**

给你一个下标从 0 开始的整数数组 nums 。如果 nums 中长度为 m 的子数组 s 满足以下条件，我们称它是一个 交替子数组 ：

m 大于 1 。
s1 = s0 + 1 。
下标从 0 开始的子数组 s 与数组 [s0, s1, s0, s1,...,s(m-1) % 2] 一样。也就是说，s1 - s0 = 1 ，s2 - s1 = -1 ，s3 - s2 = 1 ，s4 - s3 = -1 ，以此类推，直到 s[m - 1] - s[m - 2] = (-1)m 。
请你返回 nums 中所有 交替 子数组中，最长的长度，如果不存在交替子数组，请你返回 -1 。

子数组是一个数组中一段连续 非空 的元素序列。



示例 1：

输入：nums = [2,3,4,3,4]
输出：4
解释：交替子数组有 [3,4] ，[3,4,3] 和 [3,4,3,4] 。最长的子数组为 [3,4,3,4] ，长度为4 。
示例 2：

输入：nums = [4,5,6]
输出：2
解释：[4,5] 和 [5,6] 是仅有的两个交替子数组。它们长度都为 2 。


[1,29,30,5]
 */

//https://leetcode.cn/problems/longest-alternating-subarray/solutions/2615916/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/
// 这个灵神的代码太好了！
/***
适用场景：按照题目要求，数组会被分割成若干组，且每一组的判断/处理逻辑是一样的。

核心思想：

1. 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。

2. 内层循环负责遍历组，找出这一组最远在哪结束。
这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

对于本题来说，在内层循环时，假设这一组的第一个数是 333，那么这一组的数字必须形如 3,4,3,4,⋯3,4,3,4,\cdots3,4,3,4,⋯，也就是

nums[i]=nums[i−2]
另外，对于 [3,4,3,4,5,4,5][3,4,3,4,5,4,5][3,4,3,4,5,4,5] 这样的数组，第一组交替子数组为 [3,4,3,4][3,4,3,4][3,4,3,4]，第二组交替子数组为 [4,5,4,5][4,5,4,5][4,5,4,5]，这两组有一个数是重叠的，所以下面代码在外层循环末尾要把 iii 减一。
 */

func alternatingSubarray(nums []int) int {
	ans := -1
	i, n := 0, len(nums)
	for i < n-1 {
		if nums[i+1]-nums[i] != 1 {
			i++ // 直接跳过
			continue
		}
		i0 := i // 记录这一组的开始位置
		i += 2  // i 和 i+1 已经满足要求，从 i+2 开始判断
		for i < n && nums[i] == nums[i-2] {
			i++
		}
		// 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-i0)
		i--
	}
	return ans
}

//这不是一道简单的题， 是一道很好的面试题目。
// 你已经晕了吗？休息一下吧。 复杂度是O(n^2)
func alternatingSubarray(nums []int) int {
	ans := -1
	n := len(nums)

	for i := 0; i < n; i++ {
		d := 1
		j := i
		for j+1 < n {
			if nums[j]+d == nums[j+1] {
				j++
				d = d * -1
				ans = max(ans, j-i+1)
			} else {
				break
			}
		}
	}
	return ans
}
