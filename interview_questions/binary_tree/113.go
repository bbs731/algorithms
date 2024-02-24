package binary_tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(*TreeNode, int, []int)

	ans := make([][]int, 0)
	dfs = func(r *TreeNode, sum int, path[]int) {
		if  r == nil {
			return
		}
		if r.Left == nil && r.Right == nil && r.Val == sum {
			// record ans
			path = append(path, r.Val)
			ans = append(ans, append([]int{}, path...))
			path = path[:len(path)-1]
			return
		}

		path = append(path, r.Val)
		dfs(r.Left, sum-r.Val, path)
		dfs(r.Right, sum-r.Val, path)
		path = path[:len(path)-1]
	}

	dfs(root, targetSum, []int{})
	return ans
}
