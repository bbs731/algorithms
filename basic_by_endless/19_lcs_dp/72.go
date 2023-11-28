package dp

/*

	f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1

降维去掉 i 的维度，那么， f[i+1][j] 和 f[i][j] 会坍缩为 f[j]，所以需要一个额外的变量来保存 f[i][j] 记为 prevj
	f[j+1] = min(f[j+1], f[j], prevj) + 1


这个降维比 1143还难， 做不对！
 */
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	f := make([]int, m+1)

	// 用边界条件，来初始化数组
	for j := 1; j <= m; j++ {
		f[j] = j
	}

	for i := 0; i < n; i++ {
		// 这里用 f[0] 或者  i  来初始化  prevj 都是对的。
		prevj := f[0] // f[0] 来初始化， 根据上面"条件边界，来初始化数组"的内容来看  f[i][0] = i 所以初始化 prevj = i
		f[0]++        //这句话代表什么？ 没降维之前的DP代码中，边界条件 f[i+1][0] = i+1 初始化，
		//这就提现的降维之后的复杂性， f[0] 代表了多重意思， 还要反应之前 f[i+1][0] 的变化情况， 这里也可以写成  f[0] = i + 1,
		//"单独计算 j=0 的情况，注意下面的循环相当于从 f[1] 开始计算"

		for j := 0; j < m; j++ {
			tmp := f[j+1]
			if word1[i] == word2[j] {
				f[j+1] = prevj
			} else {
				f[j+1] = min(f[j+1], f[j], prevj) + 1
			}
			prevj = tmp
		}
	}
	return f[m]
}

/*
	dfs(i, j) =  min(dfs(i-1,j), dfs(i,j-1), dfs(i-1, j-1)) + 1   else  dfs(i-1)(j-1) if word1[i] == word2[j]
翻译成 DP数组

	f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
	f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
 */
func minDistance_dp(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	f := make([][]int, n+1)

	// 用边界条件，来初始化数组
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}

	for i := 0; i < n; i++ {
		f[i+1][0] = i + 1
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}

/*
	i: word1
	j: word2
	dfs(i, j)

	we consider modification to word1 (the last letter)
	insertion word1 => dfs(i, j-1)  (添加 letter 到word1 结尾， 然后和word2 结尾一起删除）
	replace word1 => dfs(i-1, j-1)
	delete word1 = > dfs(i-1, j)

 */
func minDistance_dfs(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	var dfs func(int, int) int

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		}
		ans := min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
		cache[i][j] = ans
		return ans
	}
	return dfs(n-1, m-1)
}
