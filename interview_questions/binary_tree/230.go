package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {

	var ans *TreeNode

	var dfs func(*TreeNode, int) int
	dfs = func(r *TreeNode, pre int) int {
		if r == nil {
			return 0
		}
		ls := dfs(r.Left, pre)
		if ls+pre+1 == k {
			ans = r
		}
		rs := dfs(r.Right, ls+pre+1)
		return ls + 1 + rs
	}

	dfs(root, 0)
	return ans.Val
}

/***
这个代码，更加的简洁啊！
 */
func kthSmallest(root *TreeNode, k int) int {
	var ans *TreeNode
	cnt := 0

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		cnt++
		if cnt == k {
			ans = r
		}
		dfs(r.Right)
	}

	dfs(root)
	return ans.Val
}
