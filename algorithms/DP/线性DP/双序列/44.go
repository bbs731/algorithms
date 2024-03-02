package dp


func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	// 这里是永远的痛， 永远也想不明白！
	for i:=0; i<m; i++ {
		if p[i] == '*'{
			f[0][i+1] = true
		} else {
			break
		}
	}

	for i:=0; i<n; i++ {
		for j :=0; j<m; j++ {
			if s[i] == p[j] || p[j] == '?' {
				f[i+1][j+1] = f[i][j]
			} else if p[j] == '*' {
				f[i+1][j+1] = f[i+1][j] || f[i][j+1]
			}
		}
	}
	return f[n][m]
}
