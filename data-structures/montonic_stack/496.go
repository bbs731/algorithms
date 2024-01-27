package montonic_stack

/***
单调栈的模版题目。  "下一个更大的元素"
 */
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	st := []int{}
	m := make(map[int]int)
	ans := make([]int, n)
	for i:= range ans {
		ans[i] = -1
	}

	for i, v := range nums2 {
		m[v] = i
		for len(st) > 0 && v > nums2[st[len(st)-1]]{
			j := st[len(st)-1]
			ans[j] = nums2[i]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	res := make([]int, len(nums1))
	for i, v:= range nums1 {
		res[i] = ans[m[v]]
	}
	return res
}
