package C99

import "sort"

var m map[int]*TreeNode

func InOrderTraversalTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	lf := InOrderTraversalTree(root.Left)
	lf = append(lf, root.Val)
	m[root.Val] = root
	lr := InOrderTraversalTree(root.Right)
	lf = append(lf, lr...)
	return lf
}

func recoverTree2(root *TreeNode) {
	m = make(map[int]*TreeNode)

	l := InOrderTraversalTree(root)
	cl := make([]int, len(l))
	copy(cl, l)
	sort.Ints(cl)

	for i, _ := range l {
		if l[i] != cl[i] {
			m[l[i]].Val = cl[i]
			m[cl[i]].Val = l[i]
			return
		}
	}
}
