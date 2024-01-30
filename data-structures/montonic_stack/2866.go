package montonic_stack

/****


给你一个长度为 n 下标从 0 开始的整数数组 maxHeights 。

你的任务是在坐标轴上建 n 座塔。第 i 座塔的下标为 i ，高度为 heights[i] 。

如果以下条件满足，我们称这些塔是 美丽 的：

1 <= heights[i] <= maxHeights[i]
heights 是一个 山脉 数组。
如果存在下标 i 满足以下条件，那么我们称数组 heights 是一个 山脉 数组：

对于所有 0 < j <= i ，都有 heights[j - 1] <= heights[j]
对于所有 i <= k < n - 1 ，都有 heights[k + 1] <= heights[k]
请你返回满足 美丽塔 要求的方案中，高度和的最大值 。



示例 1：

输入：maxHeights = [5,3,4,1,1]
输出：13
解释：和最大的美丽塔方案为 heights = [5,3,3,1,1] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山脉数组，峰值在 i = 0 处。
13 是所有美丽塔方案中的最大高度和。
示例 2：

输入：maxHeights = [6,5,3,9,2,7]
输出：22
解释： 和最大的美丽塔方案为 heights = [3,3,3,9,2,2] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山脉数组，峰值在 i = 3 处。
22 是所有美丽塔方案中的最大高度和。
示例 3：

输入：maxHeights = [3,2,5,5,2,3]
输出：18
解释：和最大的美丽塔方案为 heights = [2,2,5,5,2,2] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山脉数组，最大值在 i = 2 处。
注意，在这个方案中，i = 3 也是一个峰值。
18 是所有美丽塔方案中的最大高度和。

提示：
1 <= n == maxHeights <= 10^5
1 <= maxHeights[i] <= 10^9

 */

/***
这个题是不是设计的太牛逼了！

思路：
1. 先用， monotonic stack 找到， 当前点， 左右，比当前点小的数的下标。
2. 设计一下DP    dp_left[i] 表示， 以 ith number 为 peak , 左边的 [0, i-1] 满足条件的 sum和，
	同理计算  dp_right[i]  这个复杂度，是 O(n)
3. 然后枚举每个点。 用 dp_left, dp_right 来得到答案。


恭喜你哈 (还是看了提示，知道用线性DP），但是题目的解，你找到了套路， 单调栈 + 线性的DP
掺杂了两种算法， 这种面试题目， 是不是，也就是 google 的水平了？ 还有比这个更难的吗?

 */

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	left := make([]int, n)
	right := make([]int, n)
	st := []int{-1}
	for i := range right {
		right[i] = n
	}

	for i, v := range maxHeights {
		for len(st) > 1 && maxHeights[st[len(st)-1]] >= v {
			right[st[len(st)-1]] = i
			// pop
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	// 来计算 dp_left
	dp_left := make([]int, n)
	for i, v := range maxHeights {
		if left[i] == -1 {
			dp_left[i] = (i - 0 + 1) * v
		} else {
			dp_left[i] = (i-left[i])*v + dp_left[left[i]]
		}
	}

	dp_right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		v := maxHeights[i]
		if right[i] == n {
			dp_right[i] = (n - 1 - i + 1) * v
		} else {
			dp_right[i] = (right[i]-1-i+1)*v + dp_right[right[i]]
		}
	}

	ans := 0
	// 枚举每个端点
	for i, v := range maxHeights {
		ans = max(ans, dp_left[i]+dp_right[i]-v)
	}
	return int64(ans)
}
