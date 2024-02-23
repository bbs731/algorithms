package binary_tree

import "math"

/***
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

 */

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

/***
免疫了吗？
 */

func maxPathSum(root *TreeNode) int {
	ans := -math.MaxInt32

	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		lv := dfs(r.Left)
		rv := dfs(r.Right)

		ans = max(ans, lv+rv+r.Val)
		return max(max(lv, rv)+r.Val, 0)
	}

	dfs(root)
	return ans
}
