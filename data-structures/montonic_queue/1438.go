package montonic_queue

/***
神级别的题目。（多做几遍，用不同的方法）
 */
func longestSubarray(nums []int, limit int) int {
	n := len(nums)

	// 能想到，固定窗口大小，然后做单调队列，已经是很大的进步？ 但是是否维护单调队列，就一定要 fixed windows 的滑动窗口呢？
	// 答案是肯定的， 可以去掉 fixed window size 的限制， 继续优化, 把时间复杂度降低到 O（n)
	find_windows_m := func(nums []int, m int) bool {
		qmin := make([]int, 0)
		qmax := make([]int, 0)
		for i, num := range nums {
			// qmin[0] out of index  // 这里用 for 不要用 if, 有些测试用例过不了, 还是 corner case 没考虑清楚。
			for len(qmin) > 0 && qmin[0] < i-m+1 {
				qmin = qmin[1:]
			}

			for len(qmin) > 0 && nums[qmin[len(qmin)-1]] >= num {
				// pop back
				qmin = qmin[:len(qmin)-1]
			}
			qmin = append(qmin, i)

			for len(qmax) > 0 && qmax[0] < i-m+1 {
				qmax = qmax[1:]
			}

			for len(qmax) > 0 && nums[qmax[len(qmax)-1]] <= num {
				qmax = qmax[:len(qmax)-1]
			}
			qmax = append(qmax, i)

			if nums[qmax[0]]-nums[qmin[0]] <= limit && i >= m-1 { // 这里的 i >=m-1 重要，要不然 提前返回 true 了。
				return true
			}
		}
		return false
	}

	// 先 true 后 false 的序列
	l, r := 0, n+1
	for l+1 < r {
		mid := (l + r) >> 1
		if find_windows_m(nums, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	// l +1 = r
	return l

}

/***
把 二分枚举 windows 大小， 变成 滑动窗口（不在 fix window 大小，而是根据 limit 动态调整 sliding window 的左端点） 结果还能降低时间的复杂度到 O(n)
 */
func longestSubarray(nums []int, limit int) int {
	left := 0 //sliding window left
	ans := 1
	qmin := make([]int, 0)
	qmax := make([]int, 0)
	for i, num := range nums { // 枚举 sliding window 的右端点
		// insert i into monotonic queue   // 先插入， 后 pop 的逻辑，写出来更加的简洁。
		for len(qmin) > 0 && nums[qmin[len(qmin)-1]] >= num {
			// pop back
			qmin = qmin[:len(qmin)-1]
		}
		qmin = append(qmin, i)

		for len(qmax) > 0 && nums[qmax[len(qmax)-1]] <= num {
			qmax = qmax[:len(qmax)-1]
		}
		qmax = append(qmax, i)

		//这么写，更加优雅一些。
		// 调整 sliding window 的左端点
		for left <= i-1 && nums[qmax[0]]-nums[qmin[0]] > limit {
			left++
			for len(qmin) > 0 && qmin[0] < left {
				qmin = qmin[1:]
			}

			for len(qmax) > 0 && qmax[0] < left {
				qmax = qmax[1:]
			}
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

//func longestSubarray(nums []int, limit int) int {
//	left := 0 //sliding window left
//	ans := 1
//	qmin := make([]int, 0)
//	qmax := make([]int, 0)
//	for i, num := range nums {
//		// insert i into monotonic queue
//		for len(qmin) > 0 && nums[qmin[len(qmin)-1]] >= num {
//			// pop back
//			qmin = qmin[:len(qmin)-1]
//		}
//		qmin = append(qmin, i)
//
//		for len(qmax) > 0 && nums[qmax[len(qmax)-1]] <= num {
//			qmax = qmax[:len(qmax)-1]
//		}
//		qmax = append(qmax, i)
//
//		for left <= i-1 {
//			for len(qmin) > 0 && qmin[0] < left {
//				qmin = qmin[1:]
//			}
//
//			for len(qmax) > 0 && qmax[0] < left {
//				qmax = qmax[1:]
//			}
//
//			if nums[qmax[0]]-nums[qmin[0]] <= limit {
//				ans = max(ans, i-left+1)
//				break
//			} else {
//				left++
//			}
//		}
//	}
//	return ans
//}
