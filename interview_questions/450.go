package interview_questions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		/****
		 这里的逻辑，其实是可以简化的。判断，左边和右边是否为空，单独处理返回就行。 如果，左右都不是空的话， 选择右边即可。
		 看，下面的实现吧。
		 */
		var p, c *TreeNode
		p = root
		if root.Right != nil {
			c = p.Right
			for c.Left != nil {
				p = c
				c = c.Left
			}

		} else {
			c = p.Left
			for c.Right != nil {
				p = c
				c = c.Right
			}
		}
		root.Val = c.Val
		if c == p.Right {
			p.Right = deleteNode(c, c.Val)
		} else {
			p.Left = deleteNode(c, c.Val)
		}

	} else {
		if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		} else {
			root.Left = deleteNode(root.Left, key)
		}
	}

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		//var p, c *TreeNode
		// 选择右边子树即可。
		var p, c *TreeNode = root, root.Right
		for c.Left != nil {
			p = c
			c = c.Left
		}
		_ = p
		root.Val = c.Val
		// 这个技巧，其实不怎么好，因为需要，再次 traverse root.Right, 有更好的做法。
		root.Right = deleteNode(root.Right, c.Val)

		//if c == p.Right {
		//	p.Right = deleteNode(c, c.Val)
		//} else {
		//	p.Left = deleteNode(c, c.Val)
		//}

	} else {
		if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		} else {
			root.Left = deleteNode(root.Left, key)
		}
	}

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		//var p, c *TreeNode
		// 选择右边子树即可。
		var p, c *TreeNode = root, root.Right
		for c.Left != nil {
			p = c
			c = c.Left
		}
		// 可以让 c 成为新的 root, 但是有一个问题， 需要去掉root，如何做，其实有简单的做法，返回 root.Right 作为新的 root 替换掉 root 就想到与 删除 root 了。
		// 但是这个技巧不是非常的好想吧！
		c.Left = root.Left
		return root.Right
		//_ = p
		//root.Val = c.Val
		//root.Right = deleteNode(root.Right, c.Val)
		//if c == p.Right {
		//	p.Right = deleteNode(c, c.Val)
		//} else {
		//	p.Left = deleteNode(c, c.Val)
		//}

	} else {
		if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		} else {
			root.Left = deleteNode(root.Left, key)
		}
	}

	return root
}
