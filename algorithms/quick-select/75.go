package quick_select


func sortColors(nums []int) {
	n :=len(nums)
	left, right :=0, n-1
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
			i++   // 所以 i 需要 ++
		} else {
			// x == 1
			i++
		}
	}

}

