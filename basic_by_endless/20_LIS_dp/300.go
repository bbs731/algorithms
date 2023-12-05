package dp

import "sort"

/*
https://www.bilibili.com/video/BV1ub411Q7sB/?spm_id_from=333.788&vd_source=84c3c489cf545fafdbeb3b3a6cd6a112
还是看一遍完整的视频吧。重要是知道，思路是怎么来的。

涉及到的不懂的知识点：
1. 如果在时间复杂度上优化 DP 问题， （降维是空间优化）， 需要 "交换"状态和状态值。 （交换之后，这道题可以利用 g s数组的单调性）
2. 贪心和DP到底是什么关系？  本质上，贪心不涉及到重复的子问题， DP 有重复的子问题, 这好像不是本质， 本质是 DP 可以在 子问题上回退， greedy 不行。 贪心的终点是DP，如何理解这句话？

贪心算法，通过局部最优解，得到全局最优解。这个要依赖题目的性质。贪心算法在有最优子结构的问题中尤为有效。最优子结构的意思是问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解。

https://oi-wiki.org/basic/greedy/
Greedy 与动态规划的区别
贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。

好像多说无益，需要多做 greedy 的题目总结出经验。

dp 如果需要降低时间复杂度的进阶技巧是， 交换状态和状态值。
f[i] 换成  g[i], g[i] 代表的是长度为  i+1 LIS 的末尾元素的最小值。


// 解释说明， 这里不能 用单调栈或者单调队列， 因为不能弹出"无用"的元素, 无法判断一个数最后有没有用，所以不能弹出。所以严格意义上来说，
这个确实不是单调队列或者单调栈。 但是，最后这个得到的这个数组 g[i] 确实是递增的，有单调性。 因为，时间复杂度从 n^2 下降到了，n*log
但是没有下降到 n, 因为，只利用到了单调性，没有利用到移除"无用"的元素。
 */
func lengthOfLIS(nums []int) int {
	//可以用，递增的数组 g[i] (不是单调栈或者单调队列）， 把时间复杂度，下降到 O(n*logn)
	//q := []int{0}
	g := []int{}
	for _, x := range nums {
		//pos := sort.Search(len(g), func(k int) bool { return nums[g[k]] >= nums[i] })
		pos := sort.SearchInts(g, x)
		if pos == len(g) {
			g = append(g, x)
		}
		g[pos] = x
	}
	return len(g)
}

func lengthOfLIS(nums []int) int {
	l := 0
	for _, x := range nums {
		pos := sort.SearchInts(nums[:l], x)
		nums[pos] = x
		if pos == l {
			l++
		}
	}
	return l
}

/*
 dfs 翻译为 dp
 翻译成为 DP
 	f[i] = max(f[j]) + 1 // for j < i and nums[i] > nums[j]   看到一个连续区间，求 max 是不是想到了可以用单调栈，单调队列优化？

在没有任何优化的情况下：
时间复杂度是 O(n^2)
 */
func lengthOfLIS_dp(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	ans := 1

	for i, x := range nums {
		f[i] = 1
		for j := 0; j < i; j++ {
			if x > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}

/*
	这是子集型dfs 的问题， 用枚举那个的方法写简单。
	1. 选不选（需要两个角标）
	2. 枚举那个。（这道题简单）


	f[i] 是 包含 i 的， 最长升序的 长度。
	f[i] = max(f[j]） + 1   // for nums[i] > nums[j]
 */
func lengthOfLIS_dfs(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range nums {
		cache[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}
		ans := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				ans = max(ans, dfs(j)+1)
			}
		}
		cache[i] = ans
		return ans
	}

	res := 1
	for i := 0; i < n; i++ {
		res = max(res, dfs(i))
	}
	return res
}
