package sliding_window

import "fmt"

/***

我们称一个分割整数数组的方案是 好的 ，当它满足：

数组被分成三个 非空 连续子数组，从左至右分别命名为 left ， mid ， right 。
left 中元素和小于等于 mid 中元素和，mid 中元素和小于等于 right 中元素和。
给你一个 非负 整数数组 nums ，请你返回 好的 分割 nums 方案数目。由于答案可能会很大，请你将结果对 109 + 7 取余后返回。



示例 1：

输入：nums = [1,1,1]
输出：1
解释：唯一一种好的分割方案是将 nums 分成 [1] [1] [1] 。
示例 2：

输入：nums = [1,2,2,2,5,0]
输出：3
解释：nums 总共有 3 种好的分割方案：
[1] [2] [2,2,5,0]
[1] [2,2] [2,5,0]
[1,2] [2,2] [5,0]
示例 3：

输入：nums = [3,2,1]
输出：0
解释：没有好的分割方案。

 */

// where sum[pos:right] <= v
func search2(nums []int, left, right int, v int) int {
	l := left - 1
	r := right + 1

	for l+1 < r {
		mid := (l + r) / 2
		res := nums[right+1] - nums[mid]

		if res > v {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 = r
	return r
}

// where sum[l:right] >= v
func search(nums []int, left, right int, v int) int {
	l := left - 1
	r := right + 1

	for l+1 < r {
		mid := (l + r) / 2
		res := nums[right+1] - nums[mid]

		if res < v {
			r = mid
		} else {
			l = mid
		}
	}
	// l + 1 = r
	return l
}

// 好题，我感觉我能做的出来。
func waysToSplit(nums []int) int {
	total := 0
	n := len(nums)
	psum := make([]int, n+1)
	ans := 0
	mod := int(1e9) + 7

	for i, v := range nums {
		total += v
		psum[i+1] = psum[i] + v
	}

	for right := range nums {
		if right <= 1 {
			continue
		}

		sum := psum[right+1] - psum[0]
		limit := total - sum
		// no solution
		if limit < (sum+1)/2 {
			continue
		}

		// now split the sum into two parts
		// find the pos sum[pos:right] >= sum/2
		left1 := search(psum, 0, right, sum/2)
		// find the pos sum [pos: right] <=limit
		left2 := search2(psum, 0, right, limit)

		fmt.Println(left2, left1, limit, sum/2)
		if left2 >= left1 {
			ans += left2 - left1 + 1
			ans %= mod
		}
	}
	return ans
}
