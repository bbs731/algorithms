package dp

import "sort"

/*
可以用， 单调栈， 把时间复杂度，下降到 O(n*logn)
不会
 */
func lengthOfLIS(nums []int) int {
	//n := len(nums)
	//
	//if n <= 1{
	//	return n
	//}
	//
	//q := []int{0}
	////q := make([]int, n+1)
	//
	//for i:=1; i<n; i++ {
	//	if nums[i] > nums[q[len(q)-1]] {
	//		q = append(q, i)
	//		continue
	//	}
	//
	//
	//	//
	//	//for len(q)>0 && nums[q[len(q)-1]] >= x {
	//	//	q = q[:len(q)-1] // pop last element
	//	//}
	//	//q = append(q, i)
	//}
	//return len(q)
}

/*
 翻译成为 DP
 	f[i] = max(f[j]) + 1 // for j < i and nums[i] > nums[j]   看到一个连续区间，求 max 是不是想到了可以用单调栈，单调队列优化？

在没有任何优化的情况下：
时间复杂度是 O(n^2)
 */
func lengthOfLIS_dp(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	ans := 1

	for i, x :=range nums{
		f[i] = 1
		for j:=0; j<i; j++ {
			if x > nums[j] {
				f[i] = max(f[i], f[j] + 1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}


/*
	f[i] 是 包含 i 的， 最长升序的 长度。
	f[i] = max(f[j]） + 1   // for nums[i] > nums[j]
 */
func lengthOfLIS_dfs(nums []int) int {
	n :=len(nums)
	cache := make([]int, n)
	for i := range nums {
		cache[i] = -1
	}

	var dfs func(int)int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}
		ans := 1
		for j:=0; j<i; j++ {
			if nums[i]> nums[j]{
				ans = max(ans, dfs(j) + 1)
			}
		}
		cache[i] = ans
		return ans
	}

	res := 1
	for i:=0; i<n; i++ {
		res = max(res, dfs(i))
	}
	return res
}
