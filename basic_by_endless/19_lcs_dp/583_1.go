package dp

/*
	f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1   or f[i][j]       f[i+1][j] 和  f[i][j] 变成一维数组之后，会无法区分，所以需要个额外的变量。
	于是我们引入  prevj = f[i][j]
降维：
	f[j+1] = min(f[j+1], f[j]) + 1 or prevj = f[i][j] = f[j]

这道题的降维， 应该是最容易的降维了。第一遍还是错了。

 */
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	f := make([]int, m+1)

	for j := 1; j <= m; j++ {
		f[j] = j
	}

	for i := 0; i < n; i++ {
		prevj := f[0] // prevj 是用来保存 f[i][j] 的， 初始化为  f[0] 很自然的事情。
		f[0] = i + 1  // 翻译的  f[i+1][0] = i + 1 这句话
		for j := 0; j < m; j++ {
			tmp := f[j+1]
			if word1[i] == word2[j] {
				f[j+1] = prevj // 翻译的这句话： f[i+1][j+1] = f[i][j]
			} else {
				f[j+1] = min(f[j+1], f[j]) + 1 // 翻译的这句话： 	f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
			prevj = tmp
		}
	}
	return f[m]
}

/*
	f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
 */
func minDistance_dp(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}

	for j := 1; j <= m; j++ {
		f[0][j] = j // 翻译的  if i < 0  return j+1 这句话
	}

	for i := 0; i < n; i++ {
		f[i+1][0] = i + 1 // 翻译的 dfs 中  if j < 0 return  i+1 这句话
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}

	return f[n][m]
}

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

		ans := min(dfs(i-1, j), dfs(i, j-1)) + 1
		cache[i][j] = ans
		return ans
	}
	return dfs(n-1, m-1)
}
