package dp

import "sort"

/*
想法来源 LC300
 */
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	f := make([]int, 0)
	ans := []int{}

	for _, x := range obstacles {
		pos := sort.SearchInts(f, x+1) // 这里找的是 upper bound
		if pos == len(f) {
			f = append(f, x)
		} else {
			f[pos] = x
		}
		ans = append(ans, pos+1)
	}
	return ans
}

func longestObstacleCourseAtEachPosition_dp(obstacles []int) []int {
	n := len(obstacles)
	f := make([]int, n)
	ans := []int{}

	for i, x := range obstacles {
		longest := 1
		for j := 0; j < i; j++ {
			if x >= obstacles[j] {
				longest = max(longest, f[j]+1)
			}
		}
		ans = append(ans, longest)
	}
	return ans
}
