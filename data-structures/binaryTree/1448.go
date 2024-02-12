package binaryTree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/***
https://leetcode.cn/problems/count-good-nodes-in-binary-tree/solutions/2403677/jian-ji-xie-fa-pythonjavacgojs-by-endles-gwxt/
自信点， 恐惧会削弱你
 */
func goodNodes(root *TreeNode) int {
	ans := 0

	var dfs func(*TreeNode, int)
	dfs = func(r *TreeNode, p int) {
		if r == nil {
			return
		}
		if r.Val >= p {
			ans++
		}
		dfs(r.Left, max(r.Val, p))
		dfs(r.Right, max(r.Val, p))
	}

	//dfs(root, -int(1e5))
	dfs(root, math.MinInt)
	return ans
}
