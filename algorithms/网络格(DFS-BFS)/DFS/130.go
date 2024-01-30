package DFS

/***
感觉会和 1254一样恶心。
吸取了 1254 的教训， 使用 global 的 variable, res and components
 */

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	components := make([][2]int, 0)
	res := true
	var dfs func(int, int)
	dfs = func(x, y int) {
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		components = append(components, [2]int{x, y})

		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				res = false
				continue
			}
			if board[r][c] == 'O' {
				dfs(r, c)
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' && visited[i][j] == false {
				res = true
				dfs(i, j)
				if res == true {
					for _, c := range components {
						board[c[0]][c[1]] = 'X'
					}
				}
				//clear(components)  // 这个就是错误的，为什么？
				components = components[:0]
			}
		}
	}
	return
}
