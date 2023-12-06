package dp

/*

	f[i+2][1] = max(f[i+1][1], f[i][0] -prices[i])
	f[i+2][0] = max(f[i+1][0], f[i+1][1] + prices[i])
初始化： f[1][0]= 0, f[1][1] = -inf

能降维吗？ 可以。 这次的降维干的漂亮啊， 一次救过了！

	f[1] = max(f[1], f[0]_special - prices[i])
	f[0] = max(f[0], f[1] + prices[i])

	f[0]_special, f[0], f[1]

初始化: f[0] = 0, f[1] = -inf
 */

func maxProfit(prices []int) int {
	inf := int(1e10)
	n := len(prices)
	var f0, f0_special, f1 int
	f1 = -inf
	for i := range prices {
		//f[i+2][1] = max(f[i+1][1], f[i][0]-prices[i])
		//f[i+2][0] = max(f[i+1][0], f[i+1][1]+prices[i])
		f1 = max(f1, f0_special-prices[i])
		old_f0 := f0               // 记录的是 f[i+1][0]
		f0 = max(f0, f1+prices[i]) // 计算 f[i+2][0]
		if i > 1 {
			f0_special = old_f0 // 复制 f0_special 为 f[i+1][0], 到下一个循环， f[i+3][1] 的时候用到的是 f0_special f[i+1][0]
		}
	}
	return f0
}

/*

参考 122 的状态转移方程

		f[i+1][1] = max(f[i][1], f[i-1][0] - price[i])  // 冷静期一天， 反应在状态方程上
		f[i+1][0] = max(f[i][0], f[i][1] + prices[i])
初始化： f[0][0] = 0 , f[0][1] = -inf


f[i-1] 当 i = 0 时候， index 会越界。 那么我们需要把 所有 f 的 index 再加 1
index i 变换：(这是本领，多联系）

	f[i+2][1] = max(f[i+1][1], f[i][0] -prices[i])
	f[i+2][0] = max(f[i+1][0], f[i+1][1] + prices[i])
初始化： f[1][0]= 0, f[1][1] = -inf

*/

func maxProfit_dp(prices []int) int {
	inf := int(1e10)
	n := len(prices)
	f := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		f[i] = make([]int, 2)
	}
	f[1][1] = -inf

	for i := range prices {
		f[i+2][1] = max(f[i+1][1], f[i][0]-prices[i])
		f[i+2][0] = max(f[i+1][0], f[i+1][1]+prices[i])
	}
	return f[n+1][0]
}
