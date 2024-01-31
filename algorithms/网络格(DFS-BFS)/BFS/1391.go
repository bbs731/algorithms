package BFS

/***

//You are given an m x n grid. Each cell of grid represents a street. The street
// of grid[i][j] can be:
//
//
// 1 which means a street connecting the left cell and the right cell.
// 2 which means a street connecting the upper cell and the lower cell.
// 3 which means a street connecting the left cell and the lower cell.
// 4 which means a street connecting the right cell and the lower cell.
// 5 which means a street connecting the left cell and the upper cell.
// 6 which means a street connecting the right cell and the upper cell.
//
//
// You will initially start at the street of the upper-left cell (0, 0). A valid
// path in the grid is a path that starts from the upper left cell (0, 0) and ends
// at the bottom-right cell (m - 1, n - 1). The path should only follow the street
//s.
//
// Notice that you are not allowed to change any street.
//
// Return true if there is a valid path in the grid or false otherwise.
//
//
// Example 1:
//
//
//Input: grid = [[2,4,3],[6,5,2]]
//Output: true
//Explanation: As shown you can start at cell (0, 0) and visit all the cells of
//the grid to reach (m - 1, n - 1).
//
//

 */

/***

那个形状，在那个位置，可以对接那些形状， 这个条件太长了。 （还好一遍写对了）
但是， 有没有啥技巧？ 简洁的写法？

 */

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	//visited :=  // mark to 0 avoid visited 数组怎么样？

	type pair struct {
		x, y int
	}
	q := []pair{{0, 0}}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			var r, c int
			x, y := p.x, p.y
			//candidates := []pair{}
			switch grid[x][y] {
			case 1: //left and right
				r, c = x, y-1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 4 || grid[r][c] == 6) {
					q = append(q, pair{r, c})
				}
				r, c = x, y+1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 3 || grid[r][c] == 5) {
					q = append(q, pair{r, c})
				}
			case 2: //upper and lower
				r, c = x+1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 6 || grid[r][c] == 5) {
					q = append(q, pair{r, c})
				}
				r, c = x-1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 3 || grid[r][c] == 4) {
					q = append(q, pair{r, c})
				}

			case 3: //left and lower
				r, c = x, y-1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 6 || grid[r][c] == 4) {
					q = append(q, pair{r, c})
				}
				r, c = x+1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 5 || grid[r][c] == 6) {
					q = append(q, pair{r, c})
				}
			case 4: //right and lower
				r, c = x, y+1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 3 || grid[r][c] == 5) {
					q = append(q, pair{r, c})
				}
				r, c = x+1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 5 || grid[r][c] == 6) {
					q = append(q, pair{r, c})
				}
			case 5: // left and upper
				r, c = x, y-1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 4 || grid[r][c] == 6) {
					q = append(q, pair{r, c})
				}
				r, c = x-1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 3 || grid[r][c] == 4) {
					q = append(q, pair{r, c})
				}
			case 6: //right and upper
				r, c = x, y+1
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == 3 || grid[r][c] == 5) {
					q = append(q, pair{r, c})
				}
				r, c = x-1, y
				if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 2 || grid[r][c] == 3 || grid[r][c] == 4) {
					q = append(q, pair{r, c})
				}
			}

			//mark p to 0 as visited.
			grid[x][y] = 0

			//for _, c := range q {
			//	//if c.x < 0 || c.x >= m || c.y < 0 || c.y >= n {
			//	//	continue
			//	//}
			//	if grid[c.x][c.y] != 0 {
			//		q = append(q, pair{c.x, c.y})
			//	}
			//}
		}
	}

	if grid[m-1][n-1] == 0 {
		return true
	}
	return false
}
