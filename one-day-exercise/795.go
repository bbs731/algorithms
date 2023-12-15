package one_day_exercise

//Given an integer array nums and two integers left and right, return the number
// of contiguous non-empty subarrays such that the value of the maximum array elem
//ent in that subarray is in the range [left, right].
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
//
//
// Example 1:
//
//
//Input: nums = [2,1,4,3], left = 2, right = 3
//Output: 3
//Explanation: There are three subarrays that meet the requirements: [2], [2, 1]
//, [3].
//
//
// Example 2:
//
//
//Input: nums = [2,9,2,5,6], left = 2, right = 8
//Output: 7
//
//



/*
牛啊！思路惊奇！除了一个小错误一次过！

https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/solutions/1988198/tu-jie-yi-ci-bian-li-jian-ji-xie-fa-pyth-n75l/
灵神的解答更震撼啊！
 */

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	lw := 0
	nums = append(nums, right+1) // append 一个dummy node 在最后，要不然，计算不到 nums[n-1]

	for i, x := range nums {
		if x > right {
			ans += cal(lw, i-1)
			lw = i + 1
		} else {
			//x <=right
			// do nothing
		}
	}

	lw = 0
	for i, x := range nums {
		if x >= left {
			ans -= cal(lw, i-1)
			lw = i + 1
		}
	}
	return ans
}

func cal(left, right int) int {
	if left > right {
		return 0
	}
	if right == left {
		return 1
	}
	l := right - left + 1
	return l * (l + 1) / 2
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	lw, ll := 0, 0
	nums = append(nums, right+1) // append 一个dummy node 在最后，要不然，计算不到 nums[n-1]

	for i, x := range nums {
		if x > right {
			ans += cal(lw, i-1)
			lw = i + 1
		}
		if x >= left {
			ans -= cal(ll, i-1)
			ll = i + 1
		}
	}
	return ans
}
