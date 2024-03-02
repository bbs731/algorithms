package dp

import "sort"

/****

1. 变量是什么？
i 个 house
j 个油桶。

2. 枚举， 最后一个 post 的位置。
f[i][j] 怎么来的？ = min(f[i0][j-1] + cost[i0+1,i] for all possible i0)

f[i][j+1] = min (f[i0][j] + cost[i0+1, i])

3. 定义状态：
f[i][j] 用 j 个油桶， cover 前 i 个 housse 的最小和。

初始化条件
f[i][j] for j >= i  set to 0
f[0][0] = 0
f[i][1]=cost(0,i)

 */
func minDistance(houses []int, k int) int {
	 n := len(houses)
	 sort.Ints(houses)

	 costs := make([][]int, n)
	 for i := range costs {
	 	costs[i] = make([]int, n)
	 }

	 // pre compute costs
	 for i := n-2; i>=0;i-- {
	 	for j := i+1; j<n; j++ {
	 		costs[i][j] = costs[i+1][j-1] + houses[j] - houses[i]
		}
	 }

	 f := make([][]int, n)
	 for i := range f{
	 	f[i]= make([]int, k+1)
	 	for j := range f[i] {
	 		f[i][j] = int(1e15)
		}
		 f[i][1] = costs[0][i]
	 }
	 for i:=0;i<n; i++ {
	 	for j:=1; j<k && j <= i; j++ {
			for l0:=0; l0 < i; l0++ {
				/***
					这道题， 用 i+1 作为 index 就是一个灾难
				 */
				//f[i+1][j+1] = min(f[i+1][j+1], f[l0][j] + costs[l0+1][i])
				f[i][j+1] = min(f[i][j+1], f[l0][j] + costs[l0+1][i])
			}
		}
	 }
	 return f[n-1][k]
}
