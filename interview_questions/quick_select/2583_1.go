package quick_select

import "math/rand"

/***
默写模板， 默写模板， 默写模板
 */

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	level := make([]int, 0)

	for len(q) > 0 {
		tmp := q
		q = nil
		sum := 0
		for _, n := range tmp {
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			sum += n.Val
		}
		level = append(level, sum)
	}

	// need to find the kth largest value of l
	if len(level) < k {
		return -1
	}
	n := len(level)
	// find n -k
	l, r := 0, n-1
	k = n - k

	rand.Shuffle(n, func(i, j int) {
		level[i], level[j] = level[j], level[i]
	})

	for l < r {
		i, j := l, r+1
		v := level[l]
		for {
			for i++; i < r && level[i] < v; i++ {
			}
			for j--; j > l && level[j] > v; j-- {
			}
			if i >= j {
				break
			}
			level[i], level[j] = level[j], level[i]
		}

		// swap l and j
		level[l], level[j] = level[j], level[l]
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return int64(level[k])
}
