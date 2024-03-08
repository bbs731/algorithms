package binarySearchTree

import "math"

/***
越简单的题目， 死的越快！
 */

func isValidBST(root *TreeNode) bool {
	return help(root, math.MinInt64, math.MaxInt64)
}

func help(r *TreeNode, low, high int) bool {
	if r == nil {
		return true
	}
	if r.Val <= low || r.Val >= high {
		return false
	}
	return help(r.Left, low, r.Val) && help(r.Right, r.Val, high)
}
