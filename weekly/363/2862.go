package weekly

import "fmt"

/*
给你一个下标从 1 开始、由 n 个整数组成的数组。

如果一组数字中每对元素的乘积都是一个完全平方数，则称这组数字是一个 完全集 。

下标集 {1, 2, ..., n} 的子集可以表示为 {i1, i2, ..., ik}，我们定义对应该子集的 元素和 为 nums[i1] + nums[i2] + ... + nums[ik] 。

返回下标集 {1, 2, ..., n} 的 完全子集 所能取到的 最大元素和 。

完全平方数是指可以表示为一个整数和其自身相乘的数。



示例 1：

输入：nums = [8,7,3,5,7,2,4,9]
输出：16
解释：除了由单个下标组成的子集之外，还有两个下标集的完全子集：{1,4} 和 {2,8} 。
与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 8 + 5 = 13 。
与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 7 + 9 = 16 。
因此，下标集的完全子集可以取到的最大元素和为 16 。
示例 2：

输入：nums = [5,10,3,10,1,13,7,9,4]
输出：19
解释：除了由单个下标组成的子集之外，还有四个下标集的完全子集：{1,4}、{1,9}、{2,8}、{4,9} 和 {1,4,9} 。
与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 5 + 10 = 15 。
与下标 1 和 9 对应的元素和等于 nums[1] + nums[9] = 5 + 4 = 9 。
与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 10 + 9 = 19 。
与下标 4 和 9 对应的元素和等于 nums[4] + nums[9] = 10 + 4 = 14 。
与下标 1、4 和 9 对应的元素和等于 nums[1] + nums[4] + nums[9] = 5 + 10 + 4 = 19 。
因此，下标集的完全子集可以取到的最大元素和为 19 。


提示：

1 <= n == nums.length <= 104  哎， 这里是 10^4  吃了无数次亏了！
1 <= nums[i] <= 109

*/

/*
既然是子集问题，就去套用子集问题的模板啊！
去温习一下， backtracking 子集问题。 再回来做一下！

你真是太有任性了， 努力拼搏之后，在failed test case 的提示之下终于可以了。
backtracking + cache 无敌。

看灵神的答案， 非常简单：
https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/solutions/2446037/an-zhao-corei-fen-zu-pythonjavacgo-by-en-i6nu/
*/

func maximumSum(nums []int) int64 {
	ans := nums[0]
	n := len(nums)

	cache := make([][]int, n)
	for i:=0; i<n; i++ {
		cache[i]= make([]int, 101)
		for j:=0; j<=100; j++ {
			cache[i][j] = -1
		}
	}


	path := []int{}
	var dfs func(int, int)int
	dfs = func(start, i int) int {
		if start * i * i >n {
			// 统计结果。
			// len(path） = 1 的结果是合法的！
			//if len(path)==1 && candidates[path[0]]== 0 {
			//	return
			//}
			sum := 0
			for j :=0; j<len(path); j++ {
				sum += nums[path[j]-1]
			}
			//ans = max(ans, sum)
			return sum
		}

		if cache[start][i] != -1 {
			return cache[start][i]
		}

		res := 0
		// 选 i
		//if start * i *i  <= n {
		path = append(path, start*i*i)
		res = max(res, dfs(start, i+1))
		path = path[:len(path)-1]
		//}

		// 不选 i
		//res = max(res, dfs(start, i+1))  // 其实是不需要 不选的选项的。

		cache[start][i] = res
		return res
	}

	for i := 1; i <= n; i++ {
		path = append(path, i)
		ans = max(ans, dfs(i, 2)) // 那就直接变成一个 循环就可以了， 不需要 dfs + cache  了。
		path = path[:len(path)-1]
	}
	return int64(ans)
}

/*

这个想法是不对的。
*/

func maximumSum(nums []int) int64 {
	ans := 0
	n := len(nums)
	for i := 6; i <= 10 && i*i < n; i++ {
		ans = max(ans, nums[i*i])
	}

	for i := 1; i <= 5 && i < n; i++ {
		ans += nums[i]
	}

	for i := 4; i <= 10 && i < n; i++ {
		for j := 2; j < i; j++ {
			if (i*i)%j == 0 && i*i/j < n {
				ans = max(ans, nums[i]+nums[i*i/j])
			}
		}
	}
	return int64(ans)
}
