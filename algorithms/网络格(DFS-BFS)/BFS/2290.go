package BFS

import "runtime/debug"

/*
**

题目和1293一样

灵神的视频讲解：
https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/solutions/1524710/0-1-bfs-by-endlesscheng-4pjt/

看完了灵神的视频 + 自己之前的理解 总结一下：
这道题，达到的共识就是求最短路： 从 [0,0] 到 [m-1, n-1] 边权都是正的：

讨论 3 中做法：
1. 我自己的写法。本质上应该是 Bellman-Ford 算法。 因为 grid[r][c] 可以重复入队。 所以复杂度不是 O(m*n)。
Bellman Ford 的算法复杂度是 O(VE) = O(m*n * m*n) = O(m^2*n^2)

2. 可以用 dijkstra 写， 用 heap 时间复杂度是 O（V+E）logV  并且 V = E= O(m*n)

3. 灵神给出的 0-1 BFS 代码，在下面， queue 选择的 是 deque， 如果是0 插入队头， 如果是 1 插入队尾。 这样得到的 BFS 求最短路
的时间复杂度可以优化到 O(V) = O(m*n) 比 dijkstra 还低
*/
func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	visited := make(map[pair]int)
	start := pair{0, 0}
	q := []pair{start}
	visited[start] = 0

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				if v, ok := visited[pair{r, c}]; !ok || v > visited[p]+grid[r][c] {
					// first time visited or can update
					visited[pair{r, c}] = visited[p] + grid[r][c]
					q = append(q, pair{r, c})
				}
			}
		}
	}
	return visited[pair{m - 1, n - 1}]
}

func minimumObstacles(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)

	type pair struct{ x, y int }
	var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m * n
		}
	}
	dis[0][0] = 0
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		for _, d := range dir4 {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				g := grid[x][y]
				if dis[p.x][p.y]+g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q[g] = append(q[g], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1]
}

/*
**
chunlei 的代码， 有 deque 优化一下， 提交看看用时上的差异。
*/

func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	visited := make(map[pair]int)
	start := pair{0, 0}
	q := [2][]pair{{}}
	visited[start] = 0
	q[0] = append(q[0], start)

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			//p, q[1] = q[1][len(q[1])-1], q[1][:len(q[1])-1]
			p, q[1] = q[1][0], q[1][1:] // 为什么写成上面那个，会有问题呢？
		}
		//for _, p :=range tmp {
		x, y := p.x, p.y
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			if v, ok := visited[pair{r, c}]; !ok || v > visited[p]+grid[r][c] {
				// first time visited or can update
				visited[pair{r, c}] = visited[p] + grid[r][c]
				g := grid[r][c]
				q[g] = append(q[g], pair{r, c})
				//if grid[r][c] == 0 {
				//	q[0] = append(q[0], pair{r, c})
				//} else {
				//	q[1] = append(q[1], pair{r, c})
				//}
			}
		}
		//}
	}
	return visited[pair{m - 1, n - 1}]
}



func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	// 用 distance 数组， 要比用 visited hash map 快了 一倍还多， 从 570ms 到 200 ms 这是为什么？
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m * n
		}
	}
	dis[0][0] = 0

	// 模拟 deque
	q := [2][]pair{{{0, 0}}}

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			//p, q[1] = q[1][len(q[1])-1], q[1][:len(q[1])-1] // 这个是错的，会造成循环。
			p, q[1] = q[1][0], q[1][1:] // 为什么写成上面那个，会有问题呢？
		}

		x, y := p.x, p.y
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			g := grid[r][c]
			if dis[r][c] > dis[x][y] + g {
				dis[r][c] = dis[x][y] + g
				q[g] = append(q[g], pair{r, c})
			}
		}
	}
	return dis[m - 1][ n - 1]
}



// 这个版本 305ms 所以慢的是 hashmap, 应该用 distance 数组。
// 即使是 Bellman Ford 也是足够的快！
func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}
	// 用 distance 数组， 要比用 visited hash map 快了 一倍还多， 从 570ms 到 200 ms 这是为什么？
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m * n
		}
	}
	dis[0][0] = 0
	start := pair{0, 0}
	q := []pair{start}

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				g := grid[r][c]
				if  dis[r][c] > dis[x][y]+g {
					dis[r][c] = dis[x][y] + g
					q = append(q, pair{r, c})
				}
			}
		}
	}
	return dis[m - 1][ n - 1]
}
