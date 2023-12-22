package weekly

/*

假设你是一家合金制造公司的老板，你的公司使用多种金属来制造合金。现在共有 n 种不同类型的金属可以使用，并且你可以使用 k 台机器来制造合金。每台机器都需要特定数量的每种金属来创建合金。

对于第 i 台机器而言，创建合金需要 composition[i][j] 份 j 类型金属。最初，你拥有 stock[i] 份 i 类型金属，而每购入一份 i 类型金属需要花费 cost[i] 的金钱。

给你整数 n、k、budget，下标从 1 开始的二维数组 composition，两个下标从 1 开始的数组 stock 和 cost，请你在预算不超过 budget 金钱的前提下，最大化 公司制造合金的数量。

所有合金都需要由同一台机器制造。

返回公司可以制造的最大合金数。



示例 1：

输入：n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,0], cost = [1,2,3]
输出：2
解释：最优的方法是使用第 1 台机器来制造合金。
要想制造 2 份合金，我们需要购买：
- 2 份第 1 类金属。
- 2 份第 2 类金属。
- 2 份第 3 类金属。
总共需要 2 * 1 + 2 * 2 + 2 * 3 = 12 的金钱，小于等于预算 15 。
注意，我们最开始时候没有任何一类金属，所以必须买齐所有需要的金属。
可以证明在示例条件下最多可以制造 2 份合金。
示例 2：

输入：n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,100], cost = [1,2,3]
输出：5
解释：最优的方法是使用第 2 台机器来制造合金。
要想制造 5 份合金，我们需要购买：
- 5 份第 1 类金属。
- 5 份第 2 类金属。
- 0 份第 3 类金属。
总共需要 5 * 1 + 5 * 2 + 0 * 3 = 15 的金钱，小于等于预算 15 。
可以证明在示例条件下最多可以制造 5 份合金。
示例 3：

输入：n = 2, k = 3, budget = 10, composition = [[2,1],[1,2],[1,1]], stock = [1,1], cost = [5,5]
输出：2
解释：最优的方法是使用第 3 台机器来制造合金。
要想制造 2 份合金，我们需要购买：
- 1 份第 1 类金属。
- 1 份第 2 类金属。
总共需要 1 * 5 + 1 * 5 = 10 的金钱，小于等于预算 10 。
可以证明在示例条件下最多可以制造 2 份合金。


提示：

1 <= n, k <= 100
0 <= budget <= 108
composition.length == k
composition[i].length == n
1 <= composition[i][j] <= 100
stock.length == cost.length == n
0 <= stock[i] <= 108
1 <= cost[i] <= 100
 */

/*
wow my god, welcome to the algorithms world!
brilliant questions!
完全是自己想的， 太赞了！

https://leetcode.cn/problems/maximum-number-of-alloys/solutions/2446024/er-fen-da-an-fu-ti-dan-by-endlesscheng-3jdr/

可以预估一个 max upper bound instead of using 1000000000

 */
func canMake(p int, composition []int, stock []int, cost []int, budget int) bool {

	total := 0
	for i, c := range composition {
		total += max(c*p-stock[i], 0) * cost[i]
	}

	return total <= budget
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	ans := 0

	ms := 0
	for _, s := range stock {
		ms = max(ms, s)
	}
	ceiling := budget + ms

	// loop 机器
	for i := 0; i < k; i++ {
		//l, r := 0, 1000000000
		l, r := 0, ceiling
		for l <= r {
			mid := (l + r) / 2
			if canMake(mid, composition[i], stock, cost, budget) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		// 这里直接 Loop 会超时， 所以想到了可以，像上一样做二分。
		//for canMake(p, composition[i], stock, cost, budget) {
		//	p++
		//}
		ans = max(ans, l-1)
	}
	return ans
}
