package DFS


/***
这方法，真蠢啊！  好的方法是， 反向来找， 从 atlantic 的边界开始找， 能到达所有 atlantic 的点。
从 pacific 的边界开始找，所有能达到 pacific 的点。  然后并起来。
 */

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m := len(heights)
	n := len(heights[0])

	//var visited [][]bool
	var dfs func(int, int, [][]bool)
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	//visited = make([][]bool, m)  // 可以简化掉 visited 数组， pacific 和 atlantic 数组可以代替 visited 的功能。
	for i := range heights {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
		//visited[i] = make([]bool, n)
	}

	// 反过来 flood fill
	dfs = func (x, y int, sea [][]bool ) {
		//if visited[x][y] {
		//	return
		//}
		if sea[x][y] {
			return
		}

		//visited[x][y]=true
		sea[x][y] = true
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			r, c := x+d[0], y+ d[1]
			if r <0 || r>=m || c<0 || c >=n {
				continue
			}
			if heights[r][c] >= heights[x][y]{
				dfs(r, c, sea)
			}
		}
	}
	for j:=0; j <n; j++ {
		dfs(0, j, pacific)
	}
	for i:=0; i<m; i++ {
		dfs(i, 0, pacific)
	}
	// reset visited
	//for i:=0; i<m; i++ {
	//	for j:=0; j<n; j++ {
	//		visited[i][j]= false
	//	}
	//}

	for j:=0;j<n;j++ {
		dfs(m-1,j, atlantic)
	}
	for i:=0; i<m; i++ {
		dfs(i, n-1, atlantic)
	}

	for i:=0;i<m; i++ {
		for j:=0; j<n; j++ {
			if atlantic[i][j] && pacific[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}


func pacificAtlantic(heights [][]int) (ans [][]int) {
	m := len(heights)
	n := len(heights[0])

	var visited [][]bool
	var dfs func(int, int)
	var left_up, right_down bool
	dfs = func(x, y int) {
		if visited[x][y] {
			return
		}
		// mark [x, y] as visited
		visited[x][y] = true
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || c < 0 {
				left_up = true
			}
			if r >= m || c >= n {
				right_down = true
			}
			if r<0 || r >=m || c<0 || c >=n {
			//if left_up || right_down {  // 这里想取巧，结果错了哎！
				continue
			}

			if heights[r][c] <= heights[x][y] {
				dfs(r, c)
			}
		}
	}

	for i := range heights {
		for j := range heights[0]{
			// reset visited
			visited = make([][]bool, m)
			for k :=range visited {
				visited[k] = make([]bool, n)
			}
			left_up = false
			right_down = false
			dfs(i, j)
			if left_up && right_down {
				ans = append(ans, []int{i,j})
			}
		}
	}
	return
}
