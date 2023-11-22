package bfs

func levelOrder(root *TreeNode) (ans [][]int) {

	if root == nil {
		return
	}

	// 和第一种的区别就是，只使用了一个数组 (即队列)
	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		values := make([]int, l)

		for i := range values {
			// pop deque
			n := q[0]
			q = q[1:]
			//values = append(values, n.Val)
			values[i] = n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		ans = append(ans, values)
	}
	return ans
}
