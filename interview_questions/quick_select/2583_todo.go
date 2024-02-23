package quick_select

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}

	q := []*TreeNode{root}
	l := make([]int, 0)

	for len(q) > 0 {
		tmp := q
		q = nil
		sum := 0

		for _, n := range tmp {
			sum += n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		l = append(l, sum)
	}

	if len(l) < k {
		return -1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(l)))

	//return int64(l[len(l)-k])
	return int64(l[k-1])
}
