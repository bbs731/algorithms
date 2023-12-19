package _76

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

你可以对数组执行 至多 k 次操作：

从数组中选择一个下标 i ，将 nums[i] 增加 或者 减少 1 。
最终数组的频率分数定义为数组中众数的 频率 。

请你返回你可以得到的 最大 频率分数。

众数指的是数组中出现次数最多的数。一个元素的频率指的是数组中这个元素的出现次数。



示例 1：

输入：nums = [1,2,6,4], k = 3
输出：3
解释：我们可以对数组执行以下操作：
- 选择 i = 0 ，将 nums[0] 增加 1 。得到数组 [2,2,6,4] 。
- 选择 i = 3 ，将 nums[3] 减少 1 ，得到数组 [2,2,6,3] 。
- 选择 i = 3 ，将 nums[3] 减少 1 ，得到数组 [2,2,6,2] 。
元素 2 是最终数组中的众数，出现了 3 次，所以频率分数为 3 。
3 是所有可行方案里的最大频率分数。


示例 2：

输入：nums = [1,4,4,2,4], k = 0
输出：3
解释：我们无法执行任何操作，所以得到的频率分数是原数组中众数的频率 3 。
 */

/*

这道题的原理，是中位数的定理：(中位数贪心）
可以详细的证明，中位数到，所有的其他的点的距离和最短。具体看灵神 376周赛的视频证明。

题解： https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/solutions/2569301/hua-dong-chuang-kou-zhong-wei-shu-tan-xi-nuvr/

 */

/* 必须优化，要不然时间复杂度过不去  O(n)的时间复杂度。*/
func calCost(nums []int, l, r int) int {
	percentile50 := nums[(r-l+1)/2]

	cost := 0
	for i := l; i <= r; i++ {
		if nums[i] > percentile50 {
			cost += nums[i] - percentile50
		} else {
			cost += percentile50 - nums[i]
		}
	}
	return cost
}

// O(1)的时间复杂度。利用了前缀和。
func calCost_use_prefix(psums []int, nums []int, l, r int) int {
	cost := 0
	mid := (r + l) / 2 // 这里是个坑， 中位数有两个或者一个， 左边= (r+l)/2  右边 (r+l+1)/2   n 是奇数的时候，左边和右边的中位数是一个数。

	cost += (mid-l)*nums[mid] - (psums[mid] - psums[l])
	cost += psums[r+1] - psums[mid+1] - nums[mid]*(r-mid)
	return cost
}

/*

问题：
	1. 为什么能想到用 sliding-window 滑动窗口来解决这道问题？
	2. 满足什么样的条件？(单调性） 才能使用滑动窗口？
 */

func maxFrequencyScore(nums []int, k int64) int {
	ans := 0
	left, right := 0, 0 // 这里， right 应该初始化成 0， 不能是 1.  测试用例   [10]  k=0  应该返回 1
	sort.Ints(nums)
	psums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		psums[i] = psums[i-1] + nums[i-1]
	}
	for ; right < len(nums); right++ { // 枚举 window's right point
		for left <= right && calCost_use_prefix(psums, nums, left, right) > int(k) {
			left += 1 // 移动左端点
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
