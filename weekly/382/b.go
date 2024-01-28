package weekly

import "sort"

func maximumLength(nums []int) int {
	sort.Ints(nums)
	//n := len(nums)
	m := make(map[int]int)
	ans := 1

	for _, v := range nums{
		m[v]++
	}

	for i := range nums {
		if nums[i] == 1 || m[nums[i]] == 1 {
			continue
		}
		l := 2
		x := nums[i]
		for m[x*x]!=0 {
			x = x*x
			l = l+2
			if m[x] < 2{
				break
			}
		}
		ans = max(ans, l-1)
	}

	// 数字 1 特殊。
	if m[1] > 0 {
		if m[1]&1 !=0  {
			ans = max(ans, m[1])
		}  else {
			ans = max(ans, m[1]-1)
		}
	}

	return ans

}
