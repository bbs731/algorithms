package one_day_exercise

import (
	"math"
)

/****
这个 复杂度还是太高了， 灵神说的是：  O(m*n(m+n)) 具体，咋证明呢？ 不会啊！

如何优化呢？
 */
func minimumVisitedCells(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	h := make([][]int, m)
	v := make([][]int, m)
	cost := make([][]int, m)
	inqueue := make([][]bool, m)
	for i := range h {
		h[i] = make([]int, n)
		v[i] = make([]int, n)
		cost[i] = make([]int, n)
		inqueue[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h[i][j] = grid[i][j] + j
			v[i][j] = grid[i][j] + i
			cost[i][j] = math.MaxInt32
		}
	}

	cost[m-1][n-1] = 1
	type pair struct {
		i, j int
	}
	q := []pair{{m - 1, n - 1}}
	inqueue[m-1][n-1] = true

	for len(q) > 0 {
		t := q[0]
		x, y := t.i, t.j

		// 行， 向左
		for j := 0; j < y; j++ {
			if h[x][j] >= y && cost[x][j] > cost[x][y]+1 {
				cost[x][j] = cost[x][y] + 1
				if inqueue[x][j] == false {
					q = append(q, pair{x, j})
					inqueue[x][j] = true
				}
			}
		}

		// 列， 向上
		for i := 0; i < x; i++ {
			if v[i][y] >= x && cost[i][y] > cost[x][y]+1 {
				cost[i][y] = cost[x][y] + 1
				if inqueue[i][y] == false {
					q = append(q, pair{i, y})
					inqueue[i][y] = true
				}
			}
		}
		q = q[1:]
		inqueue[x][y] = false
	}

	//fmt.Println(h)
	//fmt.Println(v)
	//fmt.Println(cost)
	if cost[0][0] == math.MaxInt32 {
		return -1
	}
	return cost[0][0]
}
