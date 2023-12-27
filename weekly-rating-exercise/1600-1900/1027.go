package _600_1900

import (
	"sort"
)

/*

给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。

回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。



示例 1：

输入：nums = [3,6,9,12]
输出：4
解释：
整个数组是公差为 3 的等差数列。
示例 2：

输入：nums = [9,4,7,2,10]
输出：3
解释：
最长的等差子序列是 [4,7,10]。
示例 3：

输入：nums = [20,1,15,3,10,5,8]
输出：4
解释：
最长的等差子序列是 [20,15,10,5]。


提示：

2 <= nums.length <= 1000
0 <= nums[i] <= 500

 */
/*
难度分： 1759
标准的面试题难度吗？


1. 朴素的版本。

2. 可以用个 hash table 优化。
 */

func longestArithSeqLength(nums []int) int {
	ans := 0
	n := len(nums)

	dict := make(map[int][]int, n)
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}

	newd := make(map[int][]int, len(dict))
	for k, v := range dict {
		sort.Ints(v)
		newd[k] = v
	}

	for i := 0; i < n; i++ {
		for diff := 0; diff < 500; diff++ {
			next := nums[i] + diff
			cnt := 1
			//j := i + 1
			//for ; j < n && next <= 500; {
			//	l := newd[next]
			//	if len(l) == 0 {
			//		break
			//	}
			//	p := sort.SearchInts(l, j)
			//	if p == len(l) {
			//		break
			//	}
			//	cnt++
			//	j = p + 1
			//	next += diff
			//	ans = max(ans, cnt)
			//}
			for j := i + 1; j < n && next <= 500; j++ {
				if nums[j] == next {
					cnt++
					next += diff
					ans = max(ans, cnt)
				}
			}
		}
	}

	i, j := 0, n-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	dict = make(map[int][]int, n)
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}

	newd = make(map[int][]int, len(dict))
	for k, v := range dict {
		sort.Ints(v)
		newd[k] = v
	}

	for i := 0; i < n; i++ {
		for diff := 0; diff < 500; diff++ {
			next := nums[i] + diff
			cnt := 1
			//j := i + 1
			//for ; j < n && next <= 500; {
			//	l := newd[next]
			//	if len(l) == 0 {
			//		break
			//	}
			//	p := sort.SearchInts(l, j)
			//	if p == len(l) {
			//		break
			//	}
			//	cnt++
			//	j = p + 1
			//	next += diff
			//	ans = max(ans, cnt)
			//}
			for j := i + 1; j < n && next <= 500; j++ {
				if nums[j] == next {
					cnt++
					next += diff
					ans = max(ans, cnt)
				}
			}
		}
	}

	return ans
}
