package binary_search

import "sort"

/***
给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。

如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。

请注意，答案不一定是 arr 中的数字。



示例 1：

输入：arr = [4,9,3], target = 10
输出：3
解释：当选择 value 为 3 时，数组会变成 [3, 3, 3]，和为 9 ，这是最接近 target 的方案。
示例 2：

输入：arr = [2,3,5], target = 10
输出：5
示例 3：

输入：arr = [60864,25176,27249,21296,20204], target = 56803
输出：11361

提示：
1 <= arr.length <= 10^4
1 <= arr[i], target <= 10^5

 */

/***
这个二分答案， 墨迹了一个小时，
面试的时候肯定是废了
 */
func findBestValue(arr []int, target int) int {
	const maxn int = 1e9 + 1
	sort.Ints(arr)
	n := len(arr)
	presum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + arr[i-1]
	}

	// find the first mid that result in less then target
	l, r := -1, maxn
	diff1 := maxn
	diff2 := maxn
	// 二分答案
	for l+1 < r {
		mid := (l + r) >> 1
		p := sort.SearchInts(arr, mid)
		// arr[p] >= mid
		tot := presum[p] + (n-p)*mid
		if tot > target {
			r = mid
			diff2 = tot - target
			//diff2 = tot - target
		} else {
			l = mid
			diff1 = target - tot
		}
	}
	//l + 1 == r
	if diff1 <= diff2 {
		return min(l, arr[n-1])
	}
	return min(r, arr[n-1])
}
