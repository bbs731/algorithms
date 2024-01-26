package binary_search

import (
	"sort"
)

/****


数对 (a,b) 由整数 a 和 b 组成，其数对距离定义为 a 和 b 的绝对差值。

给你一个整数数组 nums 和一个整数 k ，数对由 nums[i] 和 nums[j] 组成且满足 0 <= i < j < nums.length 。返回 所有数对距离中 第 k 小的数对距离。


示例 1：

输入：nums = [1,3,1], k = 1
输出：0
解释：数对和对应的距离如下：
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
距离第 1 小的数对是 (1,1) ，距离为 0 。
示例 2：

输入：nums = [1,1,1], k = 2
输出：0
示例 3：

输入：nums = [1,6,1], k = 3
输出：5


提示：

n == nums.length
2 <= n <= 10^4
0 <= nums[i] <= 10^6
1 <= k <= n * (n - 1) / 2

 */

// 二分答案，可以解答！
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)
	// (l, r]  这个区间应该是怎么写都可以 (l, r) 也是可以的。
	l, r := -1, int(1e6)

	for l < r {
		mid := (l + r + 1) >> 1

		tot := 0
		for i := 0; i < n; i++ {
			p := sort.SearchInts(nums, nums[i]+mid+1) - 1
			tot += p - i
		}

		if tot >= k {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return r + 1
}

// 这是灵神风格的代码了：翻译一遍就过， 继续保持。 二分的训练题，告一段落了， 下次再捡起来的时候，就开始做 >= 2100 的题目了。
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	//l, r := -1, int(1e6)+1

	return sort.Search(int(1e6), func(mid int) bool {

		tot := 0
		for i := 0; i < n; i++ {
			p := sort.SearchInts(nums, nums[i]+mid+1) - 1
			tot += p - i
		}
		// 因为是 先 false 后 true 序列， 所以，正常返回想要的结果就可以。
		return tot >= k

	})
}
