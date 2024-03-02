package dp


func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	// 这个初始化太难了！
	for i:=0; i<m; i++ {
		if p[i] == '*' {
			f[0][i+1] = f[0][i-1]
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if p[j] == '*' {
				f[i+1][j+1]= f[i+1][j-1]  // 一定可以满足的。
				if s[i] == p[j-1] || p[j-1] == '.' {
					f[i+1][j+1] = f[i+1][j-1] || f[i][j+1]
				}
			} else {
				if p[j] == '.' || s[i] == p[j] {
					f[i+1][j+1] = f[i][j]
				}
			}
		}
	}
	return f[n][m]
}
