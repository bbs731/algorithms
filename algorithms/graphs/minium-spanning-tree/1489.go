package minium_spanning_tree

import (
	"fmt"
	"sort"
)

/*

输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
输出：[[0,1],[2,3,4,5]]
解释：上图描述了给定图。
下图是所有的最小生成树。

 */

// https://oi-wiki.org/graph/mst/   中判断 MST 是否唯一的例子，借不了这道题。
// 换一个枚举的笨方法。
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	must := make([]int, 0)
	non_must := make([]int, 0)
	cedges := make([][]int, len(edges))
	copy(cedges, edges)

	// 特判一下吧： 这个特判不成立， 存在， 所有 edges 相等， 然后，带环的情况， 处理不了。
	// https://oi-wiki.org/graph/mst/   中判断 MST 是否唯一的例子，借不了这道题。
	p := []int{}
	for j := 0; j < len(edges); j++ {
		if edges[j][2] != edges[0][2] {
			break
		}
		p = append(p, j)
	}
	if len(p) == len(edges) && len(p) != 1 {
		return [][]int{[]int{}, p}
	}

	inMst, un := mstKruskal(n, cedges)
	fmt.Println(inMst)
	fmt.Println(un)
	for i, e := range edges {
		_, found := un[[2]int{e[0], e[1]}]
		if found {
			non_must = append(non_must, i)
		} else {
			if _, ok := inMst[[2]int{e[0], e[1]}]; ok {
				//fmt.Println(e, i)
				must = append(must, i)
			}
		}
	}
	return [][]int{must, non_must}
}

func mstKruskal(n int, edges [][]int) (map[[2]int]bool, map[[2]int]bool) {
	// 边权范围小的话也可以用桶排
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := make(map[[2]int]bool) // edges with same cost
	inMst := make(map[[2]int]bool)
	sum := 0
	cntE := 0
	start := 0
	tail := -1
	sum1 := 0
	sum2 := 0
	for i, e := range edges {
		if i > tail {
			if sum1 > 1 && sum1 != sum2 {
				for k := start; k <= tail; k++ {
					ans[[2]int{edges[k][0], edges[k][1]}] = true
				}
			}
			sum1 = 0
			start = i
			for j := i; j < len(edges); j++ {
				if edges[j][2] != edges[i][2] {
					tail = j - 1
					break
				}
				if find(edges[j][0]) != find(edges[j][1]) {
					sum1++
				}
			}
			sum2 = 0
		}

		v, w, wt := e[0], e[1], e[2]
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
			sum += wt
			inMst[[2]int{v, w}] = true
			sum2++
			cntE++
		}
	}

	// 图不连通
	//if cntE < n-1 {
	//	return -1, nil
	//}
	return inMst, ans
}
