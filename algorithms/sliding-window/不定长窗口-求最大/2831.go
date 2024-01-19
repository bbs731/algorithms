package weekly

/***

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

如果子数组中所有元素都相等，则认为子数组是一个 等值子数组 。注意，空数组是 等值子数组 。

从 nums 中删除最多 k 个元素后，返回可能的最长等值子数组的长度。

子数组 是数组中一个连续且可能为空的元素序列。



示例 1：

输入：nums = [1,3,2,3,1,3], k = 3
输出：3
解释：最优的方案是删除下标 2 和下标 4 的元素。
删除后，nums 等于 [1, 3, 3, 3] 。
最长等值子数组从 i = 1 开始到 j = 3 结束，长度等于 3 。
可以证明无法创建更长的等值子数组。
示例 2：

输入：nums = [1,1,2,2,1,1], k = 2
输出：4
解释：最优的方案是删除下标 2 和下标 3 的元素。
删除后，nums 等于 [1, 1, 1, 1] 。
数组自身就是等值子数组，长度等于 4 。
可以证明无法创建更长的等值子数组。


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= nums.length
0 <= k <= nums.length

 */

// let say we have a fixed sliding window size key, what's the longest sequence?
func sliding_window(nums []int, x, k int) bool {
	// 本来想， 返回一个 windows 中 最多element 出现的次数，想维护一个 heap， 结果发现想复杂了。
	n := len(nums)
	window_size := x + k
	max_cnts := 0
	left := 0
	cnts := make(map[int]int)
	for i := 0; i <= min(window_size-2, n-1); i++ {
		cnts[nums[i]]++
		max_cnts = max(max_cnts, cnts[nums[i]])
	}

	for i := window_size - 1; i < n; i++ {
		cnts[nums[i]]++
		max_cnts = max(max_cnts, cnts[nums[i]])
		if max_cnts >= x {
			return true
		}
		// remove the left element
		cnts[nums[left]]--
		left++
	}
	return false
}

// 时间复杂度是 O(n*logn)
func longestEqualSubarray(nums []int, k int) int {
	// 可以二分答案。
	cnts := make(map[int]int)
	max_cnts := 0
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		cnts[nums[i]]++
		max_cnts = max(max_cnts, cnts[nums[i]])
	}
	if max_cnts+k >= n {
		return max_cnts
	}

	left := 1
	right := max_cnts + 1 // 这里是开区间，所以别忘了+1 这是你犯得唯一错误了。真棒！2024/Jan/18 18：00按摩回来之后，就做出来了，发现上面其实不用 heap

	for left+1 < right {
		mid := (left + right) >> 1

		if sliding_window(nums, mid, k) == true {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

// 实在没想到，还能有 O(n) 线性复杂度的解法， 学习一下
// 这个想法太牛了！
func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}
	for _, ps := range pos {
		if len(ps) <= ans {
			continue
		}
		left := 0
		for right, p := range ps {
			for p-ps[left]-right+left > k { // 要删除的数太多了  ps[right] - ps[left] +1 - (right - left +1)
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	};
	return a
}

// 灵神的答案
//https://leetcode.cn/problems/find-the-longest-equal-subarray/solutions/2396401/fen-zu-shuang-zhi-zhen-pythonjavacgo-by-lqqau/

func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i-len(pos[x]))
	}
	for _, ps := range pos {
		if len(ps) <= ans {
			continue
		}
		left := 0
		for right, p := range ps {
			for p-ps[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}
