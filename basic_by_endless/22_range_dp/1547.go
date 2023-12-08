package dp

import (
	"math"
	"sort"
	"strconv"
)

/*
ToDo:  用递推再写一遍！ 区间DP的题目，绝了！
如果需要用到很多的技巧，代表思路是错的。

 */

/*
	我们先预处理一下，分块， 然后考虑一下，如何把分块组合起来。 让组合的 cost 最小。

	dfs(i, j) 是块 i 到 j 组合 所花费的 cost  = min (dfs(i,k) + dfs(k,j) + j-i)
	这里面的 i,j 是 我们预处理过之后的 Index,  0 <=i, j <=100

	这是第三次尝试了吧？ 坚持住
 */

func minCost(n int, cuts []int) int {
	l := len(cuts)
	sort.Ints(cuts)
	inf := int(1e9)

	cache := make([][]int, l+1)
	for i := range cache {
		cache[i] = make([]int, l+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	blocks := make([][2]int, l+1)
	for i := 0; i < l; i++ {
		blocks[i][1] = cuts[i]
		blocks[i+1][0] = cuts[i]
	}
	blocks[0][0] = 0
	blocks[l][1] = n

	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i >= j {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		ans := inf
		for k := i; k < j; k++ {
			ans = min(ans, dfs(i, k)+dfs(k+1, j))
		}
		ans += blocks[j][1] - blocks[i][0]
		cache[i][j] = ans
		return ans
	}

	return dfs(0, l)
}

/*
	dfs(i,j) = dfs(i, cut[k]) + dfs(cut[k], j) + j-i   for all cuts inside [i, j]
	n 的取值范围是  10^6  memory 会 over limit
 */

func minCost_memory_limit(n int, cuts []int) int {
	sort.Ints(cuts)
	inf := int(1e9)
	//
	cache := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		// n = 7
		// [1,3,4,5]

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		// startpos > i
		startpos := sort.SearchInts(cuts, i+1)
		// endpos < j
		endpos := sort.SearchInts(cuts, j) - 1

		if startpos > endpos {
			return 0
		}

		ans := inf
		for k := startpos; k <= endpos && k < len(cuts); k++ {
			//if cuts[k] > i && cuts[k] < j {
			ans = min(ans, dfs(i, cuts[k])+dfs(cuts[k], j)+j-i)
			//}
		}

		cache[i][j] = ans
		return ans
	}

	return dfs(0, n)
}

func hashToString(cuts []int) string {
	str := ""
	for _, v := range cuts {
		str += strconv.Itoa(v)
	}
	return str
}

func minCost_time_limit(n int, cuts []int) int {
	sort.Ints(cuts)
	var dfs func(int, []int) int

	cache := make([]map[string]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make(map[string]int)
	}

	dfs = func(l int, cs []int) int {
		if len(cs) == 0 {
			return 0
		}

		if val, ok := cache[l][hashToString(cs)]; ok {
			return val
		}

		ans := math.MaxInt
		for i := 0; i < len(cs); i++ {
			cut := cs[i]
			left := make([]int, 0)
			right := make([]int, 0)
			for j := 0; j < len(cs); j++ {
				if j == i {
					continue
				}
				if cs[j] < cut {
					left = append(left, cs[j])
				} else {
					right = append(right, cs[j]-cut)
				}
			}
			ans = min(ans, dfs(cut, left)+dfs(l-cut, right)+l)
		}
		cache[l][hashToString(cuts)] = ans
		return ans
	}
	return dfs(n, cuts)
}
