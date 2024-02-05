package montonic_queue

func longestSubarray(nums []int, limit int) int {
	n := len(nums)
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
