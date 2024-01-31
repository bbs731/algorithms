package BFS

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	q := []pair{}

	for i := range grid {
		q = append(q, pair{i, 0})
	}

	dirs := [][2]int{{1, 1}, {0, 1}, {-1, 1}}
	var step int
	for ; len(q) > 0; step++ {
		tmp := q
		q = nil

		// 不知道，真正面试的时候， instack 需不需要 follow-up 因为涉及到频繁的创建和丢弃，可以用local 数组优化吗？
		instack := make(map[pair]struct{})

		for _, p := range tmp {
			for _, d := range dirs {
				r, c := p.x+d[0], p.y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] <= grid[p.x][p.y] {
					continue
				}
				// otherwise 入队
				if _, ok := instack[pair{r, c}]; !ok {
					q = append(q, pair{r, c})
					instack[pair{r, c}] = struct{}{}
				}
			}
		}
	}

	return step - 1
}

// follow-up 减少空间复杂度， 用 one-dimension array 去重。

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	q := []pair{}

	clock := make([]int, m)

	for i := range grid {
		q = append(q, pair{i, 0})
	}

	dirs := [][2]int{{1, 1}, {0, 1}, {-1, 1}}
	var step int
	for step = 1; len(q) > 0; step++ {
		tmp := q
		q = nil

		// 不知道，真正面试的时候， instack 需不需要 follow-up 因为涉及到频繁的创建和丢弃，可以用local 数组优化吗？
		//instack := make(map[pair]struct{})

		for _, p := range tmp {
			for _, d := range dirs {
				r, c := p.x+d[0], p.y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] <= grid[p.x][p.y] {
					continue
				}
				// otherwise 入队
				//if _, ok := instack[pair{r, c}]; !ok {
				// 下一位 candidate 的 col 一定是 py+1, 所以可以用一个一维的数组来去重，只看 row 就可以了。 另外把 step 当做 clock 需要从 1 开始。
				if clock[r] != step {
					q = append(q, pair{r, c})
					//instack[pair{r, c}] = struct{}{}
					clock[r] = step
				}
			}
		}
	}

	return step - 2
}
