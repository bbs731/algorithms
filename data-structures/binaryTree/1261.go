package binaryTree


/***
看看，灵神的答案，惊艳啊！
 */
type FindElements map[int]bool

func Constructor(root *TreeNode) FindElements {
	f := FindElements{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		f[val] = true
		dfs(node.Left, val*2+1)
		dfs(node.Right, val*2+2)
	}
	dfs(root, 0)
	return f
}

func (f FindElements) Find(target int) bool {
	return f[target]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	root *TreeNode
}


func Constructor(root *TreeNode) FindElements {
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			r.Left.Val = 2*r.Val+1
			dfs(r.Left)
		}
		if r.Right != nil {
			r.Right.Val = 2*r.Val +2
			dfs(r.Right)
		}
	}
	if root != nil {
		root.Val = 0
		dfs(root)
	}
	return FindElements{
		root:root,
	}
}


func (this *FindElements) Find(target int) bool {
	var dfs  func(*TreeNode) bool
	dfs = func(r *TreeNode) bool {
		if r == nil {
			return false
		}
		if r.Val == target {
			return true
		}
		return dfs(r.Left) || dfs(r.Right)
	}
	return dfs(this.root)
}




