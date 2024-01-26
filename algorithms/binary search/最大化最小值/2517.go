package binary_search

import "sort"

/****

给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。

商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。

返回礼盒的 最大 甜蜜度。



示例 1：

输入：price = [13,5,1,8,21,2], k = 3
输出：8
解释：选出价格分别为 [13,5,21] 的三类糖果。
礼盒的甜蜜度为 min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8 。
可以证明能够取得的最大甜蜜度就是 8 。
示例 2：

输入：price = [1,3,1], k = 2
输出：2
解释：选出价格分别为 [1,3] 的两类糖果。
礼盒的甜蜜度为 min(|1 - 3|) = min(2) = 2 。
可以证明能够取得的最大甜蜜度就是 2 。
示例 3：

输入：price = [7,7,7,7], k = 2
输出：0
解释：从现有的糖果中任选两类糖果，甜蜜度都会是 0 。


提示：

2 <= k <= price.length <= 10^5
1 <= price[i] <= 10^9

 */

func maximumTastiness(price []int, k int) int {
	// 二分答案， 是一个 先 true 后 false 的序列。
	l, r := 0, int(1e9)+1
	n := len(price)
	sort.Ints(price)

	for l+1 < r {
		mid := (l + r) >> 1

		// 思路是，贪心的去找。
		tot := 0
		start := 0
		next := sort.SearchInts(price, price[start]+mid)
		for next < n {
			tot++
			start = next
			next = sort.SearchInts(price, price[start]+mid)
		}

		if tot+1 >= k {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func maximumTastiness(price []int, k int) int {
	// 二分答案， 是一个 先 true 后 false 的序列。
	//l, r := 0, int(1e9)+1
	n := len(price)
	sort.Ints(price)

	// 不用调整值域。
	return sort.Search(int(1e9), func(mid int) bool {
		mid++
		tot := 0
		start := 0
		next := sort.SearchInts(price, price[start]+mid)
		for next < n {
			tot++
			start = next
			next = sort.SearchInts(price, price[start]+mid)
		}

		// 取反  tot +1  >= k
		return tot+1 < k
	})
}
