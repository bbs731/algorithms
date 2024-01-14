package weekly

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	cnt := 0
	for _, x:= range nums {
		m[x]++
		cnt = max(cnt, m[x])
	}

	ans := 0
	for _, v := range m {
		if v == cnt {
			ans += v
		}
	}
	return ans
}
