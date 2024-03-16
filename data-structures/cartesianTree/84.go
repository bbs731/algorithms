package cartesianTree

/***
https://oi-wiki.org/ds/cartesian-tree/
解题的思路，在这里给出。
 */

type ctNode struct {
	lr      [2]*ctNode
	id, val int
}

func buildCartesianTree(a []int) *ctNode {
	if len(a) == 0 {
		return nil
	}

	s := []*ctNode{}

	for i, v := range a {
		o := &ctNode{id: i, val: v}
		for len(s) > 0 {
			top := s[len(s)-1]
			if top.val < v {
				top.lr[1] = o
				break
			}
			o.lr[0] = top
			s = s[:len(s)-1]
		}
		s = append(s, o)
	}
	return s[0]
}

func largestRectangleArea(heights []int) int {
	root := buildCartesianTree(heights)
	ans := 0

	var dfs func(*ctNode) int

	dfs = func(r *ctNode) int {
		if r == nil {
			return 0
		}
		sz := dfs(r.lr[0])
		sz += dfs(r.lr[1])
		ans = max(ans, (sz+1)*r.val)
		return sz + 1
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
