package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int // 可以取巧
	var cur []*TreeNode
	var next []*TreeNode // scope 定义的有问, 既然每个循环都会 reset next, 为啥不直接定义在loop 里。
	var values []int     // scope 的问题， 没必要定义在最外层. 看 102_1_ammend.go 的修改

	if root == nil {
		return ans
	}

	cur = append(cur, root)
	values = append(values, root.Val)

	for len(cur) > 0 {
		ans = append(ans, append([]int(nil), values...))
		next = []*TreeNode{}
		values = []int{}

		for _, n := range cur {
			if n.Left != nil {
				next = append(next, n.Left)
				values = append(values, n.Left.Val)
			}
			if n.Right != nil {
				next = append(next, n.Right)
				values = append(values, n.Right.Val)
			}
		}
		cur = next
	}
	return ans
}
