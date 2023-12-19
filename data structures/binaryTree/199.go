package binaryTree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 灵神更牛叉的解法： (dfs)
 https://leetcode.cn/problems/binary-tree-right-side-view/solutions/2015061/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-r1nc/?envType=study-plan-v2&envId=top-100-liked
 */
func rightSideView(root *TreeNode) (ans []int) {
	aq := []*TreeNode{}
	if root == nil {  // 边界条件的检查是一个坑。
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		for len(q) > 0 {
			// pop the q
			n := q[0]
			q = q[1:]
			if len(q) == 0 {
				ans = append(ans, n.Val)
			}
			if n.Left != nil {
				aq = append(aq, n.Left)
			}
			if n.Right != nil {
				aq = append(aq, n.Right)
			}
		}
		q = aq
		aq = make([]*TreeNode, 0) // 这里重新生产 empty slice 是个容易出错的地方
	}
	return ans
}


/*
我说，灵神为啥这么牛！
 */
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	var dfs func(*TreeNode, int)
	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		if depth == len(ans){
			ans = append(ans, r.Val)
		}
		dfs(r.Right, depth+1)
		dfs(r.Left, depth+1)
		return
	}

	dfs(root, 0)
	return
}