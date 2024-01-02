package weekly

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	ans := make([]int64, len(queries))
	n := len(nums)
	presum := make([]int, n+1)
	for i := 1; i<=n; i++ {
		presum[i]= presum[i-1]+nums[i-1]
	}

	for i, q := range queries {
		pos := sort.SearchInts(nums, q)  // 你还是不了解二分， SearchInts 返回的index是什么，理解错误！
		//if pos == n  {
		//	ans[i] = int64(n*q - (presum[n]-presum[0]))
		//} else if pos == 0 {
		//	ans[i] = int64(presum[n]-presum[0] - n*q)
		//} else{
		//	part1 := pos*q - (presum[pos] - presum[0])
		//	part2 := presum[n] - presum[pos] - (n-pos)*q
		//	ans[i] = int64(part1 + part2)
		//}
		// 可以简化成这样！
		part1 := pos*q - (presum[pos] - presum[0])
		part2 := presum[n] - presum[pos] - (n-pos)*q
		ans[i] = int64(part1 + part2)
	}
	return ans
}
