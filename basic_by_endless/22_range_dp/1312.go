package dp

/*
	ToDo:  翻译成递推
 */
/*
	递归的做法一次过。

	dfs(i, j) defined as minnum steps to make string s[i..j] palindrome
	dfs(i,j) = min(dfs(i+1, j-1) if s[i] == s[j],
	              , dfs(i, j-1) + 1 // 添加左边，相当于删除右边,    dfs(i+1, j) + 1 // 添加右边，相当于删除左边 )

边界条件：
	dfs(i, j) = 0 if i >=j
 */
func minInsertions(s string) int {
	inf := int(1e9)
	n := len(s)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		ans := inf

		if s[i] == s[j] {
			ans = min(ans, dfs(i+1, j-1))
		}
		ans = min(ans, dfs(i, j-1)+1, dfs(i+1, j)+1)
		cache[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}
