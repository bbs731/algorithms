package _600_1900

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func isCompleteTree(root *TreeNode) bool {
	// bfs
	if root == nil {
		return true
	}
	s := []*TreeNode{root}
	sb := make([]*TreeNode, 0)
	mark := false
	for len(s) != 0 {
		for _, e := range s {
			if mark && (e.Left != nil || e.Right != nil) {
				return false
			}
			if e.Left != nil {
				sb = append(sb, e.Left)
			} else {
				mark = true
				if e.Right != nil {
					return false
				}
			}
			if e.Right != nil {
				sb = append(sb, e.Right)
			} else {
				mark = true
			}
		}
		s = sb
		//sb = sb[:0]   这里如果使用 sb = sb[：0] 是不是也改变了 s, 这样是错误的。
		sb = make([]*TreeNode, 0)
	}
	return true
}
