package binary_search

import "sort"

/***

传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。


示例 1：

输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
示例 2：

输入：weights = [3,2,2,4,1,4], days = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4
示例 3：

输入：weights = [1,2,3,1,1], days = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1


提示：

1 <= days <= weights.length <= 5 * 10^4
1 <= weights[i] <= 500

 */

func shipWithinDays(weights []int, days int) int {
	n := len(weights)
	// 你的上界设置的太粗糙了！
	l, r := 0, n*days*5*int(1e4)+1

	// 先 false 后 true
	for l+1 < r {
		mid := (l + r) >> 1
		tot := 1 // 这里还容易出错了！
		not_enough := false
		left := mid
		for _, w := range weights {
			left -= w
			if left < 0 {
				left = mid // 还有这里， 连续两个 bug
				if left < w {
					not_enough = true
					break
				}
				left = left - w
				tot++
			}
		}
		if not_enough || tot > days {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func shipWithinDays(weights []int, days int) int {
	n := len(weights)
	// 你的上界设置的太粗糙了！
	r := n*days*5*int(1e4) + 1

	// [1, r)  搜索的是这个区间， 先 false, 后true
	return 1 + sort.Search(r-1, func(x int) bool {
		// x += l
		x += 1 // 这里不要拉掉了！非常的重要！这行保证了，x 取值是从 1开始的，而且保证了 x 不可能取到 0
		tot := 1
		not_enough := false
		left := x
		for _, w := range weights {
			left -= w
			if left < 0 {
				left = x
				if left < w {
					not_enough = true
					break
				}
				left = left - w
				tot++
			}
		}
		// 如果不取反的话， 就是正常写，你期望的结果，就可以了。 （按照期望的逻辑去写）
		if not_enough || tot > days {
			return false
		}
		return true
	})
}
