package montonic_stack

/****

一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。

比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对  109 + 7 取余 的结果。

请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。

子数组 定义为一个数组的 连续 部分。



示例 1：

输入：nums = [1,2,3,2]
输出：14
解释：最小乘积的最大值由子数组 [2,3,2] （最小值是 2）得到。
2 * (2+3+2) = 2 * 7 = 14 。
示例 2：

输入：nums = [2,3,3,1,2]
输出：18
解释：最小乘积的最大值由子数组 [3,3] （最小值是 3）得到。
3 * (3+3) = 3 * 6 = 18 。
示例 3：

输入：nums = [3,1,5,6,4,2]
输出：60
解释：最小乘积的最大值由子数组 [5,6,4] （最小值是 4）得到。
4 * (5+6+4) = 4 * 15 = 60 。


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^7
 */

/***
这个第一眼，感觉，就会啊。 是不是 montonic_stack +  前嘴和

思路：
枚举每一个位置
1. 用 montonic stack 得到 满足的区间的范围  (这里求 小于当前数的，左右端点）
2. 用前缀和 知道区间的和。
3. 两者相乘， 得到结果，记录最大值。
 */

func maxSumMinProduct(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
		right[i-1] = n //蹭一下 , bug 就是这么来的？哈哈！
	}

	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] >= v {
			right[st[len(st)-1]] = i
			// pop stack
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	ans := 0
	for i, v := range nums {
		ans = max(ans, v*(sum[right[i]]-sum[left[i]+1]))
	}
	return ans % (int(1e9) + 7)
}
