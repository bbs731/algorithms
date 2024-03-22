package one_day_exercise

import (
	"container/heap"
	"math"
)

type pair struct {
	c, i, j int
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].c < h[j].c }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

/***

感觉很简单的样子， 还是参考了灵神的思路。
https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/solutions/2216329/dan-diao-zhan-you-hua-dp-by-endlesscheng-mc50/

 */
func minimumVisitedCells(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt32
		}
	}

	colh := make([]hp, n)
	for i := range colh {
		colh[i] = hp{}
	}

	heap.Push(&colh[0], pair{1, 0, 0})
	costs[0][0] = 1
	for i := range grid {
		// new row
		rowh := hp{}
		for j := range grid[i] {
			for len(rowh) > 0 && grid[rowh[0].i][rowh[0].j]+rowh[0].j < j {
				// pop heap
				heap.Pop(&rowh)
			}
			if len(rowh) > 0 {
				costs[i][j] = min(costs[i][j], rowh[0].c+1)
			}

			// check column
			for len(colh[j]) > 0 && grid[colh[j][0].i][colh[j][0].j]+colh[j][0].i < i {
				heap.Pop(&colh[j])
			}
			if len(colh[j]) > 0 {
				costs[i][j] = min(costs[i][j], colh[j][0].c+1)
			}
			if costs[i][j] != math.MaxInt32 {
				heap.Push(&rowh, pair{costs[i][j], i, j})
				heap.Push(&colh[j], pair{costs[i][j], i, j})
			}
		}
	}

	//fmt.Println(costs)
	if costs[m-1][n-1] == math.MaxInt32 {
		return -1
	}
	return costs[m-1][n-1]
}
