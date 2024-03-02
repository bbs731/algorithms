package weekly

import "sort"

/****
DP问题思考步骤：

1. 有哪些变量

2. 用变量的值 重述一遍问题。

3. 最后一步发生了什么？ （关键步骤）

4. 去掉最后一步， 问题规模缩小了， 变成什么样了，（用新的变量值，重述一遍问题)

5. 得到状态转移方程
可以通过步骤 2 (定义状态） 和 4 （状态转移方程）

6. 初始值，和 答案 分别是什么？

7. 优化转移。

1. 变量是什么？
i robots的id
j 工厂的 id
and
limit

dp[j][i]

2.
dp[j][i] = min (dp[j-1][i-k] + costs[k]) for k =0 ... limit[j])

5. 所求是什么？
	dp[n][m]

 */


 /****

 这个写法， 太朴素了， 是在 loop 值域， 所以一定会 TML
 我现在的阶段 2024-Mar-2 就是这个水平。 想了一个小时也做不出来 DP状态的定义， 和状态转移方程。
 最后看了灵神的视频解答： 写出了 dfs  2463_3.go 的解法，  接着， 有写出了递推的解法 2463_4.go
 最后回到自己最初的方法，写出了  2463_1.go 的非常赞的解法！ 独一份。 自己提高了， 但是，第一步迈出去，太困难了！

  */
 type pair struct {
 	dist int
 	j int
 }

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	n := len(robot)
	dist := make([][]pair, n)

	capacity := make(map[int]int)
	for j := range factory {
		capacity[j] = factory[j][1]
	}

	for i := range robot {
		for j := range factory {
			dist[i] = append(dist[i], pair{abs(robot[i], factory[j][0]), j})
		}
		sort.Slice(dist[i], func(m, n int) bool {
			return dist[i][m].dist < dist[i][n].dist
		})
	}

	ans := int(1e16)
	var dfs func(int, int)
	dfs = func(i int, cost int){
		if i == n {
			ans = min(ans, cost)
			return
		}
		if cost >= ans {
			return
		}
		for _, p := range dist[i] {
			if capacity[p.j] > 0 && cost + p.dist < ans {
				capacity[p.j]--
				dfs(i+1, cost + p.dist)
				capacity[p.j]++
			}
		}
	}
	dfs(0, 0)
	return int64(ans)
}


