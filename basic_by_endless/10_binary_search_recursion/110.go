package binary_tree

/*
好无聊，好无聊!
 */
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func isBalanced(root *TreeNode) bool {

	var dfs func(*TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		lb, l := dfs(node.Left)
		rb, r := dfs(node.Right)
		return lb && rb && abs(l, r) <= 1, max(l, r) + 1
	}
	ok, _ := dfs(root)
	return ok
}