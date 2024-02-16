package interview_questions


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 /**
 两遍 dfs
  */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	clones := make(map[*Node]*Node)
	var dfs func(*Node)
	dfs = func (r *Node) {
		// clone r
		clone := &Node {
			Val: r.Val,
		}
		clones[r] = clone

		for _, n := range r.Neighbors {
			if _, ok := clones[n]; !ok {
				dfs(n)
			}
		}
	}
	dfs(node)

	visited :=make(map[*Node]struct{})
	var dfs2 func(*Node)
	dfs2 = func(r *Node) {
		// mark visited
		visited[r] = struct{}{}
		// clone neighbours
		neigh := make([]*Node, 0, len(r.Neighbors))
		for _, n := range r.Neighbors {
			neigh = append(neigh, clones[n])
		}
		clones[r].Neighbors = neigh

		for _, n := range r.Neighbors {
			if _, ok := visited[n]; !ok {
				dfs2(n)
			}
		}
	}
	dfs2(node)
	return clones[node]
}


/***
一遍 dfs
 */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clones := make(map[*Node]*Node)
	visited := make(map[*Node]struct{})
	var dfs func(*Node)
	dfs = func (r *Node) {
		// mark as visited
		visited[r] = struct{}{}
		// clone r
		var clone *Node
		if _, ok :=clones[r]; !ok {
			clone = &Node{
				Val: r.Val,
			}
			clones[r] = clone
		} else {
			clone = clones[r]
		}

		neigh := make([]*Node, 0, len(r.Neighbors))

		// clone neighbours
		for _, n := range r.Neighbors {
			if v, ok := clones[n]; !ok {
				c := &Node{
					Val:n.Val,
				}
				clones[n] = c
				neigh = append(neigh, c)
			} else {
				neigh = append(neigh, v)
			}
		}
		clone.Neighbors = neigh

		for _, n := range r.Neighbors {
			if _, ok :=visited[n]; !ok {
				dfs(n)
			}
		}
	}
	dfs(node)
	return clones[node]
}
