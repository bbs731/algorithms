package weekly

import "sort"

func minimumPushes(word string) int {
	ans := 0
	tick := 0

	m :=make(map[int32]int)
	for _, w := range word {
		m[w]++
	}
	nums := make([]int, 0, len(m))
	for _, v := range m {
		nums = append(nums, v)
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] > nums[j]})

	for i, v := range nums{
		if i % 8 == 0 {
			tick++
		}
		ans += v*tick
	}
	return ans
}
