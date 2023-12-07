package dp

/*
ToDo:
	把 dfs(i) 改成 递推！
*/

/*
dfs(i, j) =  min ( dfs(i, k) + dfs(k+1, j) )  + 1  for  i <=k <j
边界条件：
 dfs(i, j) = 0  if i >=j

这个计算 dfs(i, j)  [i, j] 这个范围内的最小割的时间复杂度是 O（n^3) 的。 这道题会超时，
我们最后想要的结果, 只关心的是 dfs(0, n-1) 的最小割， 换句话，只关心 i=0 ( dfs(0, ) 的， 不关心 i 是其它值的，能否简化时间复杂度？

定义  dfs(j) 为， 【0， j] 的最小割。 dfs(j)  有某个  dfs(k) 转换而来，k < j. 我们选择这个 k 让  s[k+1, j] 是 palindrome
dfs(j) = min( dfs(k) +1 +  0  {because of palindrome(k+1, j)})

这个优化相当于我们降了一个维度， 从 3维度，降到了2维， 因此时复杂度也就变成了 O（n^2)

所以，能够优化 DP 的时间复杂度，只能靠状态的变换，这道题的技巧就是 i 可以不求通用的，定死为0 dfs(0, j) -> dfs(j)

 */
func minCut(s string) int {
	n := len(s)

	// O(n^2) 的时间复杂度，来预处理 palidrome matrix 是对的。
	p := make([][]bool, n)
	for i := n - 1; i >= 0; i-- {
		p[i] = make([]bool, n+1)
		p[i][i+1] = true
		if i+1 < n {
			p[i+1][i+1] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && p[i+1][j] {
				p[i][j+1] = true
			}
		}
	}

	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if p[0][i+1] {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}

		ans := n
		for k := 0; k < i; k++ { // k 的 index 需要从 0开始， 要不然处理不了 "ab" 这种情况。
			if p[k+1][i+1] { // index, index ,index 判断  [k+1, i] 是 palindrome
				ans = min(ans, dfs(k)+1)
			}
		}
		cache[i] = ans
		return ans
	}
	return dfs(n - 1)
}

/*
	dfs(i, j) is  palindrome
	dfs(i,j) =   true if s[i] == s[j] && dfs(i+1, j-1) == true
初始化条件：
	dfs(i, i) = true  dfs(i+1，i）= true

	f[i][j+1] =  true if s[i] == s[j] && f[i+1[[j] == true   i 倒序遍历，j 正序遍历。
初始化：
	f[i][i] = true  f[i+1][i] = true
	变化成
	f[i][i+1] = true f[i+1][i+1] = true  // 仅为 j -> j + 1

/*

dfs(i, j) =  min ( dfs(i, k) + dfs(k+1, j) )  + 1  for  i <=k <j
边界条件：
 dfs(i, j) = 0  if i >=j
这个计算 dfs(i, j)  [i, j] 这个范围内的最小割的时间复杂度是 O（n^3) 的。

*/
func minCut(s string) int {
	n := len(s)

	// O(n^2) 的时间复杂度，来预处理 palidrome matrix 是对的。
	p := make([][]bool, n)
	for i := n - 1; i >= 0; i-- {
		p[i] = make([]bool, n+1)
		p[i][i+1] = true
		if i+1 < n {
			p[i+1][i+1] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && p[i+1][j] {
				p[i][j+1] = true
			}
		}
	}

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		if p[i][j+1] {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		ans := n
		for k := i; k < j; k++ {
			ans = min(ans, dfs(i, k)+dfs(k+1, j)+1)
		}
		cache[i][j] = ans
		return ans
	}

	return dfs(0, n-1)
}
