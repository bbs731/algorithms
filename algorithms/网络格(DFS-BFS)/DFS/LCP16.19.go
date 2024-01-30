package DFS

import "sort"

/****

你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。

示例：

输入：
[
  [0,2,1,0],
  [0,1,0,1],
  [1,1,0,1],
  [0,1,0,1]
]
输出： [1,2,4]
提示：

0 < len(land) <= 1000
0 < len(land[i]) <= 1000
 */

/***

https://leetcode.cn/problems/pond-sizes-lcci/solutions/2316704/mo-ban-wang-ge-tu-dfsfu-ti-dan-by-endles-p0n1/
灵神给的， DFS 的模板！

DFS 用来处理 网络格的连通性问题。
 */
func pondSizes(land [][]int) (ans []int) {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		// 非法的检查，写在 dfs 的第一行， 这个算是需要记住的技巧吗？  通过一些题目，觉得，写在下面 call dfs 之前做检查比较好！（个人感觉）
		if x < 0 || x >= m || y < 0 || y >= n || land[x][y] != 0 {
			return 0
		}
		land[x][y] = 1 // 标记 (x,y) 被访问，避免重复访问
		cnt0 := 1
		for i := x - 1; i <= x+1; i++ { // 访问八方向的 0
			for j := y - 1; j <= y+1; j++ {
				cnt0 += dfs(i, j)
			}
		}
		return cnt0
	}
	for i, row := range land {
		for j, x := range row {
			if x == 0 { // 从没有访问过的 0 出发
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	return
}

func pondSizes(land [][]int) []int {
	m := len(land)
	n := len(land[0])
	visited := make([][]bool, m)
	for i := range land {
		visited[i] = make([]bool, n)
	}
	dir := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	var dfs func(int, int) int
	dfs = func(i, j int) (cnts int) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == true || land[i][j] != 0 {
			return
		}
		visited[i][j] = true
		cnts = 1
		for _, d := range dir {
			cnts += dfs(i+d[0], j+d[1])
		}
		return cnts
	}

	ans := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 && visited[i][j] == false {
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	return ans
}
