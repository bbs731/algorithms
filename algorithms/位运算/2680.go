package bits_operation

/***

给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。

你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。

a | b 表示两个整数 a 和 b 的 按位或 运算。



示例 1：

输入：nums = [12,9], k = 1
输出：30
解释：如果我们对下标为 1 的元素进行操作，新的数组为 [12,18] 。此时得到最优答案为 12 和 18 的按位或运算的结果，也就是 30 。
示例 2：

输入：nums = [8,1,2], k = 2
输出：35
解释：如果我们对下标 0 处的元素进行操作，得到新数组 [32,1,2] 。此时得到最优答案为 32|1|2 = 35 。

 */

/***
灵神说， 可以做到 O(1)的空间， 是怎么能做到呢？
灵神说思路来自 LC238，接雨水，前后缀分解
 */
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	preor := make([]int, n+1) // preor 可以随着循环计算，不用开这个数组
	posor := make([]int, n+1)

	//preor[0] = nums[0] // 前嘴和，不需要初始化 presum[0], 这里面犯错误了
	for i := 1; i <= n; i++ {
		preor[i] = preor[i-1] | nums[i-1]
	}

	//posor[n] = 0
	for i := n - 2; i >= 0; i-- {
		posor[i] = posor[i+1] | nums[i+1]
	}

	ans := 0
	for i, num := range nums {
		ans = max(ans, num<<k|preor[i]|posor[i])
	}
	return int64(ans)
}
