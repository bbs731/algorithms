package binary_search

import (
	"math"
	"sort"
)

/***
给你一个浮点数 hour ，表示你到达办公室可用的总通勤时间。要到达办公室，你必须按给定次序乘坐 n 趟列车。另给你一个长度为 n 的整数数组 dist ，其中 dist[i] 表示第 i 趟列车的行驶距离（单位是千米）。

每趟列车均只能在整点发车，所以你可能需要在两趟列车之间等待一段时间。

例如，第 1 趟列车需要 1.5 小时，那你必须再等待 0.5 小时，搭乘在第 2 小时发车的第 2 趟列车。
返回能满足你准时到达办公室所要求全部列车的 最小正整数 时速（单位：千米每小时），如果无法准时到达，则返回 -1 。

生成的测试用例保证答案不超过 10^7 ，且 hour 的 小数点后最多存在两位数字 。



示例 1：

输入：dist = [1,3,2], hour = 6
输出：1
解释：速度为 1 时：
- 第 1 趟列车运行需要 1/1 = 1 小时。
- 由于是在整数时间到达，可以立即换乘在第 1 小时发车的列车。第 2 趟列车运行需要 3/1 = 3 小时。
- 由于是在整数时间到达，可以立即换乘在第 4 小时发车的列车。第 3 趟列车运行需要 2/1 = 2 小时。
- 你将会恰好在第 6 小时到达。
示例 2：

输入：dist = [1,3,2], hour = 2.7
输出：3
解释：速度为 3 时：
- 第 1 趟列车运行需要 1/3 = 0.33333 小时。
- 由于不是在整数时间到达，故需要等待至第 1 小时才能搭乘列车。第 2 趟列车运行需要 3/3 = 1 小时。
- 由于是在整数时间到达，可以立即换乘在第 2 小时发车的列车。第 3 趟列车运行需要 2/3 = 0.66667 小时。
- 你将会在第 2.66667 小时到达。
示例 3：

输入：dist = [1,3,2], hour = 1.9
输出：-1
解释：不可能准时到达，因为第 3 趟列车最早是在第 2 小时发车。
 */

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	l, r := 0, int(1e7)+10

	for l+1 < r {
		mid := (l + r) >> 1
		tot := float64(0)

		for i := 0; i < n-1; i++ {
			tot += float64((dist[i] + mid - 1) / mid)
		}
		tot += float64(dist[n-1]) / float64(mid)

		if tot > hour {
			l = mid
		} else {
			r = mid
		}
	}
	// l +1 == r
	// 这里取巧了，在不可能的情况下，扩大了一下取值范围，那么，如果根本满足不了的话，再扩大也没有用。用这个条件来判断 -1 的情况！
	if r > int(1e7) {
		return -1
	}
	return r
}

/***


func minSpeedOnTime(dist []int, hour float64) int {
	h100 := int(math.Round(hour * 100))
	n := len(dist)
	if h100 <= (n-1)*100 { // hour 必须严格大于 n-1
		return -1
	}
	// 这里用到了。 sort.Search 的技巧二 ：
	// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sort.go#L287
	// 搜索的区间在 [l, r)
	return 1 + sort.Search(1e7-1, func(v int) bool {
		v++
		h := n - 1
		for _, d := range dist[:n-1] {
			h += (d - 1) / v
		}
		return (h*v+dist[n-1])*100 <= h100*v
	})
}

 */

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	h100 := int(math.Round(hour * 100))
	if h100 <= (n-1)*100 {
		return -1
	}

	// 然后我们可以确定搜索的区间变为 [1, r)
	l, r := 1, int(1e7)

	// 谢谢灵神！ 又学到了一种 sort.Search 的技巧！
	// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sort.go#L287
	return l + sort.Search(r-l, func(x int) bool {
		x += l

		tot := float64(0)
		for i := 0; i < n-1; i++ {
			tot += float64((dist[i] + x - 1) / x)
		}
		tot += float64(dist[n-1]) / float64(x)
		return tot <= hour
	})
}
