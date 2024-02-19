package Heap

import "container/heap"

/***

对于一个序列， 找到， 和 第 kth小的子序列， 这个灵神教过。


对于两个序列， 就是本题。 经典的题目。

用到了 heap 最小堆。 当 pair {i, j } 被 pop 出来的时候，
把  {i+1, j}  和 {i, j+1} 加入 Heap 中， 但是有一个问题就是 {i+1, j} 和 {i， j+1} 会导致 {i+1, j+1} 入堆两次。 为了不额外记录，pair 是否已经如果堆的信息，
决定只让 {i+1, j} 在 {i, j} 出堆的时候入堆， 前提是 {0, j} 需要先入堆, 作为初始条件 。

一旦思路是对的， 代码很难再写错了啊
 */

type pair struct{ add, i, j int }
type hp []pair

func (h hp) Len() int {
	return len(h)
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hp) Less(i, j int) bool {
	return h[i].add < h[j].add // 最小堆
}

func (h *hp) Pop() (v any) {
	a := *h
	v, *h = a[len(a)-1], a[:len(a)-1]
	return
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	n := len(nums1)
	m := len(nums2)

	h := make(hp, 0)
	heap.Init(&h)

	for j := 0; j < m; j++ {
		heap.Push(&h, pair{nums1[0] + nums2[j], 0, j})
	}

	for k > 0 {
		p := heap.Pop(&h).(pair)
		ans = append(ans, []int{nums1[p.i], nums2[p.j]})
		if p.i+1 < n { // 这里，竟然考虑到了，还犯了一个错误。
			heap.Push(&h, pair{nums1[p.i+1] + nums2[p.j], p.i + 1, p.j})
		}
		k--
	}
	return ans
}
