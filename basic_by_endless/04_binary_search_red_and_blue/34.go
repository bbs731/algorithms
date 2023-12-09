package binary_search

func lower_bound(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n //  [l, r) 左闭右开的区间

	// sort.SearchInts()  的实现，就是这个版本 左闭右开
	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func lower_bound2(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1 // [l, r]  闭区间
	/* 循环不变量是： l-1 指向的是红色
	 				r+1 指向的是蓝色
	*/

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func lower_bound3(nums []int, target int) int {
	n := len(nums)
	l, r := -1, n // (l, r) 开区间

	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return l + 1
}

func lower_bound4(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1 // [l, r]  闭区间
	/* 循环不变量是： l-1 指向的是红色
	 				r+1 指向的是蓝色
	*/

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	start := lower_bound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lower_bound(nums, target+1) - 1
	return []int{start, end}

}
