package BFS

/***

给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。


输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0,0],[0,1,0],[0,0,0]]

输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
输出：[[0,0,0],[0,1,0],[1,2,1]]
 */

type pair struct {
	x, y, dist int
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)

	for i := range res {
		res[i] = make([]int, n)
	}

	q := []pair{}
	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 0 {
				q = append(q, pair{i, j, 0})
			}
		}
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			x, y := p.x, p.y
			dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || mat[r][c] == 0 {
					continue
				}
				mat[r][c] = 0 // mark as visited
				q = append(q, pair{r, c, p.dist + 1})
				res[r][c] = p.dist + 1
			}
		}
	}

	return res
}
