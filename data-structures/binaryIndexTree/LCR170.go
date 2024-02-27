package binaryIndexTree

import "sort"

/***
在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。

示例 1:
输入：record = [9, 7, 5, 4, 6]
输出：8
解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。

限制：
0 <= record.length <= 50000

 */

type BIT struct {
	n int
	b []int // index from 1 to n
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x, v int) {
	for x <= b.n {
		b.b[x] += v
		x += lowbit(x)
	}
}

func (b *BIT) query(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.b[x]
		x -= lowbit(x)
	}
	return ans
}

func reversePairs(record []int) int {
	n := len(record)
	type pair struct {
		val int
		i   int
	}
	// 离散化值域
	// 自然就会处理， 重复的元素。重复的值， 离散化，之后， 就是1个值，就占有一个 index 位置。
	// 这个离散化，是不是写的太长了？
	m := make(map[int]struct{}, n)
	b := make([]int, 0, n)
	for _, r := range record {
		if _, ok := m[r]; !ok {
			m[r] = struct{}{}
			b = append(b, r)
		}
	}
	sort.Ints(b)
	mi := make(map[int]int)
	for i := 0; i < len(b); i++ {
		mi[b[i]] = i
	}

	bit := &BIT{
		len(mi),
		make([]int, len(mi)+1),
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		x := record[i]
		ans += bit.query(mi[x])
		bit.add(mi[x]+1, 1)
	}
	return ans
}
