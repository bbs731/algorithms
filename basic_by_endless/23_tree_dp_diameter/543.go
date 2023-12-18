package _3_tree_dp_diameter

//Given the root of a binary tree, return the length of the diameter of the tree
//.
//
// The diameter of a binary tree is the length of the longest path between any t
//wo nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges
//between them.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4,5]
//Output: 3
//Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
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
func diameterOfBinaryTree(root *TreeNode) int {
	ans := -1
	// dfs 返回是树的深度， 深度在这道题里指的是边的个数，不是 Node 的个数。
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return -1
		}
		lh := dfs(r.Left)
		rh := dfs(r.Right)
		ans = max(ans, lh+rh+2)
		return max(lh, rh) + 1
	}
	dfs(root)
	return ans
}

// 上面可以变换成 ：

func diameterOfBinaryTree(root *TreeNode) int {
	ans := -1
	// dfs 返回是树的深度， 深度在这道题里指的是边的个数，不是 Node 的个数。
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return -1
		}
		lh := dfs(r.Left) + 1
		rh := dfs(r.Right) + 1
		ans = max(ans, lh+rh)
		return max(lh, rh)
	}
	dfs(root)
	return ans
}

func diameterOfBinaryTree_node_length(root *TreeNode) int {
	ans := -1
	// dfs 返回是树的深度， 深度在这个版本里指的是Node的个数，不是边的个数。
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		lh := dfs(r.Left)
		rh := dfs(r.Right)
		ans = max(ans, lh+rh+1)
		return max(lh, rh) + 1
	}
	dfs(root)
	return ans - 1
}
