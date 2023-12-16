package _600_1900

/*
示例 3：

输入：nums = [3,2,20,1,1,3], x = 10
输出：5
解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。

这道题，可以转化为， 求一段连续的 子串， 让其和是 (sum -x), 并且让子串的长度最长。

tags:
前缀和 + 双指针

自己想的，解法我觉得挺好的！ 但是写出bug-free 的代码还是要练习啊!
*/
func minOperations(nums []int, x int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	target := presum[n] - x

	// 这个判断的条件，是不是太恶心了！ 应对 [1,2,3]  x=6  这种恶心的 test case
	if target == 0 {
		return n
	}

	l, r := 0, 0
	ans := -1
	for l <= r && r < n {
		sum := presum[r+1] - presum[l]
		if sum == target {
			ans = max(ans, r-l+1)
		}
		if l == r {
			r++
		} else {
			if sum <= target {
				r++
			} else {
				l++
			}
		}
	}
	if ans == -1 {
		return ans
	}
	return n - ans
}

/*
参考灵神的写法：
https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/
 */
func minOperations(nums []int, x int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	target := presum[n] - x

	// 这个判断的条件，是不是太恶心了！ 应对 [1,2,3]  x=6  这种恶心的 test case
	// [1, 1]  x = 3
	if target < 0 {
		return -1
	}

	l, ans :=0, -1
	for r := range nums {   // 这种 loop 不容易错是吗？
		sum := presum[r+1] - presum[l]
		for sum > target {
			sum -= nums[l]
			l++
		}
		if sum == target {
			ans = max(ans, r-l+1)
		}
	}
	if ans == -1 {
		return ans
	}
	return n - ans
}

