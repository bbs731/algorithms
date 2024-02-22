package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {

	n := len(preorder)
	if n == 0 {
		return nil
	}

	// 这么写不对 这个 hash table 创建了好几遍。
	pi := make(map[int]int)
	for i, n := range inorder {
		pi[n] = i
	}

	p := pi[preorder[0]]
	left := buildTree(preorder[1:p+1], inorder[:p])
	right := buildTree(preorder[p+1:], inorder[p+1:])

	return &TreeNode{preorder[0], left, right}
}
