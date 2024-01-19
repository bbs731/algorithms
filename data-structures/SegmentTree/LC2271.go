package weekly

import "sort"

/***

给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。

同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子的长度。

请你返回使用这块毯子，最多 可以盖住多少块瓷砖。



示例 1：

输入：tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
输出：9
解释：将毯子从瓷砖 10 开始放置。
总共覆盖 9 块瓷砖，所以返回 9 。
注意可能有其他方案也可以覆盖 9 块瓷砖。
可以看出，瓷砖无法覆盖超过 9 块瓷砖。
示例 2：



输入：tiles = [[10,11],[1,1]], carpetLen = 2
输出：2
解释：将毯子从瓷砖 10 开始放置。
总共覆盖 2 块瓷砖，所以我们返回 2 。
 */

// 我觉得最常见的面试题，就是这种类型的。

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	// 对于 Leetcode
	// 不能放在 global variable
	var lazyRoot = &lazyNode{l: 1, r: 1e9, sum: stNodeDefaultVal}

	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })

	for _, t := range tiles {
		lazyRoot.update(t[0], t[1], 1)
	}
	ans := 0
	for _, t := range tiles {
		ans = max(ans, lazyRoot.query(t[0], t[0]+carpetLen-1))
	}
	return ans
}

// 动态开点线段树·其二·延迟标记（区间修改）

const stNodeDefaultTodoVal = 0

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int
	todo   int
}

func (o *lazyNode) get() int {
	if o != nil {
		return o.sum
	}
	return stNodeDefaultVal
}

func (lazyNode) op(a, b int) int {
	return a + b // max(a, b)
}

func (o *lazyNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

// 没试过， 这个build，是怎么用的
func (o *lazyNode) build(a []int, l, r int) {
	o.l, o.r = l, r
	o.todo = stNodeDefaultTodoVal
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &lazyNode{}
	o.lo.build(a, l, m)
	o.ro = &lazyNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *lazyNode) do(add int) {
	o.todo += add                  // % mod
	o.sum += (o.r - o.l + 1) * add // % mod
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == nil {
		o.lo = &lazyNode{l: o.l, r: m, sum: stNodeDefaultVal}
	}
	if o.ro == nil {
		o.ro = &lazyNode{l: m + 1, r: o.r, sum: stNodeDefaultVal}
	}
	if todo := o.todo; todo != stNodeDefaultTodoVal {
		o.lo.do(todo)
		o.ro.do(todo)
		o.todo = stNodeDefaultTodoVal
	}
}

func (o *lazyNode) update(l, r int, add int) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int {
	if o == nil || l > o.r || r < o.l {
		return stNodeDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

const stNodeDefaultVal = 0 // 如果求最大值并且有负数，改成 math.MinInt
