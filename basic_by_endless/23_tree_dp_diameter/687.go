package _3_tree_dp_diameter

//Given the root of a binary tree, return the length of the longest path, where
//each node in the path has the same value. This path may or may not pass through
//the root.
//
// The length of the path between two nodes is represented by the number of edge
//s between them.
//
//
// Example 1:
//
//
//Input: root = [5,4,5,1,1,null,5]
//Output: 2
//Explanation: The shown image shows that the longest path of the same value (i.
//e. 5).
//
//
// Example 2:
//
//
//Input: root = [1,4,5,4,4,null,5]
//Output: 2
//Explanation: The shown image shows that the longest path of the same value (i.
//e. 4).
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10^4].
// -1000 <= Node.val <= 1000
// The depth of the tree will not exceed 1000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int // count edges instead of nodes for this question
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		lh := dfs(r.Left)
		rh := dfs(r.Right)

		if r.Left != nil && r.Val == r.Left.Val {
			lh++
		} else {
			lh = 0 // 这里是个坑啊， 如果  r.Val != r.Left.Val 或者 r.Left == nil 需要 reset lh = 0
		}
		if r.Right != nil && r.Val == r.Right.Val {
			rh++
		} else {
			rh = 0 // 同理
		}
		ans = max(ans, lh+rh)

		return max(lh, rh)
	}
	dfs(root)
	return ans
}

