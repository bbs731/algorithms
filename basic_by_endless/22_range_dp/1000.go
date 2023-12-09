package dp

/*
https://leetcode.cn/problems/minimum-cost-to-merge-stones/solutions/2207235/tu-jie-qu-jian-dpzhuang-tai-she-ji-yu-yo-ppv0/
灵神的题解，太赞了！
 */
/*
dfs(i, j) = min{ dfs(i, m) + dfs(m+1, j) }  + sum[i..j]   for m = i+ (k-1)x  if (j-i) % (k-1) == 0
dfsi, j) = min { dfs(i,m) + dfs(m+1, j)}   for m = i+ (k-1)x  if (j-1) % (k-1) != 0
因此可以去掉一个维度。

翻译成，递推
	f[i][j] = min( f[i][m] + f[m+1][j] ) + sum[i..j] if (j-i) % (k-1) == 0
	i 需要倒序遍历， j 正序遍历。
初始化， f[i][i] = 0
return : f[i][n-1]


恩， 看到最后， 有啥感想？
为啥后面，有两种方式，做的答案是不对的？ 找时间想想为什么？

 */

func mergeStones(stones []int, k int) int {

	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = inf
			for m := i; m < j; m += k - 1 {
				f[i][j] = min(f[i][j], f[i][m]+f[m+1][j])

			}
			if (j-i)%(k-1) == 0 {
				f[i][j] += psum[j+1] - psum[i]
			}
		}
	}

	return f[0][n-1]
}

/*
dfs(i, j, 1) = dfs(i, j, k) + sum[i..j]  这个想的就比较难。
dfs(i,j,p) = min(dfs(i, m, 1) ... dfs(m+1, j, p-1))  其中  i< m = i+ (k-1)x  <j && p >= 2

边界条件：
 dfs(i, i, 1) = 0

return
dfs(0, n-1,1)

可以优化为：

dfs(i, j, 1) = min{ dfs(i, m, 1) + dfs(m+1, j, k-1) }  + sum[i..j]   for m = i+ (k-1)x  if (j-i) % (k-1) == 0
dfsi, j, 1) = min { dfs(i,m, 1) + dfs(m+1, j, k-1) }   for m = i+ (k-1)x  if (j-1) % (k-1) != 0
因此可以去掉一个维度。


 */
func mergeStones(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
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
		for m := i; m < j; m += k - 1 {
			ans = min(ans, dfs(i, m)+dfs(m+1, j))
		}
		if (j-i)%(k-1) == 0 {
			ans += psum[j+1] - psum[i]
		}
		cache[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}

/*

这道题是个地域级别， k=2 时，是一道经典的， 区间DP 问题。
把 k 参数化之后就是个地狱的难度。不过可以好好锻炼一下思维。


dfs(i, j, 1) = dfs(i, j, k) + sum[i..j]  这个想的就比较难。
dfs(i,j,p) = min(dfs(i, m, 1) ... dfs(m+1, j, p-1))  其中  i< m = i+ (k-1)x  <j && p >= 2

边界条件：
 dfs(i, i, 1) = 0

return
dfs(0, n-1,1)

 */

func mergeStones(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, k+1)
			for m := range cache[i][j] {
				cache[i][j][m] = -1
			}
		}
	}

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	var dfs func(int, int, int) int
	dfs = func(i, j int, p int) int {
		ans := inf
		if cache[i][j][p] != -1 {
			return cache[i][j][p]
		}
		if p == 1 {
			if i >= j {
				return 0
			}
			ans = dfs(i, j, k) + psum[j+1] - psum[i] // 这里是不是太反直觉了？
			cache[i][j][p] = ans
			return ans
		}

		for m := i; m < j; m += k - 1 {
			ans = min(ans, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}

		cache[i][j][p] = ans
		return ans
	}
	return dfs(0, n-1, 1)
}

// 这个答案是错的， 为什么呢？
func mergeStones_dfs2(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, k+1)
			for m := range cache[i][j] {
				cache[i][j][m] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j int, p int) int {
		ans := inf

		if cache[i][j][p] != -1 {
			return cache[i][j][p]
		}

		if p == 1 {
			if i >= j {
				return 0
			}
			ans = dfs(i, j, k) + psum[j+1] - psum[i] // 这里是不是太反直觉了？
			cache[i][j][p] = ans
			return ans
		}

		for m := i; m < j; m++ {
			var rank int
			if m-i+1 < p {
				rank = m - i + 1
			} else {
				rank = (m-i+1-p)%(p-1) + 1
			}
			ans = min(ans, dfs(i, m, rank)+dfs(m+1, j, p-rank))
		}
		cache[i][j][p] = ans
		return ans
	}
	return dfs(0, n-1, 1)
}

// 这代码也是太简洁了！
/*
dfs(i, j, 1) = dfs(i, j, k) + sum[i..j]  这个想的就比较难。
dfs(i,j,p) = min(dfs(i, m, 1) ... dfs(m+1, j, p-1))  其中  i< m = i+ (k-1)x  <j && p >= 2

边界条件：
 dfs(i, i, 1) = 0

return
dfs(0, n-1,1)
 */
func mergeStones_dfs(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	var dfs func(int, int, int) int
	dfs = func(i, j int, p int) int {
		if p == 1 {
			if i >= j {
				return 0
			}
			return dfs(i, j, k) + psum[j+1] - psum[i] // 这里是不是太反直觉了？
		}

		ans := inf
		for m := i; m < j; m += k - 1 {
			ans = min(ans, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}

		return ans
	}
	return dfs(0, n-1, 1)
}

/*
错误的答案：

一会儿去掉 p 的维度之后，对比一下，看看到底错误了哪里？
dfs(i, j) 合并 从 [i..j] stones 需要最小的 cost, 那么 dfs(i, j) 是怎么来的呢？
由， dfs(i, k) 和 dfs(k+1，j) 转换而来。
dfs(i, j) = min( dfs(i, k) +dfs(k+1, j) ) + sum[i..j]

 */

func mergeStones_wrong(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		if j-i+1 <= k {
			return psum[j+1] - psum[i]
		}

		ans := inf
		for q := i + 1; q < j; q++ {
			ans = min(ans, dfs(i, q)+dfs(q+1, j))
		}
		ans += psum[j+1] - psum[i]

		return ans
	}
	return dfs(0, n-1)
}
