package binary_tree

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


func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder)== 0 {
		return nil
	}
	n := len(postorder)
	p := slices.Index(inorder, postorder[n-1])

	l := buildTree(inorder[:p], postorder[:p])
	r := buildTree(inorder[p+1:], postorder[p:n-1])  // 这个 index 需要特别的小心
	return &TreeNode{postorder[n-1], l, r}
}


func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	index := make(map[int]int, n)
	for i, x := range inorder {
		index[x] = i
	}

	var dfs func([]int, []int)*TreeNode
	dfs = func(inorder []int, postorder []int) *TreeNode {
		n := len(postorder)
		if n == 0 {
			return nil
		}
		p := index[postorder[n-1]]

		l := buildTree(inorder[:p], postorder[:p])
		r := buildTree(inorder[p+1:], postorder[p:n-1])  // 这个 index 需要特别的小心
		return &TreeNode{postorder[n-1], l, r}
	}
	return dfs(inorder, postorder)
}

