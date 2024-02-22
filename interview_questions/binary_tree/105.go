package binary_tree



func buildTree(preorder []int, inorder []int) *TreeNode {

	n := len(preorder)
	if n == 0 {
		return nil
	}

	pi := make(map[int]int)
	for i, n := range inorder {
		pi[n] = i
	}

	p := pi[preorder[0]]
	left := buildTree(preorder[1:p+1], inorder[:p])
	right := buildTree(preorder[p+1:], inorder[p+1:])

	return &TreeNode{preorder[0], left, right}
}
