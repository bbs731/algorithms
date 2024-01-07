package weekly

func maximumSetSize(nums1 []int, nums2 []int) int {
	n := len(nums1)
	h1 := make(map[int]bool)
	h2 := make(map[int]bool)
	h := make(map[int]bool)

	for i:=0; i<n; i++ {
		h1[nums1[i]] = true
		h2[nums2[i]] = true
	}

	ans := 0
	cnt:=0
	for x := range h1{
		if _, ok := h2[x]; !ok {
			cnt++
		} else {
				h[x] = true
		}
		if cnt >= n/2 {
			break
		}
	}
	diff1 := cnt
	cnt = 0
	for x := range h2{
		if _, ok :=h1[x];!ok {
			cnt++
		} else {
				h[x] = true
		}
		if cnt == n/2 {
			break
		}
	}
	diff2 :=cnt
	ans = diff1 + diff2 + min(n-diff1-diff2, len(h))

	return ans
}
