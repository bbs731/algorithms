package weekly

import "sort"

/****
这个 dfs 是灵神的版本。

多学习下， 如果才能写出来？

 */

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	m := len(robot)
	n := len(factory)

	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	inf := int(1e18)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, m)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	// f[i][j]  代表， [i, n-1] 的工厂，修理了 [j, m-1] 个机器人的最小值。
	// i 是工厂，  j 是机器人的下标。
	var dfs func(int, int) int
	dfs = func(i int, j int) int{
		if j == m {
			return 0
		}
		if i == n -1 {
			if m-j > factory[i][1] {
				return inf
			}
			s := 0
			for k := j; k<m; k++ {
				s += abs(factory[i][0], robot[k])
			}
			return s
		}

		if cache[i][j] != - 1 {
			return cache[i][j]
		}
		res := dfs(i+1, j)
		s := 0
		for k:=1; k<=factory[i][1] && j+k-1 <m; k++ {
			s += abs(factory[i][0], robot[j+k-1])
			res = min(res, dfs(i+1, j+k) + s )
		}

		cache[i][j] = res
		return res
	}
	return int64(dfs(0, 0))
}


func abs (a, b int)int {
	if a > b {
		return a - b
	}
	return b-a
}
