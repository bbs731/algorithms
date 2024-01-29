package montonic_stack

/***

给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。

由于答案可能很大，因此 返回答案模 10^9 + 7 。


示例 1：

输入：arr = [3,1,2,4]
输出：17
解释：
子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
示例 2：

输入：arr = [11,81,94,43,3]
输出：444


提示：

1 <= arr.length <= 3 * 10^4
1 <= arr[i] <= 3 * 10^4


 */

/***
第一次遇到，贡献法的问题， 太他吗的难了！  最后想到了，滑动数组中，求子数组个数的题目，猜的可以借鉴。


想了一个小时， 第一印象，想的是，用差分数组， 然 st 中的一段 accumulate 整体+1， 但是后来发现， st 上保存的区间，可能是不连续的，有 pop 过 elment 之后形成的空间。
就没办法用差分数组了。


还是看看灵神的题解： 这道题，必须作为模板题目

https://leetcode.cn/problems/sum-of-subarray-minimums/solutions/1930857/gong-xian-fa-dan-diao-zhan-san-chong-shi-gxa5/
灵神的题解，太强了，正宗的方法， 感觉我自己下面的方法，是在走偏门。
灵神的题解封神


ToDo: 用灵神的方法，重新做一遍吧！

 */

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	accumulate := make([]int, n)
	ans := 0
	st := []int{} // save the index

	for i := 0; i < n; i++ {
		accumulate[i] = 1
	}

	/****
	我自己的偏门的方法：
			 		【3， 1， 2， 4 ]
	accumulate:      [1,  1,  1,  1]

	st:              [3]
	acc: 			 [1]
	sum = 3
	ans = 3

	// pop 3,
	st:  			 [1]
	acc:			 [2]
	sum = 2
	ans = 3 + 2     (数字 3 贡献了 1 次, 数字 1贡献了 2次）

	st:				[1, 2]
	acc: 	        [2, 1]
	sum = 4
	ans = 3 + 2 + 4

	st: 			[1, 2, 4]
	acc: 		    [2, 1, 1]
	sum = 8
	ans = 3 + 2 + 4 + 8

	 */
	sum := 0
	for i, v := range arr {
		for len(st) > 0 && v < arr[st[len(st)-1]] {
			accumulate[i] += accumulate[st[len(st)-1]]

			top := st[len(st)-1]
			// pop the stack
			sum -= accumulate[top] * arr[top]
			st = st[:len(st)-1]
		}
		// insert into stack
		st = append(st, i)
		sum += accumulate[i] * v
		ans += sum
	}
	return ans % (int(1e9) + 7)
}
