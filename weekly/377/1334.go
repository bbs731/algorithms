package weekly

/*

Floyd 算法的模版题目

https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/

灵神讲解的 Floyd-Wallshall 从递归到递推的推导求解过程。 太强了！

 */


func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	inf := int(1e9)
	w := make([][]int, n)
	for i := range w{
		w[i] = make([]int, n)
		for j := range w[i]{
			w[i][j] = inf
		}
	}
	for _, e :=range edges {
		x, y , wt := e[0], e[1], e[2]
		w[x][y] = wt
		w[y][x] = wt
	}

	// Flody-Washall 最后就变成了三重循环这个最精简的版本， k 一定要在loop的最外层
	f := w
	for  k :=0; k<n; k++ {
		for i:=0; i<n; i++ {
			for j :=0; j <n; j++ {
				f[i][j] = min(f[i][j], f[i][k] + f[k][j])
			}
		}
	}

	ans := 0
	ll := n

	for i:=0; i<n; i++ {
		cnt := 0
		for j :=0; j<n; j++ {
			if i !=j && f[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ll {
			ll = cnt
			ans = i
		}
	}
	return ans
}



func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	inf := int(1e9)
	w := make([][]int, n)
	for i := range w{
		w[i] = make([]int, n)
		for j := range w[i]{
			w[i][j] = inf
		}
	}
	for _, e :=range edges {
		x, y , wt := e[0], e[1], e[2]
		w[x][y] = wt
		w[y][x] = wt
	}
	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[0] {
			cache[i][j] = make([]int, n)
			for k :=0;k<n; k++ {
				cache[i][j][k]= -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(k , i, j int) int {
		if k < 0 {
			return w[i][j]
		}
		if cache[k][i][j] != -1 {
			return cache[k][i][j]
		}
		// 不选 k
		res := dfs(k-1, i, j)
		//选 k
		res = min(res, dfs(k-1, i, k) + dfs(k-1, k, j))
		//res = min(dfs(k-1, i,j), dfs(k-1, i, k)+ dfs(k-1, k, j))  // 也可以写成一行，更加的简洁
		cache[k][i][j] = res
		return res
	}


	ans := 0
	ll := n

	for i:=0; i<n; i++ {
		cnt := 0
		for j :=0; j<n; j++ {
			//if i !=j && f[i][j] <= distanceThreshold {
			if i !=j && dfs(n-1, i,j) <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ll {
			ll = cnt
			ans = i
		}
	}
	return ans
}