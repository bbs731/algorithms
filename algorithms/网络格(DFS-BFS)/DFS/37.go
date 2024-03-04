package DFS

func solveSudoku(board [][]byte) {
	n := len(board)
	//ans := make([][]byte, n)
	//res := make([][]byte, n)
	colsm := make([]int, n)
	rowsm := make([]int, n)
	blocks := [3][3]int{}

	cnts := 0

	for i := range board {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				cnts++
				continue
			}
			d := 1 << (board[i][j] - '1')

			colsm[j] |= d
			rowsm[i] |= d
			blocks[i/3][j/3] |= d
		}
	}

	var dfs func(int, int) bool
	dfs = func(i int, cnt int) bool {
		if i == n {
			if cnt == 0 {
				return true
			}
			return false
		}

		// 像后面的优化， 使用 spaces 模拟， 可以省去一层 loop j   //好好体会这句话， 这是枚举的力量。
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := 1; k <= 9; k++ {
				if colsm[j]&(1<<(k-1)) != 0 {
					continue
				}
				if rowsm[i]&(1<<(k-1)) != 0 {
					continue
				}
				if blocks[i/3][j/3]&(1<<(k-1)) != 0 {
					continue
				}
				colsm[j] |= 1 << (k - 1)
				rowsm[i] |= 1 << (k - 1)
				blocks[i/3][j/3] |= 1 << (k - 1)
				//ans[i][j] = byte(k) + '0'
				board[i][j] = byte(k + '0')
				if dfs(i+1, cnt-1) {
					return true
				}
				board[i][j] = '.'
				colsm[j] ^= 1 << (k - 1)
				rowsm[i] ^= 1 << (k - 1)
				blocks[i/3][j/3] ^= 1 << (k - 1)
			}
		}
		return false
	}

	dfs(0, cnts)
	return
}
