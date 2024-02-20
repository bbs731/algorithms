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

	sort.Slice(cuboids, func (i, j int) bool {
		if cuboids[i][0] < cuboids[j][0] {
			return true
		}
		if cuboids[i][0] == cuboids [j][0] {
			if cuboids[i][1] < cuboids[j][1] {
				return true
			}


		}
		return false
	})
	g := make([][]int, 0)




}
