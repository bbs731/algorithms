package weekly

import (
	"cmp"
	"slices"
	"sort"
)

// 标准， 最基本的框架。 技巧也是最难的。

// https://leetcode.cn/problems/maximum-sum-queries/solutions/2305395/pai-xu-dan-diao-zhan-shang-er-fen-by-end-of9h/
// 来自灵神， 单调栈的解法。
// 不使用更高级的数据结构。 1. 需要更高的技巧  2. 思考过程比较难,比较难以想到！
func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Compare(b.x, a.x) })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}
	slices.SortFunc(qid, func(i, j int) int { return cmp.Compare(queries[j][0], queries[i][0]) })

	ans := make([]int, len(queries))
	type data struct{ y, s int }
	st := []data{} // 单调栈， 按照 y 的值是递增的，  按照 x+y 的sum值是递减的。
	j := 0

	// 倒排 query, 倒排 pairs 数组，也是难以想到。 但是这个思路是可以被训练出来的，加油！
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]
		for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
			for len(st) > 0 && st[len(st)-1].s <= a[j].x+a[j].y { // a[j].y >= st[len(st)-1].y
				st = st[:len(st)-1]
			}
			if len(st) == 0 || st[len(st)-1].y < a[j].y {
				st = append(st, data{a[j].y, a[j].x + a[j].y})
			}
		}

		// 因为 是单调栈， 按照 y 递增的，所以，才可以二分并满足时间复杂度。
		p := sort.Search(len(st), func(i int) bool { return st[i].y >= y })
		if p < len(st) {
			ans[i] = st[p].s
		} else {
			ans[i] = -1
		}
	}
	return ans
}
