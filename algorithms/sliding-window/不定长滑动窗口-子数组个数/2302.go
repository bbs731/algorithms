package sliding_window


// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/solutions/1595722/by-endlesscheng-b120/
//
//

func countSubarrays(nums []int, k int64) int64 {
	left := 0
	ans := 0
	sum := 0

	for i, v:= range nums {
		sum += v
		for sum *(i-left+1) >= int(k) {
			sum -= nums[left]
			left++
		}
		// 这里的统计是关键， 如何理解？  当区间 [left, right] 中，所有已 right 为右端点的子集的个数。
		// 因为右端点是一值往后枚举的， 所以 ans += right - left +1  可以只计算当前  right 为右端点合法子集的个数。
		// 枚举循环结束， ans 即为所求， 因为所有可能 的子集都计算过了。

		/***
		固定子数组右端点 right，寻找满足条件的左端点 left，如果 sum(left...right) * (right - left + 1) < k，那么当 i > left 时，一定有 sum(i...right) * (right - i + 1) < k。
也就是，固定子数组右端点，寻找子数组分数 < k 的最长的那个子数组 [left...right]，那么比它短的所有子数组分数一定 < k，所以子数组个数是 right - left + 1 ~
		 */
		ans += i-left + 1  //
	}
	return int64(ans)
}


// 这个思路是错的。
func countSubarrays(nums []int, k int64) int64 {
	left := 0
	sum := 0
	ans := 0

	for i, v := range nums {
		sum += v
		if sum * (i-left +1) >= int(k) {
			// calculate the subset betwee [left, i-1]
			ans +=  (i-left) * (i-left+1)/2

			rsum := v
			for j := i-1; (rsum + nums[j])* (i-j+1) < int(k); j-- {
				ans += 1
				rsum += nums[j]
			}

			// shit left to right pos and start over again
			left = i
			sum = v
		}
	}
	// last element
	ans += 1
	return int64(ans)
}
