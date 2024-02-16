package Heap


// 这道题，太巧妙了。 模版题目。


//func kSum(nums []int, k int) int64 {
//	n := len(nums)
//	sum := 0
//	for i, x := range nums {
//		if x >= 0 {
//			sum += x
//		} else {
//			nums[i] = -x
//		}
//	}
//	sort.Ints(nums)
//	h := &hp{{0, 0}}
//	for ; k > 1; k-- {
//		p := heap.Pop(h).(pair)
//		if p.i < n {
//			heap.Push(h, pair{p.sum + nums[p.i], p.i + 1}) // 保留 nums[p.i-1]
//			if p.i > 0 {
//				heap.Push(h, pair{p.sum + nums[p.i] - nums[p.i-1], p.i + 1}) // 不保留 nums[p.i-1]，把之前减去的加回来
//			}
//		}
//	}
//	return int64(sum - (*h)[0].sum)
//}
//
//type pair struct{ sum, i int }
//type hp []pair
//
//func (h hp) Len() int            { return len(h) }
//func (h hp) Less(i, j int) bool  { return h[i].sum < h[j].sum }  // 这里是最小堆 。 用最大堆也是可以的，但是主程序的的逻辑要变。
//func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
//func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
