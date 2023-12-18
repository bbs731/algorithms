package _3_tree_dp_diameter

/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

 */

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func maxPathSum(root *TreeNode) int {

	ans := -int(1e9)
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		tmp := 0
		lh := dfs(r.Left)
		rh := dfs(r.Right)
		if lh > 0 {
			tmp += lh
		}
		if rh > 0 {
			tmp += rh
		}
		ans = max(ans, tmp+r.Val)
		return max(max(lh, rh)+r.Val, 0) //这里是一个大坑！
	}
	dfs(root)
	return ans
}

/*
灵神说，上面的代码可以简化！
https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/2227021/shi-pin-che-di-zhang-wo-zhi-jing-dpcong-n9s91/
 */

func maxPathSum(root *TreeNode) int {

	ans := -int(1e9)
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		lh := dfs(r.Left)
		rh := dfs(r.Right)
		ans = max(ans, lh+rh+r.Val)      // 因为下面 return 之前和 0 取了最大值，所以这里，可以直接加 lh, rh 他们都 >=0
		return max(max(lh, rh)+r.Val, 0) //这里是一个大坑！
	}
	dfs(root)
	return ans
}
