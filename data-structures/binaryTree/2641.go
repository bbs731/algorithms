package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/***
你这么写， 过不去面试啊。 题如果，太简单了， 往往要求题目的解，是最优的，这个比较要命。
 */

func replaceValueInTree(root *TreeNode) *TreeNode {
	st := []*TreeNode{root}
	for len(st) > 0 {
		tmp := []*TreeNode{}
		total_next_level := 0
		for _, n := range st {
			if n.Left == nil && n.Right == nil {
				continue
			}
			if n.Left != nil {
				total_next_level += n.Left.Val
			}
			if n.Right != nil {
				total_next_level += n.Right.Val
			}
		}
		for _, n := range st {
			siblings := 0
			if n.Left != nil {
				siblings += n.Left.Val
			}
			if n.Right != nil {
				siblings += n.Right.Val
			}
			if n.Left != nil {
				n.Left.Val = total_next_level - siblings
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				n.Right.Val = total_next_level - siblings
				tmp = append(tmp, n.Right)
			}
		}
		st = tmp
		tmp = nil // clear tmp
	}

	root.Val = 0
	return root
}

type content struct {
	index int
	nums  []int
	total int // sum of nums
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	saved := make(map[*TreeNode]*content)

	st := []*TreeNode{root}
	for len(st) > 0 {
		tmp := []*TreeNode{}
		nums := []int{}
		index := 0
		total := 0
		for _, n := range st {
			if n.Left == nil && n.Right == nil {
				continue
			}
			lv, rv := 0, 0
			if n.Left != nil {
				lv = n.Left.Val
				saved[n.Left] = &content{index, nil, 0}
			}
			if n.Right != nil {
				rv = n.Right.Val
				saved[n.Right] = &content{index, nil, 0}
			}
			index++
			nums = append(nums, lv+rv)
			total += lv + rv
		}
		for _, n := range st {
			if n.Left != nil {
				saved[n.Left].nums = nums
				saved[n.Left].total = total
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				saved[n.Right].nums = nums
				saved[n.Right].total = total
				tmp = append(tmp, n.Right)
			}
		}
		st = tmp
		tmp = nil // clear tmp
	}

	// second round of BFS
	root.Val = 0
	st = []*TreeNode{root}

	for len(st) > 0 {
		tmp := []*TreeNode{}
		for _, n := range st {
			if n.Left != nil {
				if l, ok := saved[n.Left]; ok {
					n.Left.Val = l.total - l.nums[l.index]
				}
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				if r, ok := saved[n.Right]; ok {
					n.Right.Val = r.total - r.nums[r.index]
				}
				tmp = append(tmp, n.Right)
			}
		}
		st = tmp
		tmp = nil
	}
	return root
}
