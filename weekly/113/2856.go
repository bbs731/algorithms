package _13

import "sort"

// 这道题 是标准的面试题目。

/*
看看灵神优雅的解答：
 https://leetcode.cn/problems/minimum-array-length-after-pair-removals/solutions/2446146/olog-n-tan-xin-er-fen-cha-zhao-pythonjav-t3qn/

func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	x := nums[n/2]
	maxCnt := sort.SearchInts(nums, x+1) - sort.SearchInts(nums, x)
	return max(maxCnt*2-n, n%2)  // 这个max 用的也是 太巧妙了吧
}

 */
func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	cnt:=1
	maxCnt := 0
	for i:=1; i<n; i++ {
		if nums[i]== nums[i-1]{
			cnt++
			maxCnt = max(maxCnt, cnt)
		} else {
			cnt = 1
		}
	}

	if 2 * maxCnt > n  {
		return 2*maxCnt - n
	}
	if n % 2 == 1 {
		return 1
	}
	return 0
}


// 错误的答案
func minLengthAfterRemovals(nums []int) int {
	left := 0
	n := len(nums)

	for left < n {
		pos := sort.SearchInts(nums, nums[left]+1) - 1
		if pos-left+1 > n-1-pos {
			return pos - left + 1 - n + 1 + pos
		}
		left = left + (pos-left+1)*2
	}
	return 0
}
// 错误的答案。
func minLengthAfterRemovals(nums []int) int {
	left := 0
	n := len(nums)
	right := n - 1

	for left < right {
		if nums[left] != nums[right] {
			left++
			right--
		} else {
			break
		}
	}

	if left > right {
		return 0
	}
	return right - left + 1
}



