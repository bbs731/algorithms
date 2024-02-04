package dp

/*
	f[i+1][j+1] = f[i][j+1] || f[i+1][j]
	f[j+1] = f[j+1] || f[j]
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	if n+m != len(s3) {
		return false
	}

	f := make([]bool, m+1)
	for j := 0; j <= m; j++ {
		f[j] = s2[:j] == s3[:j] //翻译的 dfs 中， if i < 0  return s2[:j+1] == s3[j+1]  这句话
	}

	for i := 0; i < n; i++ {
		f[0] = s1[:i+1] == s3[:i+1] //翻译的是 dfs 中  if j <0  return s1[:i+1] == s3[i+1] 的这句话
		for j := 0; j < m; j++ {
			var a, b bool
			if s1[i] == s3[i+j+1] {
				a = f[j+1]
			}
			if s2[j] == s3[i+j+1] {
				b = f[j]
			}
			f[j+1] = a || b
		}
	}
	return f[m]
}

/*
	f[i][j] =  f[i-1][j] || f[i][j-1]

	f[i+1][j+1] = f[i][j+1] || f[i+1][j]
 */
func isInterleave_dp(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	if n+m != len(s3) {
		return false
	}

	f := make([][]bool, n+1)
	// 应该这样写初始化
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
		f[i][0] = s1[:i] == s3[:i]
	}
	for j := 1; j <= m; j++ {
		f[0][j] = s2[:j] == s3[:j]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var a, b bool
			if s1[i] == s3[i+j+1] {
				a = f[i][j+1]
			}
			if s2[j] == s3[i+j+1] {
				b = f[i+1][j]
			}
			f[i+1][j+1] = a || b
		}
	}
	return f[n][m]
}

func isInterleave_dfs(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	if n+m != len(s3) {
		return false
	}
	var dfs func(int, int) bool

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	dfs = func(i, j int) bool {
		if i < 0 {
			return s2[:j+1] == s3[:j+1]
		}
		if j < 0 {
			return s1[:i+1] == s3[:i+1]
		}
		var ans bool
		if cache[i][j] != -1 {
			return cache[i][j] == 1
		}

		if s1[i] == s3[i+j+1] {
			ans = dfs(i-1, j)
		}
		if s2[j] == s3[i+j+1] {
			ans = ans || dfs(i, j-1)
		}
		if ans {
			cache[i][j] = 1
		} else {
			cache[i][j] = 0
		}
		return ans
	}
	return dfs(n-1, m-1)
}
