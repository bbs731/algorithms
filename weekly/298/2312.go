package weekly

/****
1。 变量
i,  j 长 和 高

2.
用长 i,  高 j 的木块， 最多切出来的价值。

3. dp[i][j] 是状态
如果是竖着切，  枚举 k from 1 to i-1  max( dp[k][j] + dp[i-k][j] )   = dp[i][j]    for all k
如果横着切，  枚举  dp[i][j] = max( dp[i][p] + dp[i][j-p])  for all p from 1 to j-1


4. 初始化 木块 横着和竖着的。
排满 m 或者  n, 感觉是不用这么复杂， 下面step 5 的枚举会处理每种情况。


5. loop i from 1 to m,  loop  j from 1 to n
 */

 /****
 这道题，要比，切 1443 pizza 的那道题简单啊！
  */
func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i]= make([]int, n+1)
	}

	for _, p := range prices {
		w, h, v := p[0], p[1], p[2]
		dp[w][h] = v
		if h <=m && w <=n {
			// 这道题， 不让旋转木块啊， 所以下面的不用处理。
			//dp[h][w] = v
		}
	}

	for i:=1; i<=m; i++ {
		for j :=1; j<=n; j++ {
			for k:=1; k<=j-1; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k] + dp[i][j-k])
			}
			for k:=1; k<=i-1; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j] + dp[i-k][j])
			}
		}
	}

	return int64(dp[m][n])
}
