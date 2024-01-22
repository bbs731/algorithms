package weekly

import "sort"

/***
看看，你自己的提交记录， 写的越来越简洁了 （偶尔发挥好点，哈哈)

sliding-window O(n) 的， 但是受限于， sort 所以总体复杂度是 O(n*logn)

比较不同的写法，发现，大多数的情况下， 还是枚举右端点，移动左端点的写法，比较简洁
 */

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	ans := 1
	sort.Ints(nums)

	// 前缀和，总不写，还是容易出错啊， psum[0]=0 index 从  1开始。
	psum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + nums[i-1]
	}

	left := 0
	for i := 1; i < n; i++ {
		diff := (i-left)*nums[i] - (psum[i] - psum[left])
		for diff > k {
			left++
			diff = (i-left)*nums[i] - (psum[i] - psum[left])
		}

		ans = max(ans, i-left+1)
	}
	return ans
}
