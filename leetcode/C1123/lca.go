package C1123

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

var mem map[*TreeNode]int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func DepthOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := mem[root]; ok {
		return v
	}
	mem[root] = max(DepthOfTree(root.Left), DepthOfTree(root.Right)) + 1
	return mem[root]
}

func recurLcaDeepestLeaves(root *TreeNode) *TreeNode {
	l, r := DepthOfTree(root.Left), DepthOfTree(root.Right)
	if l == r {
		return root
	} else if l > r {
		return recurLcaDeepestLeaves(root.Left)
	} else {
		return recurLcaDeepestLeaves(root.Right)
	}
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	mem = make(map[*TreeNode]int)
	return recurLcaDeepestLeaves(root)
}
