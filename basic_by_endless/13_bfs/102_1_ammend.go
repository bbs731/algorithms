package bfs

func levelOrder(root *TreeNode) (ans [][]int) {

	if root == nil {
		return
	}

	// 声明加赋值
	cur := []*TreeNode{root}

	for len(cur) > 0 {

		// next and values need to be reset every iterator loop
		next := make([]*TreeNode, 0)
		values := make([]int, len(cur))

		for i, n := range cur {
			values[i] = n.Val

			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		// since values every iteration will create new one, so no need to copy here
		//ans = append(ans, append([]int(nil), values...))
		ans = append(ans, values)
		cur = next
	}

	return ans
}
