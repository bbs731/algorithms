package interview_questions

import "sort"

/***
这玩仍， 能背下来吗？
 */
func merge(intervals [][]int) [][]int {
	st := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for _, i := range intervals {
		s, e := i[0], i[1]
		// 这段逻辑，可以为下面处理， 所以，其实是可以省略的。
		//if len(st) == 0 || s > st[len(st)-1][1] {
		//	st = append(st, []int{s, e})
		//	continue
		//}

		// do the merge
		for len(st) > 0 {
			if s < st[len(st)-1][0] {
				// pop stack
				st = st[:len(st)-1]
			} else {
				if s <= st[len(st)-1][1] {
					s = min(st[len(st)-1][0], s)
					st = st[:len(st)-1]
				}
				break
			}
		}
		st = append(st, []int{s, e})
	}
	return st
}
