package interview_questions


/***
太简单，但是有陷阱
 */
func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	ans := 0
	half := 0

	for _, num := range nums {
		m[num]++
	}

	for n1, v1:= range m {
		if  v2, ok := m[k-n1]; ok {
			if k-n1 == n1 {
				half += v1/2
			} else {
				ans += min(v1, v2)  // 总共被加了2遍
			}
		}
	}
	return ans/2 + half
}
