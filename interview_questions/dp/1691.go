package dp

import "sort"

/***

300 LIS
354 俄罗斯套娃
1691 stacking box

这是啥逻辑？
现在变成 3D 的套娃！
 */

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	for _, c := range cuboids {
		sort.Ints(c)
	}

	/***
		这个 sorting 的顺序非常的讲究啊！ 哎！！！！！！！
	 */
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][2] < cuboids[j][2] {
			return true
		}
		if cuboids[i][2] == cuboids [j][2] {
			if cuboids[i][1] < cuboids[j][1] {
				return true
			}
			if cuboids[i][1] == cuboids[j][1] {
				return cuboids[i][0] < cuboids[j][0]
			}
		}
		return false
	})

	ans := 0
	f := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] = f[i] + cuboids[i][2]
		ans = max(ans, f[i])
	}
	return ans
}
