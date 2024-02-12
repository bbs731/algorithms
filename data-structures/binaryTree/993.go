package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCousins(root *TreeNode, x int, y int) bool {

	xl, yl, xp, yp := -1, -2, -1, -2

	var dfs func(*TreeNode, int, int)
	dfs = func(cur *TreeNode, parent, level int) {
		if cur == nil {
			return
		}
		if cur.Val == x {
			xp = parent
			xl = level
		}
		if cur.Val == y {
			yp = parent
			yl = level
		}
		dfs(cur.Left, cur.Val, level+1)
		dfs(cur.Right, cur.Val, level+1)
	}

	dfs(root, -1, 0)

	return xl == yl && xp != yp
}
