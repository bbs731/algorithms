package weekly

import (
	"container/heap"
	"sort"
)

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n := len(nums)
	ml := len(changeIndices) // seconds

	x := sort.Search(ml, func (m  int )bool {
		h := &hp{}
		heap.Init(h)

		all := make(map[int]struct{}, n)
		for i := m - 1; i >= 0; i-- {
			all[changeIndices[i]] = struct{}{}
			heap.Push(h, pair{changeIndices[i], i + 1})
			if len(all) == n {
				break
			}
		}
		all = make(map[int]struct{}, n)
		cost := 0
		for h.Len() > 0 {
			v := heap.Pop(h).(pair)
			if _, ok := all[v.i]; ok{
				continue
			}
			if cost+nums[v.i-1] < v.time {
				cost += nums[v.i-1]
				// mark i as completed
				all[v.i] = struct{}{}
				if len(all) == n {
					return true
				}
			} else {
				return false
			}
		}
		return false
	})
	if x > ml {
		return -1
	}
	return x
}



type pair struct {
	i int // pos i in nums to mark nums[i] as 0
	time int // the second of time, acutall the index in changIndeces.
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less (i, j int)bool {
	return h[i].time < h[j].time
}

func (h hp)Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp)Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *hp)Pop() any {
	a :=*h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
