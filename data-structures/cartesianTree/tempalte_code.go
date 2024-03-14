package cartesianTree

type ctNode struct {
	lr      [2]*ctNode
	id, val int
}

/***
Treap 是 笛卡尔树的一种实现， 其中 treap 的 val 是随机的。


灵神笛卡尔树的板子来自：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/cartesian_tree.go
 */

func buildCartesianTree(a []int) *ctNode {
	if len(a) == 0 {
		return nil
	}

	s := []*ctNode{}
	for i, v := range a {
		o := &ctNode{id: i, val: v}
		for len(s) > 0 {
			top := s[len(s)-1]
			if top.val < v {
				top.lr[1] = o
				break
			}
			o.lr[0] = top
			s = s[:len(s)-1]
		}
		s = append(s, o)
	}
	return s[0]
}

// 非指针的版本
func buildCartesianTree2(a []int) [][2]int {
	n := len(a)
	lr := make([][2]int, n)
	for i := range lr {
		lr[i] = [2]int{-1, -1}
	}

	s := []int{}
	for i, v := range a {
		for len(s) > 0 {
			top := s[len(s)-1]
			if a[top] < v {
				lr[top][1] = i
				break
			}
			lr[i][0] = top
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return lr
}
