package dp


/***

f[i][j] = f[i-1][j-1]  if s1[i] == s2[j]
		=  min( f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1     //拉了一个 f[i-1][j-1] 的替换操作 对于最后一个 字符。

 */

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	inf := int(1e9)

	f := make([][]int, n+1)
	for i:=range f{
		f[i] = make([]int, m+1)
		for j:= range f[i] {
			f[i][j] = inf
			if i== 0 {
				f[0][j] = j
			}
			if j == 0 {
				f[i][0] = i
			}
		}
}


	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if word1[i] == word2[j]{
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}


