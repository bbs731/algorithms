package _76

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	valid := true
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			valid = false
			break
		}
		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}

	if !valid {
		return [][]int{}
	}
	return ans
}
