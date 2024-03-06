package quick_select

func sortColors(nums []int) {
	n := len(nums)
	left, right := 0, n-1
	i := 0
	/***
	题解：  https://leetcode.cn/problems/sort-colors/solutions/2651983/yi-kan-jiu-dong-de-dai-ma-by-bo-bo-t1-jlht/
	[0, left)   是 0
	[left, right]  是 1
	(right, n-1] 是 2
	***/
	for i <= right {
		x := nums[i]
		if x == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else if x == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++ // left 指向的位置是 1
			i++    // 所以 i 需要 ++
		} else {
			// x == 1
			i++
		}
	}
}

func sortColors(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	// invariant  <l is 0  and >r is 2

	for i := 0; i < n; {
		c := nums[i]
		if c == 1 || i < l {
			// all zeros 所以直接 i++
			i++
			continue
		}
		if i > r {
			// now all 2s, can break now
			break
		}

		// 不能更新 i, 因为 swap 的 nums[l], nums[r] 现在变成了 nums[i], nums[i]的值是什么不确定， 需要再次被loop处理。
		if c == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		}
	}
}
