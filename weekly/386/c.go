package weekly

import "container/heap"

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

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n := len(nums)
	m := len(changeIndices) // seconds

	h := &hp{}
	heap.Init(h)

	all := make(map[int]struct{}, n)
	for i :=m-1;i>=0; i--{
		all[changeIndices[i]]= struct{}{}
		heap.Push(h, pair{ changeIndices[i],i+1})
		if len(all)== n {
			break
		}
	}

	all = make(map[int]struct{}, n)
	cost := 0
	for h.Len() > 0 {
		v := heap.Pop(h).(pair)
		if cost + nums[v.i] < v.time {
			cost += nums[v.i-1]
			// mark i as completed
			all[v.i] = struct{}{}
			if len(all) == n {
				if cost +1 == v.time {
					return v.time +1
				}
				return v.time
			}
		} else {
			break // failed
			//return -1
		}
	}
	return -1
}
