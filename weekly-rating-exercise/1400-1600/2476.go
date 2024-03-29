package _400_1600

import (
	"slices"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 my goodness, you are good!
 边界条件，太容易错了。

 https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/solutions/1981349/zhong-xu-bian-li-er-fen-cha-zhao-by-endl-m8ez/
 灵神的答案，思路是一样的, 赞!
  */
func closestNodes(root *TreeNode, queries []int) [][]int {
	tree := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func( r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		tree = append(tree, r.Val)
		dfs(r.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, x:=range queries {
		ans[i] = make([]int, 2)
		// >=x 的最小值
		lower, _ := slices.BinarySearch(tree, x)
		//lower := sort.SearchInts(tree, x)  // 这道题，考的就是二分查找。
		// <=x 的最大值
		upper := sort.SearchInts(tree, x+1) -1

		if upper <0  {  // 这里的边界条件，太容易错了。
			ans[i][0] = -1
		} else {
			ans[i][0] = tree[upper]
		}

		if lower >= len(tree) {  // 这里的边界条件
			ans[i][1] = -1
		} else {
			ans[i][1] = tree[lower]
		}
	}
	return ans
}
