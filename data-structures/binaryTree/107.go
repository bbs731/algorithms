package binaryTree

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	//ans = append(ans, []int{root.Val})

	for len(q)>0 {
		tmp := q
		q = nil
		level := make([]int, 0)
		for _, n := range tmp{
			level = append(level, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		ans = append(ans, level)
	}

	//j := len(ans) -1
	//for i:=0; i<j; i++ {
	//	ans[i], ans[j]= ans[j], ans[i]
	//	j--
	//}
	// 灵神教会的，有  slice 翻转的 api 可以直接使用
	slices.Reverse(ans)
	return ans
}
