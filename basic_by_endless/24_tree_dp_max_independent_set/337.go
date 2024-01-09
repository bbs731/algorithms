package _4_tree_dp_max_independent_set

// code copied from:  https://leetcode.cn/problems/house-robber-iii/solution/shi-pin-ru-he-si-kao-shu-xing-dppythonja-a7t1/
// 视频讲解： https://www.bilibili.com/video/BV1vu4y1f7dn/?spm_id_from=333.788&vd_source=84c3c489cf545fafdbeb3b3a6cd6a112

// 这代码，也是太优雅了吧！

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}
		l_rob, l_not_rob := dfs(r.Left)
		r_rob, r_not_rob := dfs(r.Right)

		rob := l_not_rob + r_not_rob + r.Val
		not_rob := max(l_rob, l_not_rob) + max(r_rob, r_not_rob)

		return rob, not_rob
	}
	return max(dfs(root))
}
