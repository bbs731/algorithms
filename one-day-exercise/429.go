package one_day_exercise


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return  nil
	}
	q := []*Node{root}
	ans := make([][]int, 0)

	for len(q) > 0 {
		tmp := q
		q = nil

		level := make([]int, 0, len(tmp))
		for _, n := range tmp {
			level = append(level, n.Val)
			q = append(q, n.Children...)
		}
		ans = append(ans, level)
	}
	return ans
}
