package dp

/*
	dfs(i, j)  i start position from word1 [0.. n1-1],   j end position from word2 [0..n2-1]
	定义为， 能总成的最长 palindrome 子串的最大长度。

	dfs(i, j) = max { dfs(i+1,j-1) + 2 //if word1[i] == word2[j],  dfs(i+1,j-1) // 都不选， dfs(i+1, j), dfs(i, j-1) }


边界条件：
	i > n1  return 0 or j+1  word2[j] is palindrome
	j < 0 return 0 or n1- i + 1 word1[i:] is palindrome

	return  dfs(0, n2-1)

 */

func longestPalindrome(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	inf := int(1e9)
	isP2 := make([]bool, n2)
	isP1 := make([]bool, n1)

	var dfs func(int, int) int

	// 预处理一下 palindrome

	dfs = func(i, j int) int {
		if i >= n1 {
			if isP2[j] {
				return j + 1
			}
			return 0
		}

		if j < 0 {
			if isP1[i] {
				return n1 - i + 1
			}
			return 0
		}

		ans := -inf
		ans = max(ans, dfs(i+1, j-1), dfs(i+1, j), dfs(i, j-1))
		if word1[i] == word2[j] {
			ans = max(ans, dfs(i+1, j-1)+2)
		}
		return ans
	}

}
