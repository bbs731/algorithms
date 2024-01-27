package loop

import "sort"


// 按照模版写, 赞赞赞！
func reductionOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for i:=0; i<n; {
		start := i
		for ; i<n && nums[i]== nums[start]; i++ {
		}
		ans += n-i // n-1 -i + 1
	}
	return ans
}
