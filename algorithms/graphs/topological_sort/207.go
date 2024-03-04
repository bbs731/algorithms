package topological_sort

/***

 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	n := numCourses
	g := make([][]int, n)
	deg := make([]int, n)

	for _, c := range prerequisites {
		pp, p := c[0], c[1]
		g[p] = append(g[p], pp)
		deg[pp]++
	}

	q := []int{}
	// topological sort result list
	l := []int{}
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		l = append(l, v)
		q = q[1:]
		for _, w := range g[v] {
			deg[w]--
			if deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	return len(l) == numCourses
}
