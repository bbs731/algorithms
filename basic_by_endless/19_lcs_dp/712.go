package dp

/*

不错啊， 这次降维的代码一次就过了。 找到 pattern 了。

 */
func minimumDeleteSum(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)

	f := make([]int, m+1)

	for j := 1; j <= m; j++ {
		f[j] = stringCost(s2[:j])
	}

	for i := 0; i < n; i++ {
		prevj := f[0]
		f[0] = stringCost(s1[:i+1])
		for j := 0; j < m; j++ {
			tmp := f[j+1] // 相当于翻译的这句话  tmp := f[i+1][j+1]
			if s1[i] == s2[j] {
				f[j+1] = prevj
			} else {
				f[j+1] = min(f[j+1]+int(s1[i]), f[j]+int(s2[j]))
			}
			prevj = tmp
		}
	}
	return f[m]
}

/*
    dfs(i, j) := min(dfs(i-1, j)+int(s1[i]), dfs(i, j-1)+int(s2[j]))
翻译。
	f[i+1][j+1] = min(f[i][j+1] + int(s1[i[) ,  f[i+1][j] + int(s2[j]))  if s1[i] != s2[j] else f[i][j]

 */
func minimumDeleteSum_dp(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}

	for j := 1; j <= m; j++ {
		f[0][j] = stringCost(s2[:j])
	}

	for i := 0; i < n; i++ {
		f[i+1][0] = stringCost(s1[:i+1])
		for j := 0; j < m; j++ {
			if s1[i] == s2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1]+int(s1[i]), f[i+1][j]+int(s2[j]))
			}
		}
	}

	return f[n][m]
}

func minimumDeleteSum_dfs(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return stringCost(s2[:j+1])
		}
		if j < 0 {
			return stringCost(s1[:i+1])
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if s1[i] == s2[j] {
			return dfs(i-1, j-1)
		}

		ans := min(dfs(i-1, j)+int(s1[i]), dfs(i, j-1)+int(s2[j]))
		cache[i][j] = ans
		return ans
	}
	return dfs(n-1, m-1)
}

func stringCost(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}
	return sum
}
