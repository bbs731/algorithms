package binary_tree

import "slices"

/***

输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]

 */

/***
 pre order and post order 不能唯一确定一棵树。

 所以这道题应该是存在多解，的情况。题目要求，返回其中一种就可以。

我们规定，在分变不出来，子树是左孩子还是右孩子的时候， 统一按照右孩子处理。


赞啊！  这种问题，应该是难不倒你了！
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	//prei := make(map[int]int, n)
	//posi := make(map[int]int, n)

	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	// 当前 node 只有左孩子， 或者右孩子的情况。 统一按照右孩子来处理。
	if preorder[1] == postorder[n-2] {
		return &TreeNode{preorder[0], nil, constructFromPrePost(preorder[1:], postorder[:n-1])}
	}

	// now we both have left and right
	p2 := slices.Index(postorder, preorder[1])
	p1 := slices.Index(preorder, postorder[n-2])

	return &TreeNode{preorder[0], constructFromPrePost(preorder[1:p1], postorder[:p2+1]),
		constructFromPrePost(preorder[p1:], postorder[p2+1:n-1])}

}
