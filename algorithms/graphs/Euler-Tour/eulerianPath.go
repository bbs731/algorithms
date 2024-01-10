package Euler_Tour

import (
	"slices"
	"sort"
)

// code borrowed from :
// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go
// video explaination from : https://www.youtube.com/watch?v=8MpoO2zA2l4

// 找 Eulerian Circle 可以用 Eulerian Path 的算法， 因为 euler circle 可以从任意起点开始 dfs。

// namespace
type graph struct{}

// 时间复杂度是 linear 的：  O（V+E) or O(n+m)
// 最后 path 的 len(path) = m+1
func (*graph) eulerianPathOnDirectedGraph(n, m int) []int {
	type neighbour struct{ to, eid int } // vertex id number, and edge id number
	g := make([][]neighbour, n)
	inDeg := make([]int, n) // inDegree 入度
	// read g ... and 统计入度

	// 排序，保证字典序 (什么题目要求字典序？例子: LC332）
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].to < es[j].to })
	}

	st := -1
	end := -1
	for i, es := range g {
		if len(es) == inDeg[i]+1 { // 出度 = 入度 + 1, 起点
			if st >= 0 {
				return nil // 错误，无欧拉路
			}
			st = i
		}
		if len(es)+1 == inDeg[i] { // 入度 = 出度 + 1， 终点
			if end >= 0 {
			}
			return nil // 错误
		}
		end = i
	}

	if st < 0 {
		st = 0 // 可以随便选起点， 因为存在 Eulerian Circle, 当然就存在从任意起点开始的 Eulerian Path
	}

	path := make([]int, 0, m+1)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			dfs(e.to) //如果想要 edge id 的话，可以放入  e.eid
		}
		path = append(path, v)
	}
	dfs(st)
	slices.Reverse(path)
	// len(path) = m +1
	return path
}

func (*graph) eulerianPathOnUndirectedGraph(n, m int) []int {
	type neighbour struct{ to, eid int }
	g := make([][]neighbour, n)
	// read g...

	//排序，保证字典序最小
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].to < es[j].to })
	}
	var st int
	oddDegCnt := 0

	for i := len(g) - 1; i >= 0; i-- { //倒着遍历保证起点的字典序最小
		if deg := len(g[i]); deg > 0 {
			if deg&1 == 1 {
				st = i
				oddDegCnt++
			} else if oddDegCnt == 0 {
				st = i
			}
		}
	}
	if oddDegCnt > 2 {
		return nil
	}

	path := make([]int, 0, m+1)
	vis := make([]bool, m)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			i := e.eid
			if vis[i] {
				continue
			}
			vis[i] = true
			dfs(e.to)
		}
		// 输出点的写法， len(path) 最后 = m+1
		path = append(path, v)
	}
	dfs(st)
	slices.Reverse(path)
	return path
}
