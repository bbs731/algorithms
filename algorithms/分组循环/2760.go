package loop

import "runtime"

//You are given a 0-indexed integer array nums and an integer threshold.
//
// Find the length of the longest subarray of nums starting at index l and endin
//g at index r (0 <= l <= r < nums.length) that satisfies the following conditions
//:
//
//
// nums[l] % 2 == 0
// For all indices i in the range [l, r - 1], nums[i] % 2 != nums[i + 1] % 2
// For all indices i in the range [l, r], nums[i] <= threshold
//
//
// Return an integer denoting the length of the longest such subarray.
//
// Note: A subarray is a contiguous non-empty sequence of elements within an arr
//ay.
//
//
// Example 1:
//
//
//Input: nums = [3,2,5,4], threshold = 5
//Output: 3
//Explanation: In this example, we can select the subarray that starts at l = 1
//and ends at r = 3 => [2,5,4]. This subarray satisfies the conditions.
//Hence, the answer is the length of the subarray, 3. We can show that 3 is the
//maximum possible achievable length.
//

// 模板， 越简单的题越容易出错！
func longestAlternatingSubarray(nums []int, threshold int) int {
	i := 0
	ans := 0
	n := len(nums)
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++
			continue
		}
		start := i
		i++ // 这个巧妙啊
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-start)
	}
	return ans
}


func longestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	ans := 0
	for i:=0; i< n; {
		//for ; i <n && (nums[i] > threshold || nums[i]&1 != 0); i++ {
		//}
		//if i >= n {  // 这里感觉好不优雅啊， 不如上面 if continue 的写法。
		//	return ans
		//}
		if nums[i]> threshold || nums[i]&1 != 0 {
			i++
			continue
		}
		start := i
		for i++; i<n && nums[i] <=threshold && nums[i]%2 != nums[i-1]%2; i++{
		}
		ans = max(ans, i-1 -start + 1)
	}
	return ans
}
