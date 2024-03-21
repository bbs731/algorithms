package dp

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	s1, s2 := 0, 0
	n := len(nums1)
	for i := range nums1 {
		s1 += nums1[i]
		s2 += nums2[i]
	}
	//
	//if s1 < s2 {
	//	nums1, nums2 = nums2, nums1
	//	s1, s2 = s2, s1
	//}

	// we have nums1 >= nums2 and s1 >= s2
	// what  the fuck, 正反，都需要试，这个没想到。 所以根本，就不用比较 s1 和 s2 的大小
	l := make([]int, n)
	r := make([]int, n)
	for i := range l {
		l[i] = nums2[i] - nums1[i]
		r[i] = nums1[i] - nums2[i]
	}

	f0 := 0
	f1 := 0
	f := 0
	fr := 0
	for i := 1; i <= n; i++ {
		f0 = max(f0+l[i-1], l[i-1], 0)
		f = max(f, f0)

		f1 = max(f1+r[i-1], r[i-1], 0)
		fr = max(fr, f1)
	}
	return max(f+s1, fr+s2)
}
