package dp

/*

能降到 1维度吗？  不知道。 反正是没做出来！

为什么不能再降维了？ 是顺序有冲突了吗？
第一次降维已经要求， j 需要降序遍历，从 k-1 到 0， 但是
f[j+1][1] = max(f[j+1][1] , f[j][0] -prices[i])  是需要需要 正序遍历 j 的， 先算出  f[j] 再算 f[j+1]。 这里产生了矛盾， 导致
不能降维一维数组吗？ （只是猜想而已， 咋证明）

下面的答案是错的。
		f[j+1][1] = max(f[j+1][1] , f[j][0] -prices[i])
		f[j+1][0] = max(f[j+1][0], f[j+1][1] + prices[i])
初始化为 f[0..k][0] = 0  其它都为 -inf

		f[1] = max(f[1], prev0 - prices[i])
		f[0] = max(f[0], f[1] + prices[i])
f0 = 0 prev0 =0, f1 = -inf
 */

// 错误的答案！
func maxProfit(k int, prices []int) int {
	//inf := int(1e10)
	//var f1, f0 int
	//prev0 := make([]int, k)
	//n := len(prices)
	//f1 = -inf
	//for i := 0; i < n; i++ {
	//	for j := k - 1; j >= 0; j-- {
	//		f0 = max(f0, f1+prices[i])
	//		f1 = max(f1, prev0[j]-prices[i])
	//		prev0[j] = f0
	//	}
	//}
	//return f0
}

/*

		f[i+1][j+1][1] = max(f[i][j+1][1], f[i][j][0] - prices[i])
		f[i+1][j+1][0] = max(f[i][j+1][0], f[i][j+1][1] + prices[i])
初始化： j = 0  f[][0][] = -inf   and  i == 0   f[0][][1] = -inf  f[0][][0]= 0
特殊的边界条件， f[0][0][0] 应该给 0

降 1 个维度： 这个降维漂亮， 一次过了。 能降到 1维度吗？
		f[j+1][1] = max(f[j+1][1] , f[j][0] -prices[i])
		f[j+1][0] = max(f[j+1][0], f[j+1][1] + prices[i])

初始化为 f[0..k][0] = 0  其它都为 -inf
 */

func maxProfit(k int, prices []int) int {
	inf := int(1e10)
	n := len(prices)
	f := make([][2]int, k+1)
	// 初始化
	for j := range f {
		f[j][0] = 0
		f[j][1] = -inf
	}

	for i := 0; i < n; i++ {
		for j := k - 1; j >= 0; j-- {
			f[j+1][0] = max(f[j+1][0], f[j+1][1]+prices[i])
			f[j+1][1] = max(f[j+1][1], f[j][0]-prices[i])
		}
	}
	return f[k][0]
}

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
