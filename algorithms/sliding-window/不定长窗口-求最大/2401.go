package weekly

/***

给你一个由 正 整数组成的数组 nums 。

如果 nums 的子数组中位于 不同 位置的每对元素按位 与（AND）运算的结果等于 0 ，则称该子数组为 优雅 子数组。

返回 最长 的优雅子数组的长度。

子数组 是数组中的一个 连续 部分。

注意：长度为 1 的子数组始终视作优雅子数组。



示例 1：

输入：nums = [1,3,8,48,10]
输出：3
解释：最长的优雅子数组是 [3,8,48] 。子数组满足题目条件：
- 3 AND 8 = 0
- 3 AND 48 = 0
- 8 AND 48 = 0
可以证明不存在更长的优雅子数组，所以返回 3 。
示例 2：

输入：nums = [3,1,5,11,13]
输出：1
解释：最长的优雅子数组长度为 1 ，任何长度为 1 的子数组都满足题目条件。

 */

// 很好的题目，如果不提示 sliding window 的话，感觉是做不成来的。

func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 { // 有交集
			or ^= nums[left] // 从 or 中去掉集合 nums[left]
			left += 1
		}
		or |= x // 把集合 x 并入 or 中
		ans = max(ans, right-left+1)
	}
	return
}

// 下面的代码是时间复杂度实际上，最坏情况下是 O（n^2) 有 O(n) 的写法。
// 如果 leetcode 想 OI 一样卡时间的话，可能就过不去了。
func longestNiceSubarray(nums []int) int {
	n := len(nums)
	ans := 1
	l := 0

	for i := 1; i < n; i++ {
		x := nums[i]
		j := i - 1
		for ; j >= l; j-- {
			if nums[j]&x != 0 {
				break
			}
		}
		l = j + 1
		// 两种情况下，都能归纳为 l = j+1
		//if j == i-1 {
		//	l = i
		//} else {
		//	l = j + 1
		//}
		ans = max(ans, i-l+1)
	}
	return ans
}
