package Biconnected_Components

import "sort"

package one_day_exercise

import (
"fmt"
"slices"
"sort"
)


// 题解来自这里。
// https://leetcode.cn/problems/s5kipK/solutions/1425679/by-tsreaper-z8by/

// 我考虑清楚了， 需要找到 所有的 bcc 然后取, 前  len(bcc) -1 个 bcc componenet,  每个 bcc componenet 里取最小的 vertex cost  然后相加。
// 思路是对的， 但是，犯了几个错误。
// 1. 没区分 scc 和 bcc 的区别， 这个，需要再学习bcc, 因为之前没写过 bcc 的代码, 这个知识点拉掉了。
// 2. 像题解里说的。 不是每个 bcc 都要参与到最后的计算当中的， 只有哪些，缩点之后是叶子节点的 bcc 在需要考虑在最后的答案计算当中的。 （想想为什么？ 链接 2个 或者两个以上 cut-vertex 的 bcc 缩点之后，就是non-leaf 节点。）
// 3. 如何知道一个 bcc 缩点之后是叶子节点， （题目作的时候不需要真的缩点） 当， bcc 中所有的 vertex 中，有且只有一个 cut vertex 那么，这个 bcc 就是缩点之后的一个叶子节点。
// 4. bcc 的代码，来自 灵神的代码库 codeforeces-go/graph.go  还没读过。 抓紧。

// 另外， 这真是一个好题。  图根据 bcc 缩点之后，变成一颗树. 考虑所有的树的叶子节点， 然后把 排名 n-1 之前的叶子节点的 cost 都加起来就是的答案。 多么好的题啊。
func findVertexBCC(g [][]int) (comps [][]int,  isCut[]bool) {
	bccIDs := make([]int, len(g)) // ID 从 1 开始编号
	idCnt := 0
	isCut = make([]bool, len(g))

	dfn := make([]int, len(g))
	dfsClock := 0
	type edge struct{ v, w int } // eid
	st := []edge{}               // 存边是为了解决一些特殊题目（基本写法存点就行）
	var tarjan func(v, fa int) int
	tarjan = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			e := edge{v, w} // ne.eid
			if dfn[w] == 0 {
				st = append(st, e)
				childCnt++
				lowW := tarjan(w, v)
				if lowW >= dfn[v] {
					isCut[v] = true
					idCnt++
					comp := []int{}
					//eids := []int{}
					for {
						e, st = st[len(st)-1], st[:len(st)-1]
						if bccIDs[e.v] != idCnt {
							bccIDs[e.v] = idCnt
							comp = append(comp, e.v)
						}
						if bccIDs[e.w] != idCnt {
							bccIDs[e.w] = idCnt
							comp = append(comp, e.w)
						}
						//eids = append(eids, e.eid)
						if e.v == v && e.w == w {
							break
						}
					}
					comps = append(comps, comp)
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] {
				st = append(st, e) // 简单写法中，可以省略
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 {
			isCut[v] = false
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			if len(g[v]) == 0 { // 零度，即孤立点（isolated vertex）
				idCnt++
				bccIDs[v] = idCnt
				comps = append(comps, []int{v})
				continue
			}
			tarjan(v, -1)
		}
	}
	return
}


func minimumCost(cost []int, roads [][]int) int64 {
	n := len(cost)
	g := make([][]int, n)
	const inf int = 1e17
	for _, e := range roads {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dcc, isCut := findVertexBCC(g)
	ans := 0

	if len(dcc) == 1 {
		//return int64(min(cost))
		f := inf
		for _, x := range cost{
			f = min(f, x)
		}
		return int64(f)
	}

	candidates := make([]int, 0)

	for _, c := range dcc {
		cnt :=0
		mn := inf
		for _, x := range c {
			if isCut[x] {
				cnt++
			} else {
				mn = min(mn, cost[x])
			}
		}
		if cnt == 1 {
			candidates = append(candidates, mn)
		}
	}

	sort.Ints(candidates)

	for i:=0; i+1 < len(candidates); i++ {
		ans += candidates[i]
	}
	return int64(ans)
}
