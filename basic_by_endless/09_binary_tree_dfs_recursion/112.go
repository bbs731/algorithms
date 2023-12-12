package binary_tree

//Given the root of a binary tree and an integer targetSum, return true if the t
//ree has a root-to-leaf path such that adding up all the values along the path eq
//uals targetSum.
//
// A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//Output: true
//Explanation: The root-to-leaf path with the target sum is shown.
//
//
// Example 2:
//
//
//Input: root = [1,2,3], targetSum = 5
//Output: false
//Explanation: There two root-to-leaf paths in the tree:
//(1 --> 2): The sum is 3.
//(1 --> 3): The sum is 4.
//There is no root-to-leaf path with sum = 5.
//
//
// Example 3:
//
//
//Input: root = [], targetSum = 0
//Output: false
//Explanation: Since the tree is empty, there are no root-to-leaf paths.
//
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	对于 leaf node 的理解 不对。  leaf node 是左右 node 都为 nil 的 node
 */

func hasPathSum(r *TreeNode, targetSum int) bool {
	if r == nil {
		return false
	}
	if r.Left == nil && r.Right == nil && r.Val == targetSum {
		return true
	}

	return hasPathSum(r.Left, targetSum-r.Val) || hasPathSum(r.Right, targetSum-r.Val)
}

