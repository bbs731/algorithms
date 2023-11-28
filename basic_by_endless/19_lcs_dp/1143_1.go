package dp

/*
 	f[i+1][j+1] =   1. f[i][j] + 1
					2. f[i][j+1] , f[i+1][j]

如果降维到 1维的数组的话，f[i][j] 和 f[i+1][j] 这两个值，在1维数组里，表示的是同一个数字，没办法区分， 所以需要额外的一个变量来保存。

难！
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	f := make([]int, m+1)

	for i := 0; i < n; i++ {
		prev := f[0] // = 0  // 这里多用一个 prev 变量来保存  f[i][j] 的值, 降维之后就是 f[j] 的值。
		for j := 0; j < m; j++ {
			tmp := f[j+1] //这是下个循环的 f[j] 因此需要保存，并且最后需要赋值给 prev 变量，以便下个循环的时候使用。
			if text1[i] == text2[j] {
				//f[j+1] = f[j] + 1
				f[j+1] = prev + 1
			} else {
				f[j+1] = max(f[j+1], f[j])
			}
			prev = tmp
		}
	}
	return f[m]
}

/*
 降维到2个一维的数组：
	f[2][m+1]   牛啊， 把所有的第一维度的  i+1 和 i  都 %2， 结果 n 也需要 %2 即可。
	这个都是标准的套路。
 */
func longestCommonSubsequence_2_dim(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	f := make([][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[(i+1)%2][j+1] = f[i%2][j] + 1
			} else {
				f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j])
			}
		}
	}
	return f[n%2][m]
}

/*
 	f[i][j] = max(f[i-1][j] , f[i][j-1]) + (text1[i] == text2[j])
	f[i+1][j+1] = max(f[i][j+1], f[i+1][j]) + (text1[i] == text2[j])
 */
func longestCommonSubsequence_dp_basic(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

func longestCommonSubsequence_dfs(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		ans := max(dfs(i-1, j), dfs(i, j-1))
		cache[i][j] = ans
		return ans
	}

	return dfs(n-1, m-1)
}
