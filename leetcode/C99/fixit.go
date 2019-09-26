package C99

/**
 * definition for a binary tree node.
 * type treenode struct {
 *     val int
 *     left *treenode
 *     right *treenode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Space Complexity is O(1), we have Algorithm with Space Complexity O(n) but with easy implementation
var Prev *TreeNode
var Found1 *TreeNode
var Found2 *TreeNode

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	if Prev != nil {
		if Prev.Val > root.Val {
			if Found1 == nil {
				Found1 = Prev
				Found2 = root // treat the special case [3, 1, 4, null, null, 2]
			} else {
				if Found1.Val > Prev.Val {
					Found2 = root
				} else {
					Found2 = Prev
				}
				return
			}
		}
	}
	Prev = root
	InOrderTraversal(root.Right)
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	Prev = nil
	Found1 = nil
	Found2 = nil

	InOrderTraversal(root)

	tmp := Found1.Val
	Found1.Val = Found2.Val
	Found2.Val = tmp
}
