package Strongly_Connected_Component

import "slices"

// code from  endlesswang  codeforces-go
// video explaination:  https://www.youtube.com/watch?v=wUgWX0nc4NY&list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P&index=23

func sccTarjan(g [][]int) ([][]int, []int) {
	scc := [][]int{}
	dfn := make([]int, len(g))
	dfsClock := 0
	st := []int{}
	inSt := make([]bool, len(g))

	var tarjan func(int)
	tarjan = func(v int) {
		dfsClock++
		dfn[v] = dfsClock
		lowW := dfsClock // low-link
		st = append(st, v)
		inSt[v] = true
		for _, w := range g[v] {
			if dfn[w] == 0 { //unvisited
				tarjan(w)
			}
			if inSt[w] {
				lowW = min(lowW, dfn[w])
			}
		}

		if dfn[v] == lowW {
			comp := []int{}
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				inSt[w] = false
				comp = append(comp, w)
				if w == v { // v is one of the SCC root. See the video explaination
					break
				}
			}
			scc = append(scc, comp)
		}
	}

	for i, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(i)
		}
	}

	//// https://stackoverflow.com/questions/32750511/does-tarjans-scc-algorithm-give-a-topological-sort-of-the-scc
	slices.Reverse(scc)

	// Tarjan SCC 算法到此结束了。 下面的，把 SCC 缩点变成一个新的 graph
	sid := make([]int, len(g))
	for i, cc := range scc {
		for _, v := range cc {
			sid[v] = i
		}
	}
	ns := len(scc)
	g2 := make([][]int, ns)
	deg := make([]int, ns)
	for v, ws := range g {
		v = sid[v]
		for _, w := range g[v] {
			if v != sid[w] {
				g2[v] = append(g2[v], w)
				deg[w]++
			}
		}
	}
	return scc, sid
}
