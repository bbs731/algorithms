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

3. 能想到是 DP ，但是不知道如何定义状态
https://leetcode.cn/problems/longest-arithmetic-subsequence/solutions/2239191/ji-yi-hua-sou-suo-di-tui-chang-shu-you-h-czvx/

maxLen[d]=max(maxLen[d],dfs(j)[d]+1)
是不是太不好想了， 让 dfs 返回一个 hashmap

参考一下灵神的答案， 然后用 DP 在做一遍

[24,13,1,100,0,94,3,0,3]
 */

/*
这是典型的， DP 优化朴素loop 循环的例子。
定义  dp[i] 为以 nums[i] 结尾的等差数列的长度。 那么，如何得到 dp[i]  就是对所有 j <i   dp[j] (以 nums[j]结尾的， nums[i]-nums[j] 的等差数列的长度 + 1） 我们还需要维护一个 i, d 的长度信息，可以想到用map 存储这样一个信息比较方便。

时间复杂度： O(n^2)  = 状态数 (O(n^2) * 单个状态更新的时候的 cost (O(1))
 */
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	ans := 0
	tables := make([]map[int]int, n)
	var dfs func(int) map[int]int
	dfs = func(i int) map[int]int {
		if tables[i] != nil {
			return tables[i]
		}

		t := make(map[int]int)
		tables[i] = t
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			//if t[d] == 0 {   // 这里可以优化。
			t[d] = max(t[d], dfs(j)[d]+1)
			ans = max(ans, t[d])
			//}
		}
		return t
	}
	dfs(n - 1)
	return ans + 1
}

/* 把上面的 dfs 变成递推。
dfs 改成 f 数组；
递归改成循环（每个参数都对应一层循环）；
递归边界改成 f 数组的初始值。


两层 loop
时间复杂度是 O（n^2)
*/

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	ans := 0
	//tables := make([]map[int]int, n)
	tables := make([][1001]int, n) // 可以用数组，这样速度更快，不用动态的 create map 了。
	//tables[0] = make(map[int]int) // 初始化条件

	for i := 1; i < n; i++ {
		//tables[i] = make(map[int]int)
		//t := tables[i]  // 这样写是不对的。 这样是copy array 不是 in-place 修改
		t := &tables[i] // 需要引用，因为需要修改
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j] + 500
			if t[d] == 0 { // 这里的 if 可以让速度提高一点点。
				t[d] = max(t[d], tables[j][d]+1)
				//t[d] = tables[j][d] + 1
				ans = max(ans, t[d])
			}
		}
	}
	return ans + 1
}

// 时间复杂度是 O( n * n * logn) 可以看一看， DP的解法是有何优化循环的。
func solve(nums []int) int {
	ans := 0
	n := len(nums)
	if n <= 2 {
		return n
	}

	dict := make(map[int][]int, n)
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}

	for k := range dict {
		sort.Ints(dict[k])
	}

	for i := 0; i < n; i++ {
		for diff := 0; diff < 500; diff++ {
			next := nums[i] + diff
			cnt := 1
			j := i + 1
			for ; j < n && next <= 500; {
				l := dict[next]
				if len(l) == 0 {
					break
				}
				p := sort.SearchInts(l, j)
				if p == len(l) {
					break
				}
				cnt++
				j = l[p] + 1
				next += diff
				ans = max(ans, cnt)
			}
		}
	}
	return ans
}
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	ans := solve(nums)

	i, j := 0, n-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	ans = max(ans, solve(nums))
	return ans
}
