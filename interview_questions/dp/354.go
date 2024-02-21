package dp

import "sort"

/***
300 题目的升级版本。

和 1691 一样，3D版本
 */

/***
这个会超时啊！
 */
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		//return a[0] < b[0] || a[0] == b[0] && a[1] >= b[1] // 这里写 >= 或者 > 都是可以的， 因为下面的LIS检查，是严格的。
		// O(n^2) 的解法， 不需要考虑 第二个维度。
		return a[0] < b[0]
	})

	ans := 0
	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += 1
		ans = max(ans, f[i])
	}
	return ans
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] >= b[1] // 这里写 >= 或者 > 都是可以的， 因为下面的LIS检查，是严格的。
	})

	g := make([][]int, 0)
	for i := 0; i < n; i++ {
		p := sort.Search(len(g), func(k int) bool {
			return g[k][1] >= envelopes[i][1]
		})
		if p == len(g) {
			g = append(g, envelopes[i])
		} else {
			g[p] = envelopes[i]
		}
	}
	return len(g)
}
