package _600_1900

import (
	"container/heap"
	"sort"
)

/*
这道题，非常好的一道题，
证明了两个事情：

1.  你不会写 golang 的 heap
2.  面试题，就是这种难度。 你需要锻炼这种难度的题。不要为了挑战高难度，去学很难的知识点。 先把 1700 -1900 这段给秒了。

果然一道 1700 就会卡死你
 */

// 事件扫描线+堆

// golang 里面如何写一个 heap ？

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	unoccupied := hp{make([]int, n)}
	for i := 0; i < n; i++ {
		unoccupied.IntSlice[i] = i
	}
	belong := make([]int, n)

	events := make([][2][]int, 1e5+1) // 这个初始化太难了。
	for i, t := range times {
		events[t[0]][1] = append(events[t[0]][1], i) // 朋友来到
		events[t[1]][0] = append(events[t[1]][0], i) // 朋友离开
	}

	for _, e := range events {
		for _, f := range e[0] {
			heap.Push(&unoccupied, belong[f]) // return the chair   // 你说 golang 的 heap 坑不坑？
		}

		for _, f := range e[1] {
			c := heap.Pop(&unoccupied)
			belong[f] = c.(int)
			if f == targetFriend {
				return belong[f]
			}
		}
	}

	return -1
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

// 这个想法为什么是错的？ 好吧， 你的 chair 的数字从来没减少过， 怎么可能是对的呢？
func smallestChair(times [][]int, targetFriend int) int {
	tables := make(map[int]int)
	index := make(map[int]int) // start time to index mapping
	n := len(times)

	//l := make(recordSlice, 0, n)
	l := make([]int, 0, n)
	for i, t := range times {
		l = append(l, t[0])
		index[t[0]] = i
	}
	sort.Ints(l)
	chair := 0
	for true {
		friend := index[l[0]]
		tables[friend] = chair
		if friend == targetFriend {
			return chair
		}
		//for true {
		// fill up all time intervals that use this chair
		pos := sort.SearchInts(l, times[friend][1]+1) - 1 // for index which  >= friend.endtime
		if l[pos] >= times[friend][1] {
			tables[pos] = chair // pos friend take this chair as well
			friend = pos
		} else {
			pos = pos + 1
			if pos < n {
				tables[pos] = chair
				friend = pos
			}
			//break
		}
		//}
		if friend == targetFriend {
			return chair
		}

		l = make([]int, 0)
		for i, t := range times {
			if _, ok := tables[i]; !ok {
				l = append(l, t[0])
			}
		}
		if len(l) == 0 {
			panic("impossible")
		}
		sort.Ints(l)
		chair++
	}
	return -1
}
