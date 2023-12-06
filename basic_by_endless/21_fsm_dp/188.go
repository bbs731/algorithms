package dp

/*

		f[i][j][1] = max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		f[i][j][0] =  max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])

i , j = > i+1, j+1
		f[i+1][j+1][1] = max(f[i][j+1][1], f[i][j][0] - prices[i])
		f[i+1][j+1][0] = max(f[i][j+1][0], f[i][j+1][1] + prices[i])
初始化： j = 0  f[][0][] = -inf   and  i == 0   f[0][][1] = -inf  f[0][][0]= 0
特殊的边界条件， f[0][0][0] 应该给 0

 */

func maxProfit(k int, prices []int) int {
	n := len(prices)
	inf := int(1e10)
	f := make([][][]int, n+1)

	// 这里的， 初始化才是最难得!
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			f[i][j] = make([]int, 2)
			f[i][j][1] = -inf
		}
	}
	for j := 0; j < k+1; j++ {
		f[0][k][0] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			f[i+1][j+1][1] = max(f[i][j+1][1], f[i][j][0]-prices[i])
			f[i+1][j+1][0] = max(f[i][j+1][0], f[i][j+1][1]+prices[i])
		}
	}

	return f[n][k][0]
}

func maxProfit_dfs(k int, prices []int) int {
	n := len(prices)
	inf := int(1e10)

	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			cache[i][j] = make([]int, 2)
			cache[i][j][0] = -1
			cache[i][j][1] = -1
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i int, j int, hold int) int {
		// 先判断  j < 0  然后再判断  i < 0  这他娘的也是一个坑！
		if j < 0 {
			return -inf
		}

		if i < 0 {
			if hold == 1 {
				return -inf
			}
			return 0
		}

		if cache[i][j][hold] != -1 {
			return cache[i][j][hold]
		}

		var ans int

		if hold == 1 {
			ans = max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		} else {
			ans = max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i]) // 因为 k 是买卖的测试， 所以要买买的时候 -1 ，要么卖的时候 -1 不能买卖的时候都去减，本解选择买的时候 -1
		}
		cache[i][j][hold] = ans
		return ans
	}
	return dfs(n-1, k, 0)
}
