package binary_search

import "sort"

/****

珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。

珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。



示例 1：

输入：piles = [3,6,7,11], h = 8
输出：4
示例 2：

输入：piles = [30,11,23,4,20], h = 5
输出：30
示例 3：

输入：piles = [30,11,23,4,20], h = 6
输出：23


提示：

1 <= piles.length <= 10^4
piles.length <= h <= 10^9
1 <= piles[i] <= 10^9

 */

/***
两种解法， 一遍过！ 赞！
 */

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	l, r := 0, n*int(1e9)+1

	for l+1 < r {
		mid := (l + r) >> 1
		tot := 0

		for _, v := range piles {
			tot += (v + (mid - 1)) / mid
		}

		if tot > h {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	r := n*int(1e9) + 1

	// sort.Search 的使用技巧·其二  指定上下界  [l, r)
	// 这个问题是 先 false, 后 true 的序列， 所以，正常返回期望的结果就好，不需要考虑取反
	return 1 + sort.Search(r-1, func(x int) bool {
		x += 1
		tot := 0
		for _, v := range piles {
			tot += (v + (x - 1)) / x
		}
		return tot <= h
	})
}
