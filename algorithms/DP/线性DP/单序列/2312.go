package dp

/****
f[i][j] =  max ( f[i-a(k)][j] + f[a(k)][j],    f[i][j-b(k)] + f[i][b(k)])   for all posiisble (ak, bk) such that ak <=i and bk <=j

灵神的题解：
https://leetcode.cn/problems/selling-pieces-of-wood/solutions/1611240/by-endlesscheng-mrmd/

 */

func sellingWood(m int, n int, prices [][]int) int64 {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for _, p := range prices {
		a, b := p[0], p[1]
		f[a][b] = p[2]
	}

	// 这个是不对的，为啥？  这里面涉及到一个更新的顺序，下面的这种做法，可能保证不了最优的更新策略！
	// 因为，还有金额， 所以，以什么样的顺序更新，没办法确定。
	//for k := 0; k < len(prices); k++ {
	//	a, b := prices[k][0], prices[k][1]
	//	for i := prices[k][0]; i <= m; i++ {
	//		for j := prices[k][1]; j <= n; j++ {
	//			f[i][j] = max(f[i-a][j]+f[a][j], f[i][j-b]+f[i][b])
	//		}
	//	}
	//}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 0; k < m && i >= k; k++ {
				f[i][j] = max(f[i][j], f[i-k][j]+f[k][j])
			}
			for k := 0; k < n && j >= k; k++ {
				f[i][j] = max(f[i][j], f[i][j-k]+f[i][k])
			}
		}
	}

	return int64(f[m][n])
}
