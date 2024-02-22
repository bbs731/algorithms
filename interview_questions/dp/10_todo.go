package dp


/***

太棒了， 终于把 10 和 44 两道题都考虑清楚了！

f[i][j] =   f[i-1][j-1]   if s[i] == p[j] || p[j] == '.'
		=   f[i-1][j] ||   if p[j] == '*' and s[i] == p[j-1] || p[j] == '.'
		=   f[i][j-1]      if p[j] == '*'  这个选择一直都成立


// 灵神，给的答案！ 哈哈
func isMatch(s, p string) bool {
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}

 */

func isMatch(s string, p string) bool {
	n :=len(s)
	m := len(p)

	f := make([][]bool, n+1)
	for i:= range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	// 这个初始化， 太难了。
	for j:=0; j<m; j++ {
		if p[j] == '*' {
			f[0][j+1] = f[0][j-1]
		}
	}

	for i :=0; i<n; i++ {
		for j :=0; j<m; j++ {
			if p[j] == '*'{
				// 不用 (是有条件的， 条件就是 p[j-1] == '.' or p[j-1] == s[i])
				if p[j-1] == s[i] || p[j-1] == '.' {
					f[i+1][j+1] = f[i+1][j+1] || f[i][j+1]
				}
				// 用 , 这个选择一值都在。 因为用 0 次是合法的，不管 p[j-1] 是什么东西，都可以把 p[j-1] 去掉。
				f[i+1][j+1] = f[i+1][j+1] || f[i+1][j-1] // 用 0 次， 把 * 前面的 c 去掉

			}else {
				if p[j] == '.' || p[j] == s[i] {
					f[i+1][j+1] = f[i+1][j+1] || f[i][j]
				}
			}
		}
	}
	return f[n][m]
}
