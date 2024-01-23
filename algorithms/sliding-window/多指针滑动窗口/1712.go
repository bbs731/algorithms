package sliding_window

import (
	"fmt"
	"sort"
)

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


 //https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/solutions/544682/golang-jian-ji-xie-fa-by-endlesscheng-xaad/
// 看题解， 整理一下
func waysToSplit(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r := 2; r < n && 3*sum[r] <= 2*sum[n]; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sum[n]) + 1
		ans += sort.SearchInts(sum[l1:r], sum[r]/2+1)
	}
	return ans % (1e9 + 7)
}



// where sum[l:pos] <= v
func search(nums []int, l, r int, v int) int {
	for l+1 < r {
		mid := (l + r) / 2
		res := nums[mid]

		if res <= v {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 = r
	return l
}



// where sum[pos:right] >= v
func search2(psum[]int, l, r int, v int) int {
	for l+1 < r {
		mid := (l + r) / 2
		//res := psum[r] - psum[mid]
		if psum[mid]>=v {
			r = mid
		} else {
			l = mid
		}
	}
	// l + 1 = r
	return r
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
		sum := psum[right] - psum[0]
		limit := total - sum
		// no solution
		if limit < (sum+1)/2 {
			continue
		}
		if nums[0] > sum/2 {
			continue
		}

		if nums[right-1] > limit {
			continue
		}


		// now split the sum into two parts
		// find the pos sum[pos:right] >= sum/2
		left1 := search(psum[1:], -1, right, (sum)/2)
		//left1 := sort.SearchInts(psum, (sum+1)/2+1) -1
		// find the pos sum [pos: right] <=limit
		left2 := search2(psum[1:], -1, right, 2*psum[right]-total)

		fmt.Println(left2, left1, right, limit, sum, sum/2 )
		if left1 == left2 {
			ans += 1
		}
		if left1 > left2 {
			ans += left1 -left2
		} else {
			ans += left2 - left1
		}
		//if left1 > left2{
		//	ans +=  left1- left2
		//	ans %= mod
		//}
	}
	return ans %mod
}
