package binary_search

import (
	"container/heap"
	"sort"
)

/***

给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。



示例 1:

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:

输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
示例 3:

输入: nums1 = [1,2], nums2 = [3], k = 3
输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]


提示:

1 <= nums1.length, nums2.length <= 10^5
-10^9 <= nums1[i], nums2[i] <= 10^9
nums1 和 nums2 均为 升序排列
1 <= k <= 10^4
k <= nums1.length * nums2.length

 */

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n := len(nums1)
	m := len(nums2)
	ml, mr := nums1[0]+nums2[0], nums1[n-1]+nums2[m-1]

	// 答案上二分  是一个 先 false, 后 true 的正常序列。
	l := ml - 1
	r := mr + 1
	for l+1 < r {
		mid := (l + r) >> 1
		tot := 0
		for _, v1 := range nums1 {
			// search for  >= mid-v1
			p := sort.SearchInts(nums2, mid-v1+1) - 1
			tot += p + 1
		}

		if tot < k {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 == r
	//return r  因为，最后的结果不是要 r 而是要，最准确的 element pairs 所以，应该最开始就用 heap
	//搞不懂啊！

	ans := make([][]int, 0, k)
	//for _, v1 := range nums1 {
	//	for _, v2 := range nums2 {
	//		if v1+v2 > r {
	//			break
	//		}
	//		ans = append(ans, []int{v1, v2})
	//	}
	//	if len(ans) > k {
	//		break
	//	}
	//}
	//sort.Slice(ans, func(i, j int) bool {
	//	if ans[i][0]+ans[i][1] < ans[j][0]+ans[j][1] {
	//		return true
	//	}
	//	if ans[i][0]+ans[i][1] == ans[j][0]+ans[j][1] {
	//
	//	}
	//})
	return ans[:k]
}

/***

找机会专门练习一下，堆栈的操作！

 */
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	ans := make([][]int, 0, min(k, n*m)) // 预分配空间
	h := make(hp, min(k, n))
	for i := range h {
		h[i] = tuple{nums1[i] + nums2[0], i, 0}
	}
	for len(h) > 0 && len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
