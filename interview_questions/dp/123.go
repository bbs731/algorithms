package dp

/****
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1:

输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

 */

func maxProfit(prices []int) int {
	n := len(prices)
	f := make([][][]int, n+1)
	//for i := range f {
	//	f[i] = make([][2]int, 3)
	//}
	inf := int(1e10)
	k := 2

	// 这里的， 初始化才是最难得!
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			f[i][j] = make([]int, 2)
			f[i][j][1] = -inf // 这个初始化太难了！
		}
	}
	for j := 0; j < k+1; j++ {
		f[0][k][0] = 0
	}

	for i := 0; i < n; i++ {
		for k := 0; k < 2; k++ {
			f[i+1][k+1][0] = max(f[i][k+1][1]+prices[i], f[i][k+1][0])
			f[i+1][k+1][1] = max(f[i][k+1][1], f[i][k][0]-prices[i])
		}
	}

	return f[n][k][0]

}