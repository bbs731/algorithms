package dp

import (
	"fmt"
	"sort"
)
/*
现在的水平，感觉就是，题目有坑一定会掉进去，有几个，掉进去几次。
 */
type indexSlice struct {
	index  []int
	nums   []int
	scores []int
}

func (x *indexSlice) Len() int { return len(x.index) }
func (x *indexSlice) Less(i, j int) bool {
	if x.nums[i] < x.nums[j] {
		return true
	}
	if x.nums[i] == x.nums[j] {
		return x.scores[i] < x.scores[j]
	}
	return false
}
func (x *indexSlice) Swap(i, j int) {
	x.index[i], x.index[j] = x.index[j], x.index[i]
	// 这里， nums, 和 scores item 需要互换位置， 一个坑。
	x.nums[i], x.nums[j] = x.nums[j], x.nums[i]
	x.scores[i], x.scores[j] = x.scores[j], x.scores[i]
}

//func LIS(nums []int) int {
//	ans := 0
//	sum := make([]int, len(nums))
//	g := []int{}
//
//	for _, x := range nums {
//		pos := sort.SearchInts(g, x)
//		if pos == len(g) {
//			g = append(g, x)
//		} else {
//			g[pos] = x
//		}
//		prevsum := 0
//		if pos > 0 {
//			prevsum = sum[pos-1]
//		}
// 这里也是一个坑
//		sum[pos] = prevsum + x   // 这样写有坑啊， 这里没办法处理，怎么写 sum[pos] 都是错的。 考虑一下
//[8 2 9]
//[5 1 2 3]
// 没有办法处理。
//		ans = max(ans, sum[pos])
//	}
//	return ans
//}

func LIS(nums []int) int {
	n := len(nums)
	//f := make([]int, n)
	sum := make([]int, n)
	ans := 0

	for i, x := range nums {
		sum[i] = x
		for j := 0; j < i; j++ {
			if x >= nums[j] {
				//f[i] = max(f[i], f[j]+1)
				sum[i] = max(sum[i], sum[j]+x)
			}
		}
		ans = max(ans, sum[i])
	}
	return ans
}

func bestTeamScore(scores []int, ages []int) int {

	n := len(ages)
	index := make([]int, n)
	for i := range ages {
		index[i] = i
	}
	s := &indexSlice{index, ages, scores}
	// 根据
	sort.Sort(s)
	return LIS(s.scores)
}