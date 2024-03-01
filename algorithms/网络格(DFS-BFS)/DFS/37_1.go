package DFS

func solveSudoku(board [][]byte) {
	n := len(board)

	colsm := make([]int, n)
	rowsm := make([]int, n)
	blocks := [3][3]int{} //这个是个关键的优化，不容易想到
	var spaces [][2]int   //这个是一个非常高效的枚举。
	cnts := 0

	for i := range board {

		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				cnts++
				spaces = append(spaces, [2]int{i, j})
				continue
			}
			d := 1 << (board[i][j] - '1')

			colsm[j] |= d
			rowsm[i] |= d
			blocks[i/3][j/3] |= d

		}
	}

	var dfs func(int, int) bool
	dfs = func(pos int, cnt int) bool {
		// 其实，这个 count 是多余的， 我们直接考虑  pos 就可以了。
		if cnt == 0 {
			return true
		}
		if pos == len(spaces) {
			return false
		}
		i, j := spaces[pos][0], spaces[pos][1]
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
			if dfs(pos+1, cnt-1) {
				return true
			}
			board[i][j] = '.'
			colsm[j] ^= 1 << (k - 1)
			rowsm[i] ^= 1 << (k - 1)
			blocks[i/3][j/3] ^= 1 << (k - 1)
		}
		return false
	}

	dfs(0, cnts)
	return
}
