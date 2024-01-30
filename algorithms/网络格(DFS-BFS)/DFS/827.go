package DFS


/***
dfs flood fill 的时候， 把一个 connected_comments 周边的 0 的位置， 譬如 grid[p][q] = 0  adj[p][q] += num_of_nodes(in this scc)
最后我们统计，那个 adj 里的位置的 counter 值最大，即为所求。
 */
func largestIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	adj := make([][]int, m)
	for i:= range visited {
		visited[i] = make([]bool, n)
		adj[i] = make([]int, n)
	}
	type pair struct {
		x, y int
	}
	var dfs func(int, int)
	var cnts int
	allzero := true
	components := make(map[pair]struct{})

	dfs = func(x, y int) {
		if visited[x][y] {
			return
		}
		visited[x][y]=true
		cnts++
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d:=range dirs {
			r, c := x+ d[0], y + d[1]
			if r <0 || r>=m || c<0 || c>=n {
				continue
			}
			if grid[r][c] == 0 {
				components[pair{r,c}] = struct{}{}
			} else {
				dfs(r,c)
			}
		}
	}
	ans := 0
	for i:= range grid{
		for j := range grid[0]{
			if visited[i][j] == false && grid[i][j] == 1 {
				//components = make([]pair, 0)
				cnts = 0
				dfs(i,j)
				for  c:= range components {
					adj[c.x][c.y] += cnts
					ans = max(ans, adj[c.x][c.y])
				}
				if cnts > 0 {
					allzero = false
				}
				components = make(map[pair]struct{})
			}
		}
	}

	// 这里的特判有点多啊！ 感觉，肯定还有更简明的逻辑。
	if ans == 0 {
		if allzero {
			return 1
		}
		return m*n
	}
	return min(ans+1, m*n)
}
