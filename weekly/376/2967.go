package _76

import (
	"fmt"
	"sort"
)

/*

给你一个长度为 n 下标从 0 开始的整数数组 nums 。

你可以对 nums 执行特殊操作 任意次 （也可以 0 次）。每一次特殊操作中，你需要 按顺序 执行以下步骤：

从范围 [0, n - 1] 里选择一个下标 i 和一个 正 整数 x 。
将 |nums[i] - x| 添加到总代价里。
将 nums[i] 变为 x 。
如果一个正整数正着读和反着读都相同，那么我们称这个数是 回文数 。比方说，121 ，2552 和 65756 都是回文数，但是 24 ，46 ，235 都不是回文数。

如果一个数组中的所有元素都等于一个整数 y ，且 y 是一个小于 109 的 回文数 ，那么我们称这个数组是一个 等数数组 。

请你返回一个整数，表示执行任意次特殊操作后使 nums 成为 等数数组 的 最小 总代价。



示例 1：

输入：nums = [1,2,3,4,5]
输出：6
解释：我们可以将数组中所有元素变为回文数 3 得到等数数组，数组变成 [3,3,3,3,3] 需要执行 4 次特殊操作，代价为 |1 - 3| + |2 - 3| + |4 - 3| + |5 - 3| = 6 。
将所有元素变为其他回文数的总代价都大于 6 。


示例 2：

输入：nums = [10,12,13,14,15]
输出：11
解释：我们可以将数组中所有元素变为回文数 11 得到等数数组，数组变成 [11,11,11,11,11] 需要执行 5 次特殊操作，代价为 |10 - 11| + |12 - 11| + |13 - 11| + |14 - 11| + |15 - 11| = 11 。
将所有元素变为其他回文数的总代价都大于 11 。
示例 3 ：

输入：nums = [22,33,22,33,22]
输出：22
解释：我们可以将数组中所有元素变为回文数 22 得到等数数组，数组变为 [22,22,22,22,22] 需要执行 2 次特殊操作，代价为 |33 - 22| + |33 - 22| = 22 。
将所有元素变为其他回文数的总代价都大于 22 。



这道题，能把大佬们恶心死！ 哈哈！
 */

func isPalindrome(n int) bool {
	if n == 0 {
		return true
	}
	digits := []int{}
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	if len(digits) == 1 {
		return true
	}
	i, j := 0, len(digits)-1
	for i < j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func minimumCost(nums []int) int64 {
	ans := int64(1e16)
	sort.Ints(nums)
	n := len(nums)
	mid := nums[(n-1)/2]
	if isPalindrome(mid) {
		ans = calSum(nums, mid)
	}

	var l, r int
	for l = mid - 1; true; l-- {
		if isPalindrome(l) {
			break
		}
	}
	for r = mid + 1; true; r++ {
		if isPalindrome(r) {
			break
		}
	}
	return min(calSum(nums, l), calSum(nums, r), ans)
}

/*
灵神的答案：
https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/solutions/2569308/yu-chu-li-hui-wen-shu-zhong-wei-shu-tan-7j0zy/
看灵神的思路， 中位数贪心， 证明很精彩。 思路清洗正规（不像自己是个草台班子的想法）
 */

func calSum(nums []int, k int) int64 {
	ans := 0
	for _, x := range nums {
		if x >= k {
			ans += x - k
		} else {
			ans += k - x
		}
	}
	return int64(ans)
}

func minimumCost(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	mid := nums[(n-1)/2]

	pnums := genNumbers(9)
	pnums = append(pnums, power(10, 9)+1)
	i := sort.SearchInts(pnums, mid)
	if pnums[i] <= nums[n/2] {
		return calSum(nums, mid)
	}
	return min(calSum(nums, pnums[i-1]), calSum(nums, pnums[i]))
}

/* 来一个快速幂的实现 */
func power(base, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * base
		}
		n = n >> 1
		base = base * base
	}
	return res
}

func genNumbers(zeros int) []int {
	ans := []int{}
	base := 1
	for base < power(10, (zeros+1)/2) {
		for i := base; i < base*10; i++ {
			// generate odd
			//t := i
			p := i
			//t /= 10
			for t := i; t > 0; t /= 10 {
				p = p*10 + t%10
				//t /= 10
			}
			ans = append(ans, p)
		}

		//generate even
		if base < power(10, zeros/2) {
			for i := base; i < base*10; i++ {
				t := i
				p := i
				for t > 0 {
					p = p*10 + t%10
					t /= 10
				}
				ans = append(ans, p)
			}
		}
		base = base * 10
		fmt.Println(base)
	}
	return ans
}
