package monotonic_stack

import "sort"

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	stack1 := make([]int,0)
	stack2 := make([]int, 0)

	for index, x := range nums {
		n2 := sort.Search(len(stack2), func(i int) bool { return nums[stack2[i]] >= x})
		for j :=0; j<n2; j++ {
			ans[stack2[j]] = x
		}
		// pop stack2 and insert x
		stack2 = stack2[n2:]

		n1 := sort.Search(len(stack1), func (i int) bool { return nums[stack1[i]] >= x})
		for j:=0; j<n1; j++ {
			// move candidate from stack1 to stack2
			candidate := nums[stack1[j]]
			pos := sort.Search(len(stack2), func(i int) bool {return nums[stack2[i]] >= candidate})
			if pos == 0 {
				stack2 = append([]int{stack1[j]}, stack2[:]...)
			} else {
				stack2 = append(stack2[:pos], append([]int{stack1[j]}, stack2[pos:]...)...)
			}
		}
		// pop stack1
		stack1 = stack1[n1:]
		stack1 = append([]int{index}, stack1[:]...)
		//stack2 = append([]int{x}, stack2[n2:]...)
	}

	for j :=0; j < len(stack1); j++ {
		ans[stack1[j]]	= -1
	}
	for j :=0; j<len(stack2); j++ {
		ans[stack2[j]]= -1
	}
	return ans
}
