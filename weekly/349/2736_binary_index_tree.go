package weekly

import (
	"cmp"
	"slices"
	"sort"
)

// 代码参考：
// https://leetcode.cn/problems/frequency-of-the-most-frequent-element/solutions/883643/gong-shui-san-xie-cong-mei-ju-dao-pai-xu-kxnk/

// 用 BIT 来实现 range query, 不像 开点 SegmentTree 需要对值域离散化。
// BIT 对值域需要做离散化，还是很容易出错的。
// 另外一个出错的点，BIT 的 index 是从 1 开始的，所以， newFenwickTree 的时候，size 要注意开的足够大。

const fenwickInitVal = -1

type fenwick []int

func newFenwickTree(n int) fenwick {
	t := make(fenwick, n)
	for i := range t {
		t[i] = fenwickInitVal
	}
	return t
}

func (fenwick) op(a, b int) int {
	return max(a, b)
}

// 单点更新。
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = f.op(f[i], val)
	}
}

// 求前缀和 max(a[1],  ... , a[i])
// 1<=i<=n
func (f fenwick) pre(i int) int {
	res := fenwickInitVal
	for ; i > 0; i -= i & -i {
		res = f.op(res, f[i])
	}
	return res
}

func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	// 相比于 开点 segmentTree， BIT 需要做值域的离散化
	// 需要对 a.y 和 q.y 做离散化。
	// 发现，对值域，做离散化，还是挺麻烦的， 特别容易出错。
	my := make(map[int]int)

	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
		my[nums2[i]] = 0
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Compare(b.x, a.x) })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
		my[queries[i][1]] = 0
	}
	slices.SortFunc(qid, func(i, j int) int { return cmp.Compare(queries[j][0], queries[i][0]) })

	sz := len(my)
	l := make([]int, 0, sz)
	for k := range my {
		l = append(l, k)
	}
	sort.Ints(l)
	for i, v := range l {
		my[v] = i
	}

	ans := make([]int, len(queries))
	j := 0

	fwTree := newFenwickTree(sz + 1)
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]
		for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
			fwTree.update(sz-my[a[j].y], a[j].x+a[j].y)
		}
		ans[i] = fwTree.pre(sz - my[y])
	}
	return ans
}
