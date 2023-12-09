package binary_search


func lower_bound1(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1 // [l, r]  闭区间
	/* 循环不变量是： l-1 指向的是红色   <=l-1 都是红色
	 				r+1 指向的是蓝色   >= r+1 都是蓝色
	*/

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l  // l = r+1
}


func lower_bound2(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n //  [l, r) 左闭右开的区间

	// sort.SearchInts()  的实现，就是这个版本 左闭右开
	/* 循环不变量是什么？
		l-1 指向的是红色   <=l-1 都是红色
		r 指向的是蓝色   // r-1 指向的是红色，这句话是错误的， 因为初始化的时候，就不满足，循环的时候也不满足。
		>= r 指向蓝色
	 */
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


func lower_bound3(nums []int, target int) int {
	n := len(nums)
	l, r := -1, n // (l, r) 开区间
	/*
	   循环不变量： l 指向的是红色  <=l 都是红色
				  r 指向的是蓝色  >=r 都是蓝色
	 */
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
	l, r := -1, n-1 // (l, r] 左开右闭区间

	/*
		循环不变量：
				<=l   都是红色
				>=r+1 都是蓝色
	 */
	for l < r {
		mid := (r + l + 1)/2   // 因为下面可能出现 l = mid  所以， 这里需要向上取整

		if nums[mid] < target {
			l = mid
		} else {
			r = mid -1
		}
	}
	return r+1 // l == r
}
