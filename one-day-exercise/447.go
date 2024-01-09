package one_day_exercise

/*

给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）。

返回平面上所有回旋镖的数量。


示例 1：

输入：points = [[0,0],[1,0],[2,0]]
输出：2
解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
示例 2：

输入：points = [[1,1],[2,2],[3,3]]
输出：2
示例 3：

输入：points = [[1,1]]
输出：0

 */

func ek_distance(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	ans := 0

	for i, p1 := range points {
		dm := make(map[int]int)
		for j := 0; j < n; j++ {
			if j != i {
				d := ek_distance(p1, points[j])
				//sd[distance(p1, points[j])]++
				dm[d]++
				ans += (dm[d] - 1) * 2
			}
		}
	}
	return ans
}
