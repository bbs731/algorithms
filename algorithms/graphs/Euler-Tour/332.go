package Euler_Tour

import (
	"slices"
	"sort"
)

/*

给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。

所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。


输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
输出：["JFK","MUC","LHR","SFO","SJC"]



输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。


 */

// 采用的代码是， williams 的 eulerianPath 的代码。
// 应该是使用 有向图和无向图，都可以。 本题是有向图。建图的时候注意

func findItinerary(tickets [][]string) []string {
	type neighbour struct {
		to  string
		eid int
	} // eid edge's id 这道题输出答案是 vertex 所以应该没用到。
	g := make(map[string][]neighbour)

	// 建图
	for _, citys := range tickets {
		s, e := citys[0], citys[1]
		g[s] = append(g[s], neighbour{e, len(g[e])})
		//g[e] = append(g[e], neighbour{s, true, len(g[s]) - 1})
	}

	// 字典序 sort
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].to < es[j].to })
	}

	st := "JFK"

	path := make([]string, 0)

	var dfs func(string)
	dfs = func(v string) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			dfs(e.to)
		}
		path = append(path, v)
	}
	dfs(st)

	slices.Reverse(path)
	return path
}
