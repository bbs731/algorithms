package interview_questions

// find the different digit of the two missing number
func missingTwo(nums []int) []int {
	s := 0
	n := len(nums)
	for i:=1; i<=n+2; i++ {
		s = s^i
	}
	for _, d := range nums {
		s = s^d
	}

	dd := s &(-s)  // 这个 dd 说明， 这两个 missing number 在这一个位置上是不同的。 可以用这一个 bit ,把 1..n 个数分成两个 group

	s1, s2 := 0, 0
	for i:=1; i<=n+2; i++ {
		if i & dd == 0 {
			s1 = s1^i
		} else {
			s2 = s2^i
		}
	}
	for _, d := range nums {
		if d &dd == 0 {
			s1 = s1^d
		}else {
			s2 = s2^d
		}
	}
	return []int{s1, s2}
}
