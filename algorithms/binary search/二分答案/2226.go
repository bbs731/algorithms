package binary_search

import "sort"

/****

给你一个 下标从 0 开始 的整数数组 candies 。数组中的每个元素表示大小为 candies[i] 的一堆糖果。你可以将每堆糖果分成任意数量的 子堆 ，但 无法 再将两堆合并到一起。

另给你一个整数 k 。你需要将这些糖果分配给 k 个小孩，使每个小孩分到 相同 数量的糖果。每个小孩可以拿走 至多一堆 糖果，有些糖果可能会不被分配。

返回每个小孩可以拿走的 最大糖果数目 。



示例 1：

输入：candies = [5,8,6], k = 3
输出：5
解释：可以将 candies[1] 分成大小分别为 5 和 3 的两堆，然后把 candies[2] 分成大小分别为 5 和 1 的两堆。现在就有五堆大小分别为 5、5、3、5 和 1 的糖果。可以把 3 堆大小为 5 的糖果分给 3 个小孩。可以证明无法让每个小孩得到超过 5 颗糖果。
示例 2：

输入：candies = [2,5], k = 11
输出：0
解释：总共有 11 个小孩，但只有 7 颗糖果，但如果要分配糖果的话，必须保证每个小孩至少能得到 1 颗糖果。因此，最后每个小孩都没有得到糖果，答案是 0 。


1 <= candies.length <= 10^5
1 <= candies[i] <= 10^7
1 <= k <= 10^12

 */

// 真是好题啊。

/***
按照灵神给的答案， 按照 sort.Search() 的路子再写一下。
 */

func maximumCandies(candies []int, k int64) int {
	// 这也是一个 先 true 后false 的问题吧。 没有重复的元素，四种写法都可以，不一定非要用 （ ] 区间的写法 (除非有重复的元素）。
	// 你对下面这个 开区间的写法， 返回的值域，还是不确定，分析一下， 对于 【2，5】 k=11 为啥可以正确的返回 0。 按照这道题分析的话，因为初始化 l = 0, 终止
	// 条件是   l+1 < r 那么 l 的值域就是 [0, r-1] 当然可以返回 0
	// 能返回的值域到底是什么？ [0, r-1] 吗？ 对于本题。
	l, r := 0, int(1e12)+1
	for l+1 < r {
		mid := (l + r) >> 1
		tot := 0
		for _, v := range candies {
			tot += v / mid
		}
		if tot < int(k) {
			r = mid
		} else {
			l = mid
		}
	}
	// l + 1
	return l
}

func maximumCandies(candies []int, k int64) int {
	//尝试一下灵神  sort.Search 的写法
	// wow! 这里有套路，还没学会。
	// 灵神说的： // sort.Search 的使用技巧·其一
	// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sort.go#L248

	return sort.Search(int(1e12), func(i int) bool {
		//循环不变量是  f[i-1] = false,  f[i] = f[j] = true
		i++
		tot := 0
		for _, v := range candies {
			tot += v / i
		}
		return tot < int(k) // 这里填什么？ 你最后想要的结果吗？还是取反？
		// 本来想要  >=k, 需要取反，变成  <k
	})
}
